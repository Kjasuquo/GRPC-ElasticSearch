package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gitlab.com/grpc-buffer/proto/go/pkg/proto"
	"google.golang.org/grpc/codes"
	"log"
)

func (s *ElasticSearchServer) SearchService(ctx context.Context, req *proto.InSearchRequest) (*proto.InSearchResponse, error) {
	index := req.GetIndex()
	value := req.GetValue()

	var result []map[string]interface{}

	if index == "package" {
		if value == "" {
			fmt.Println(index, ":", value)
			result = s.Elasticsearch.SearchAllData("package")
		} else {
			result = s.Elasticsearch.SearchData("package", elastic.NewMatchQuery("Name", value))
		}
	} else if index == "item" {
		if value == "" {
			result = s.Elasticsearch.SearchAllData("item")

		} else {
			result = s.Elasticsearch.SearchData("item", elastic.NewMatchQuery("Name", value))
		}
	}

	// Variable to store the response from the search service in the form of a list of items
	var items []*proto.SearchItem

	// Variable to store the response from the search service in the form of a list of packages
	var packages []*proto.SearchPackage

	_, okP := result[0]["Images"]
	_, okI := result[0]["Image"]

	// Checking if the response is a list of items or a list of packages
	if okI {
		Data, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(Data, &items)
		if err != nil {
			log.Fatalf("error%v", err)
		}
		log.Println("Items gotten")

	} else if okP {
		Data, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(Data, &packages)
		if err != nil {
			log.Fatalf("error%v", err)
		}
		log.Println("packages gotten")
	}

	log.Println("This is the package gotten Data", packages)
	log.Println("This is the item gotten Data", items)

	return &proto.InSearchResponse{
		Status:          codes.OK.String(),
		PackageResponse: packages,
		ItemResponse:    items,
	}, nil

}
