package rabbitMQ

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"gitlab.com/dh-backend/search-service/internal/elasticsearch"
	"log"
	"time"
)

var Connection *amqp.Connection

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func ConnectRabbitMq(rabbitMQURL string) {

	fmt.Println("Connecting to RabbitMQ ...")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(rabbitMQURL)

	log.Println("url", rabbitMQURL)

	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %s", err)
	}

	fmt.Println("Connected to RabbitMQ ...")

	Connection = connectRabbitMQ
}

func PublishToElasticCreationQueue(item map[string]interface{}, orderChannel string) error {
	ch, err := Connection.Channel()
	failOnError(err, "Failed to open a channel at Publish")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		orderChannel, // name
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	body, err := json.Marshal(item)
	if err != nil {
		log.Printf("cannot marshal this: %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "json",
			Body:        body,
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Printf("Sent %v\n", body)
	fmt.Println("Successfully published message to the queue")
	return nil
}

func ReadFromQueueToInsertInES(name string, localChannelName string) {
	ch, err := Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		localChannelName, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var Data = make(map[string]interface{})

	var forever chan struct{}

	go func() {
		for d := range msgs {
			// log.Printf("our consumer received a message from: %v", os.Getpid())
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal((d.Body), &Data)
			if err != nil {
				log.Fatalf("error%v", err)
			}

			client := elasticsearch.GetESClient()
			es := elasticsearch.NewElasticSearchDB(client)
			es.InsertData(name, Data)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
