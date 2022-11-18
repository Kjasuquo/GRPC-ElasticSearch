package elasticsearch

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

type DB struct {
	Client *elastic.Client
}

func GetESClient(addr string) *elastic.Client {

	client, err := elastic.NewClient(elastic.SetURL(addr),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	fmt.Println("Elastic Search initialized...")

	return client
}

func NewElasticSearchDB(db *elastic.Client) *DB {
	return &DB{Client: db}
}
