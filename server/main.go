package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ICUMD/simple_grpc_client_server/api"
	"google.golang.org/grpc"
)

func main() {
	// main start a gRPC server and waits for connection
	// create a server instance
	s := api.Server{}
	lis, error := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if error != nil {
		log.Fatalf("Failed to listen: %v", error)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterStreamServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
