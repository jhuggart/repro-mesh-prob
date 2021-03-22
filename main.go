package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jhuggart/repro-mesh-prob/gen/go/customers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jhuggart/repro-mesh-prob/internal/server"
)

var port = 9090

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	customers.RegisterCustomersServer(grpcServer, server.NewServer())

	//NOTE: This is very important. graphql-mesh relies on reflection to figure out what endpoints are available
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to start server", err)
	}
}
