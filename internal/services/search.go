package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gitlab.com/grpc-buffer/proto/go/pkg/proto"
	//"gitlab.com/dh-backend/search-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *ElasticSearchServer) SearchItems(ctx context.Context, req *proto.InSearchItemRequest) (*proto.InSearchItemResponse, error) {
	value := req.GetValue()
	var result []map[string]interface{}

	// Variable to store the response from the search service in the form of a list of items
	var items []*proto.SearchItem

	var err error

	if value == "" {
		result, err = s.Elasticsearch.SearchAllData(ItemIndex)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Search all Items Error: %v\n", err))
		}

	} else {
		result, err = s.Elasticsearch.SearchData(ItemIndex, elastic.NewMatchQuery(QueryField, value))
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Search Items Error: %v\n", err))
		}
	}

	Data, err := json.Marshal(result)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot Marshall Result: %v\n", err))
	}

	err = json.Unmarshal(Data, &items)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot unmarshall items: %v\n", err))
	}

	log.Println("This is the item gotten Data", items)

	return &proto.InSearchItemResponse{
		Status:       codes.OK.String(),
		ItemResponse: items,
	}, nil

}

func (s *ElasticSearchServer) SearchPackages(ctx context.Context, req *proto.InSearchPackageRequest) (*proto.InSearchPackageResponse, error) {
	value := req.GetValue()
	var result []map[string]interface{}

	// Variable to store the response from the search service in the form of a list of packages
	var packages []*proto.SearchPackage

	var err error

	if value == "" {
		result, err = s.Elasticsearch.SearchAllData(PackageIndex)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Search all Packages Error: %v\n", err))
		}
	} else {
		result, err = s.Elasticsearch.SearchData(PackageIndex, elastic.NewMatchQuery(QueryField, value))
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Search Packages Error: %v\n", err))
		}
	}

	Data, err := json.Marshal(result)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot Marshall Result: %v\n", err))
	}

	err = json.Unmarshal(Data, &packages)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot unmarshall packages: %v\n", err))
	}

	log.Println("This is the packages gotten Data", packages)

	return &proto.InSearchPackageResponse{
		Status:          codes.OK.String(),
		PackageResponse: packages,
	}, nil
}

func (s *ElasticSearchServer) SearchService(ctx context.Context, req *proto.InSearchRequest) (*proto.InSearchResponse, error) {
	index := req.GetIndex()
	value := req.GetValue()

	var result []map[string]interface{}
	var err error

	if index == "package" {
		if value == "" {
			fmt.Println(index, ":", value)
			result, err = s.Elasticsearch.SearchAllData(PackageIndex)
			if err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("error searching packages in search-service: %v\n", err))
			}
		} else {
			result, err = s.Elasticsearch.SearchData(PackageIndex, elastic.NewMatchQuery(QueryField, value))
			if err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("error searching packages in search-service: %v\n", err))
			}
		}
	} else if index == "item" {
		if value == "" {
			result, err = s.Elasticsearch.SearchAllData(ItemIndex)
			if err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("error searching items in search-service: %v\n", err))
			}

		} else {
			result, err = s.Elasticsearch.SearchData(ItemIndex, elastic.NewMatchQuery(QueryField, value))
			if err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("error searching items in search-service: %v\n", err))
			}
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
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot marshall items: %v\n", err))
		}
		err = json.Unmarshal(Data, &items)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot unmarshall items: %v\n", err))
		}
		log.Println("Items gotten")

	} else if okP {
		Data, err := json.Marshal(result)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot marshall packages: %v\n", err))
		}

		err = json.Unmarshal(Data, &packages)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot unmarshall packages: %v\n", err))
		}
		log.Println("packages gotten")
	}

	return &proto.InSearchResponse{
		Status:          codes.OK.String(),
		PackageResponse: packages,
		ItemResponse:    items,
	}, nil

}

// SearchItemSuggestions is for suggesting items to be searched based on prefixes and available items in DB
func (s *ElasticSearchServer) SearchItemSuggestions(ctx context.Context, req *proto.InSearchItemSuggestionsRequest) (*proto.InSearchItemSuggestionsResponse, error) {

	value := req.GetValue()

	var result []string

	result, err := s.Elasticsearch.SearchSuggestion(ItemSuggestionIndex, QueryField, value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Items Suggestion Error: %v\n", err))
	}

	return &proto.InSearchItemSuggestionsResponse{
		Status:      codes.OK.String(),
		Suggestions: result,
	}, nil
}

// SearchPackageSuggestions is for suggesting packages to be searched based on prefixes and available items in DB
func (s *ElasticSearchServer) SearchPackageSuggestions(ctx context.Context,
	req *proto.InSearchPackageSuggestionsRequest) (*proto.InSearchPackageSuggestionsResponse, error) {
	value := req.GetValue()

	var result []string

	result, err := s.Elasticsearch.SearchSuggestion(PackageSuggestionIndex, QueryField, value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Packages Suggestion Error: %v\n", err))
	}
	log.Println("This is the package suggestions gotten", result)
	return &proto.InSearchPackageSuggestionsResponse{
		Status:      codes.OK.String(),
		Suggestions: result,
	}, nil
}
