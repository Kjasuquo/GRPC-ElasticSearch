package services

import (
	"gitlab.com/dh-backend/search-service/config"
	"gitlab.com/dh-backend/search-service/internal/elasticsearch"
	"gitlab.com/dh-backend/search-service/internal/rabbitMQ"
	"log"
)

func RabbitMQServicesForMenu() {
	channel := make(chan map[string]interface{})
	configs := config.ReadConfigs(".")
	client := elasticsearch.GetESClient(configs.ElasticSearchUrl)
	es := elasticsearch.NewElasticSearchDB(client)

	go rabbitMQ.ReadFromItemQueueToInsertInES(channel, "menuCreationChannel")

	var forever chan struct{}
	go func() {

		// Read from channel
		for {

			Data, _ := <-channel
			_, okI := Data["Image"]
			_, okP := Data["Images"]

			if okI {
				es.InsertData("item", Data)
				log.Println("Item Inserted")
			} else if okP {
				es.InsertData("package", Data)
				log.Println("Package Inserted")
			}

		}

	}()

	<-forever
}
