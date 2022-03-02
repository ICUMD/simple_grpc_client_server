package main

import (
	"context"
	"log"

	"github.com/ICUMD/simple_grpc_client_server/api"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewStreamClient(conn)
	ctx := context.Background()
	response, err := c.SendData(ctx, &api.StreamingData{CurrData: "Hello there!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Message back from server: %s", response.CurrData)
}
