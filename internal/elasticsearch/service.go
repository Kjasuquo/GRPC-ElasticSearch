package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

const mapping = `{
	  "mappings": {
    "properties": {
      "Name": {
        "type": "completion"
      }
    }
  },
  "settings": {
    "analysis": {
      "filter": {
        "edge_ngram_filter": {
          "type": "edge_ngram",
          "min_gram": 1,
          "max_gram": 20
        }
      },
      "analyzer": {
        "autocomplete": {
          "type": "custom",
          "tokenizer": "standard",
          "filter": [
            "lowercase",
            "edge_ngram_filter"
          ]
        }
      }
    }
  }
}`

func (es *DB) CreateIndex(indexName string) error {
	// create a new index
	ctx := context.Background()
	_, err := es.Client.CreateIndex(indexName).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("%s created\n", indexName)

	return nil
}

func (es *DB) CreateSuggestionIndex(indexName string) error {
	ctx := context.Background()
	// Create a new index.
	_, err := es.Client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}

	fmt.Println("index created with mapping")
	return nil
}

func (es *DB) SearchSuggestion(indexName string, field string, query string) ([]string, error) {
	ctx := context.Background()
	Suggester := elastic.NewCompletionSuggester("suggest").Fuzziness(1).Text(query).Field(field).SkipDuplicates(true).Analyzer("autocomplete")

	searchSource := elastic.NewSearchSource().
		Suggester(Suggester).
		FetchSource(false).
		TrackScores(true)

	searchResult, err := es.Client.Search().
		Index(indexName).
		SearchSource(searchSource).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		// Handle error
		return nil, err
	}

	suggestedNames := searchResult.Suggest["suggest"]
	var results []string
	for _, options := range suggestedNames {
		for _, option := range options.Options {
			results = append(results, option.Text)
		}
	}
	return results, nil
}

func (es *DB) InsertData(indexName string, data interface{}) error {
	ctx := context.Background()
	_, err := es.Client.Index().
		Index(indexName).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	fmt.Println("data inserted")
	return nil
}

func (es *DB) SearchData(indexName string, query elastic.Query) ([]map[string]interface{}, error) {
	ctx := context.Background()
	searchResult, err := es.Client.Search().
		Index(indexName).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		// Handle error
		return nil, err
	}
	var results []map[string]interface{}
	var ttyp map[string]interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if result, ok := item.(map[string]interface{}); ok {
			results = append(results, result)
		}
	}
	return results, nil
}

func (es *DB) SearchAllData(indexName string) ([]map[string]interface{}, error) {
	ctx := context.Background()
	searchResult, err := es.Client.Search().
		Index(indexName).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		// Handle error
		return nil, err
	}
	var answer []map[string]interface{}
	var ttyp map[string]interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if result, ok := item.(map[string]interface{}); ok {
			answer = append(answer, result)
		}
	}
	return answer, nil
}

func (es *DB) DeleteData(indexName string, query elastic.Query) error {
	ctx := context.Background()

	_, err := es.Client.DeleteByQuery(indexName).Query(query).
		Do(ctx)

	if err != nil {
		return err
	}
	fmt.Println("data deleted")
	return nil
}

func (es *DB) DeleteIndex(indexName string) error {
	ctx := context.Background()
	_, err := es.Client.DeleteIndex(indexName).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	fmt.Println(indexName, " deleted")
	return nil
}
