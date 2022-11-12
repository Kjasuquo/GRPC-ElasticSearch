package services

import (
	"fmt"
	"gitlab.com/dh-backend/search-service/config"
	"gitlab.com/dh-backend/search-service/internal/elasticsearch"
	"gitlab.com/dh-backend/search-service/internal/ports"
	"gitlab.com/grpc-buffer/proto/go/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

type ElasticSearchServer struct {
	proto.UnimplementedMenuServiceServer
	Elasticsearch ports.Elasticsearch
}

func Start() {

	configs := config.ReadConfigs(".") // Read configs from config file

	PORT := fmt.Sprintf(":%s", configs.GrpcPort)
	if PORT == ":" || PORT == "" {
		PORT = ":8080"
	}
	fmt.Println("PORT:", PORT)

	client := elasticsearch.GetESClient()
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
		config.ServiceRegistryWithConsul("search-grpc", "search", PORT, configs.ConsulAddress, []string{"GRPC", "backend"})
		config.ServiceRegistryWithConsul("search-http", "search", ":8205", configs.ConsulAddress, []string{"HTTP", "envoy"})
	}()

	RabbitMQServicesForMenu()

	proto.RegisterMenuServiceServer(grpcServer, elasticServiceServer)
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
