package main

import (
	"context"
	"log"

	"github.com/ICUMD/simple_grpc_client_server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Authentication holds the login/password
type Authentication struct {
	Login    string
	Password string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	var conn *grpc.ClientConn
	// Create the client TLS credentials
	creds, error := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	if error != nil {
		log.Fatalf("could not load tls cert: %s", error)
	}
	auth := Authentication{
		Login:    "John",
		Password: "Doe",
	}
	conn, err := grpc.Dial(":7777", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewCommunicationClient(conn)
	ctx := context.Background()
	response, err := c.SendData(ctx, &api.ExchangeData{CurrData: "Hello there!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Message back from server: %s", response.CurrData)
}
