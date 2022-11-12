package ports

import "github.com/olivere/elastic/v7"

type Elasticsearch interface {
	CreateIndex(indexName string)
	InsertData(indexName string, data interface{})
	SearchData(indexName string, query elastic.Query) []map[string]interface{}
	DeleteData(indexName string, query elastic.Query)
	DeleteIndex(indexName string)
	SearchAllData(indexName string) []map[string]interface{}
}
