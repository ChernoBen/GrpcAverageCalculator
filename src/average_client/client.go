package main

import (
	"computeAverage/src/average_proto"
	"context"
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
	doClientStreaming(c)
}

func doClientStreaming(c average_proto.CalculateServiceClient) {
	fmt.Println("Consulta media")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Erro na requisição: %v \n", err)
	}
	numbers := []int32{1, 2, 3, 4, 5, 6}
	for _, num := range numbers {
		//enviando dados ao servidor
		stream.Send(&average_proto.ComputeAverageRequest{
			Number: num,
		})
	}
	//fechando stream e obtendo retorno
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Erro na resposta: %v\n", err)
	}
	fmt.Printf("A média é: %v\n", res.GetAverage())
}
