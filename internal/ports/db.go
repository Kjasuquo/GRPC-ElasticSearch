package ports

import "github.com/olivere/elastic/v7"

type Elasticsearch interface {
	CreateIndex(indexName string) error
	InsertData(indexName string, data interface{}) error
	SearchData(indexName string, query elastic.Query) ([]map[string]interface{}, error)
	DeleteData(indexName string, query elastic.Query) error
	DeleteIndex(indexName string) error
	SearchAllData(indexName string) ([]map[string]interface{}, error)
	CreateSuggestionIndex(indexName string) error
	SearchSuggestion(indexName string, field string, query string) ([]string, error)
}
