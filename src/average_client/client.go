package main

import (
	"computeAverage/src/average_proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Falha ao conectar com servidor: %v\n", err)
	}
	defer conn.Close()
	c := average_proto.NewCalculateServiceClient(conn)
	if c != nil {
		fmt.Println("Conexão com servidor criada com sucesso!")
	}
}
