package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc_test_example/pb"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	baseServer := grpc.NewServer()
	pb.RegisterTelephoneServer(baseServer, NewServer())

	log.Printf("Starting server on %s", lis.Addr().String())
	baseServer.Serve(lis)
}
