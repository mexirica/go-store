package main

import (
	pb "go-store/discount/pkg/grpc"
	"go-store/discount/internal/server"
	"go-store/pkg/logging"
	"log"
	"net"

	"google.golang.org/grpc"
)

const ElasticAddress = "http://localhost:9200"

func main() {
	hook, err := logging.NewElasticHook([]string{ElasticAddress})
	if err != nil {
		log.Fatalf("Error creating hook: %v", err)
	}
	
	logger,err := logging.NewLogger("logs/log.json")
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	
	logger.AddHook(hook)
	defer hook.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDiscountProtoServiceServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
