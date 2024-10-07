package main

import (
	pb "go-store/discount/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const ElasticAddress = "http://localhost:9200"

type server struct {
	pb.DiscountProtoServiceServer
}

func main() {
	//hook, err := logging.NewElasticHook([]string{ElasticAddress})
	//
	//if err != nil {
	//	log.Fatalf("Error creating hook: %v", err)
	//}
	//
	//server.Logger = logrus.New()
	//server.Logger.AddHook(hook)
	//defer hook.Close()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDiscountProtoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
