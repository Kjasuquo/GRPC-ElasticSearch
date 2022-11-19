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

	go rabbitMQ.ReadFromItemQueueToInsertInES(channel, MenuChannel)

	var forever chan struct{}
	go func() {

		// Read from channel
		for {

			Data, _ := <-channel
			_, okI := Data["ItemCategory"]
			_, okP := Data["packageID"]

			if okI {
				err := es.InsertData(ItemIndex, Data)
				if err != nil {
					log.Printf("Error while inserting data in ES: %v", err)
				}
				err = es.InsertData(ItemSuggestionIndex, Data)
				if err != nil {
					log.Printf("Error Insering Item for suggestion: %v", err)
				}
				log.Println("Item Inserted")
			} else if okP {
				err := es.InsertData(PackageIndex, Data)
				if err != nil {
					log.Printf("Error while inserting data in ES: %v", err)
				}
				err = es.InsertData(PackageSuggestionIndex, Data)
				if err != nil {
					log.Printf("Error Insering Package for suggestion: %v", err)
				}
				log.Println("Package Inserted")
			}

		}

	}()

	<-forever
}
