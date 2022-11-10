package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)

func (es *DB) CreateIndex(indexName string) {
	// create a new index
	ctx := context.Background()
	_, err := es.Client.CreateIndex(indexName).Do(ctx)
	if err != nil {
		// Handle error
		log.Fatalf("failed to create index: %v", err)
	}
	fmt.Println("index created")
}

func (es *DB) InsertData(indexName string, data interface{}) {
	ctx := context.Background()
	_, err := es.Client.Index().
		Index(indexName).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		// Handle error
		log.Fatalf("failed to insert data: %v", err)
	}
	fmt.Println("data inserted")
}

func (es *DB) SearchData(indexName string, query elastic.Query) []map[string]interface{} {
	ctx := context.Background()
	searchResult, err := es.Client.Search().
		Index(indexName).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		// Handle error
		log.Fatalf("failed to search data: %v", err)
	}
	var results []map[string]interface{}
	var ttyp map[string]interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if result, ok := item.(map[string]interface{}); ok {
			results = append(results, result)
		}
	}
	return results
}

func (es *DB) DeleteData(indexName string, query elastic.Query) {
	ctx := context.Background()

	_, err := es.Client.DeleteByQuery(indexName).Query(query).
		Do(ctx)

	if err != nil {
		log.Fatalf("failed to delete data: %v", err)
	}
	fmt.Println("data deleted")
}

func (es *DB) DeleteIndex(indexName string) {
	ctx := context.Background()
	_, err := es.Client.DeleteIndex(indexName).Do(ctx)
	if err != nil {
		// Handle error
		log.Fatalf("failed to delete index: %v", err)
	}
	fmt.Println("index deleted")
}

//func (es *ElasticSearchDB) UpdateData(indexName string, query elastic.Query, data interface{}) {
//	ctx := context.Background()
//	_, err := es.Client.UpdateByQuery().
//		Index(indexName).
//		Query(query).
//		Script(elastic.NewScriptInline("ctx._source.age = 100")).
//		Do(ctx)
//
//	if err != nil {
//		fmt.Println("[Update]Error=", err)
//		return
//	}
//
//	fmt.Println("Updated")
//}
