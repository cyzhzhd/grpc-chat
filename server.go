package main

import (
	"log"
	"net"

	"grpcChatServer/chatserver"

	"google.golang.org/grpc"
)

func main() {
	Port := "6000"

	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", Port, err)
	}
	log.Println("Listening @ : " + Port)

	grpcserver := grpc.NewServer()

	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcserver, &cs)

	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
