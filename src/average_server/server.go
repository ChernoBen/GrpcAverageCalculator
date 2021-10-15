package main

import (
	"computeAverage/src/average_proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	average_proto.UnimplementedCalculateServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v\n", err)
	}
	log.Printf("Server running...\n")
	s := grpc.NewServer()
	average_proto.RegisterCalculateServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao setar server: %v\n", err)
	}
}
