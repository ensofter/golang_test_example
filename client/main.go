package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc_test_example/pb"
)

func main() {
	grpcAddr := net.JoinHostPort("localhost", "9090")
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to grpc server: %v", err)
	}

	client := pb.NewTelephoneClient(conn)

	// runGetContact(client)

	// runListContacts(client)

	// runRecordCallHistory(client)

	runSendMessage(client)
}
