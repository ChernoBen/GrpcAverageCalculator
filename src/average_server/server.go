package main

import (
	"computeAverage/src/average_proto"
	"fmt"
	"io"
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

func (*server) ComputeAverage(stream average_proto.CalculateService_ComputeAverageServer) error {
	fmt.Println("Calcula média:")
	sum := int32(0)
	count := 0
	//obter entradas do usuario
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			//ao chegar no final do arquivo retorne o resultado(response)
			return stream.SendAndClose(&average_proto.ComputeAverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Erro de entrada: %v\n", err)
		}
		sum += req.GetNumber()
		count++
	}
}
