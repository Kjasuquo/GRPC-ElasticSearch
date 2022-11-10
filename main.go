package main

import (
	myService "gitlab.com/dh-backend/search-service/internal/services"
	"log"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//client := elasticsearch.GetESClient()
	//es := elasticsearch.NewElasticSearchDB(client)
	//fmt.Println(es.SearchData("data", elastic.NewMatchQuery("restaurantID", "minim")))
	myService.Start()
}
