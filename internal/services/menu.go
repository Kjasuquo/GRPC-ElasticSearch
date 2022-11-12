package services

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"gitlab.com/dh-backend/search-service/internal/elasticsearch"
	"gitlab.com/dh-backend/search-service/internal/rabbitMQ"
	"log"
)

func RabbitMQServicesForMenu() {
	channel := make(chan map[string]interface{})
	client := elasticsearch.GetESClient()
	es := elasticsearch.NewElasticSearchDB(client)

	go rabbitMQ.ReadFromItemQueueToInsertInES(channel, "menuCreationChannel")

	var forever chan struct{}
	go func() {

		// Read from channel
		for {

			Data, _ := <-channel
			v, okQ := Data["index"]
			_, okI := Data["Image"]
			_, okP := Data["Images"]

			// If the data is of type item, package or just a query
			if okQ {
				fmt.Println(Data["value"], "This Data is from Queue")
				if v == "package" {
					if Data["value"] == "" {
						fmt.Println(Data["index"], ":", Data["value"])
						result := es.SearchAllData("package")
						err := rabbitMQ.PublishToElasticCreationQueue(result, "searchedResultChannel")
						if err != nil {
							log.Fatalf("error publishing to queue: %v", err)
						}
						log.Println("result published to searchedResultChannel", result)
					} else {
						result := es.SearchData("package", elastic.NewMatchQuery("Name", Data["value"]))
						err := rabbitMQ.PublishToElasticCreationQueue(result, "searchedResultChannel")
						if err != nil {
							log.Fatalf("error publishing to queue: %v", err)
						}
						log.Println("result published to searchedResultChannel", result)
					}
				} else if v == "item" {
					if Data["value"] == "" {
						result := es.SearchAllData("item")
						err := rabbitMQ.PublishToElasticCreationQueue(result, "searchedResultChannel")
						if err != nil {
							log.Fatalf("error publishing to queue: %v", err)
						}
						log.Println("result published to searchedResultChannel")
					} else {
						result := es.SearchData("item", elastic.NewMatchQuery("Name", Data["Value"]))
						err := rabbitMQ.PublishToElasticCreationQueue(result, "searchedResultChannel")
						if err != nil {
							log.Fatalf("error publishing to queue: %v", err)
						}
						log.Println("result published to searchedResultChannel")
					}
				}
			} else if okI {
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
