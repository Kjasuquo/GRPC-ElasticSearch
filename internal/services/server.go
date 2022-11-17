package services

import (
	"fmt"
	"gitlab.com/dh-backend/search-service/config"
	"gitlab.com/dh-backend/search-service/internal/elasticsearch"
	"gitlab.com/dh-backend/search-service/internal/ports"
	"gitlab.com/grpc-buffer/proto/go/pkg/proto"
	//"gitlab.com/dh-backend/search-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

const (
	ItemIndex              = "item"
	PackageIndex           = "package"
	ItemSuggestionIndex    = "item-suggestions"
	PackageSuggestionIndex = "package-suggestions"
	QueryField             = "Name"
	MenuChannel            = "menuCreationChannel"
)

type ElasticSearchServer struct {
	proto.UnimplementedSearchServiceServer
	Elasticsearch ports.Elasticsearch
}

func Start() {

	configs := config.ReadConfigs(".") // Read configs from config file

	PORT := fmt.Sprintf(":%s", configs.GrpcPort)
	if PORT == ":" || PORT == "" {
		PORT = ":8080"
	}
	fmt.Println("PORT:", PORT)

	log.Println("elastic search url", configs.ElasticSearchUrl)

	client := elasticsearch.GetESClient(configs.ElasticSearchUrl)
	es := elasticsearch.NewElasticSearchDB(client)
	elasticServiceServer := &ElasticSearchServer{
		Elasticsearch: es,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0%v", PORT))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// grpc server
	grpcServer := grpc.NewServer()

	fmt.Println("consul address: ", configs.ConsulAddress)

	healthpb.RegisterHealthServer(grpcServer, health.NewServer())

	go func() {
		_, err := elasticServiceServer.Elasticsearch.SearchAllData(PackageSuggestionIndex)
		if err != nil {
			err := elasticServiceServer.Elasticsearch.CreateSuggestionIndex(PackageSuggestionIndex)
			if err != nil {
				log.Fatalf("Cannot create package-suggestion Index: %v", err)
			}
		}
	}()

	go func() {
		_, err = elasticServiceServer.Elasticsearch.SearchAllData(ItemSuggestionIndex)
		if err != nil {
			err := elasticServiceServer.Elasticsearch.CreateSuggestionIndex(ItemSuggestionIndex)
			if err != nil {
				log.Fatalf("Cannot create item-suggestion Index: %v", err)
			}
		}
	}()

	go func() {
		config.ServiceRegistryWithConsul("search-grpc", "search", PORT, configs.ConsulAddress, []string{"GRPC", "backend"})
		config.ServiceRegistryWithConsul("search-http", "search", ":8205", configs.ConsulAddress, []string{"HTTP", "envoy"})
	}()

	go func() {
		RabbitMQServicesForMenu()
	}()

	proto.RegisterSearchServiceServer(grpcServer, elasticServiceServer)
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
