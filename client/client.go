package main

import (
	"context"
	"fmt"
	"log"

	"grpcChatServer/chatserver"

	"google.golang.org/grpc"
)

func main() {
	serverUrl := "localhost:6000"
	log.Printf("Connecting : " + serverUrl)

	conn, err := grpc.Dial(serverUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Faile to conncet to gRPC server :: %v", err)
	}
	defer conn.Close()

	client := chatserver.NewServicesClient(conn)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}
	fmt.Println(stream)

	// blocker
	blocker := make(chan bool)
	<-blocker
}
