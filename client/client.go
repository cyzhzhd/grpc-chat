package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"grpcChatServer/chatserver"

	"google.golang.org/grpc"
)

type clientHandle struct {
	stream     chatserver.Services_ChatServiceClient
	clientName string
}

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

	ch := clientHandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	blocker := make(chan bool)
	<-blocker
}

func (ch *clientHandle) clientConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your name : ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console :: %v", err)
	}
	ch.clientName = strings.Trim(name, "\r\n")
}

func (ch *clientHandle) sendMessage() {
	for {
		reader := bufio.NewReader(os.Stdin)

		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &chatserver.FromClient{
			Name: ch.clientName,
			Body: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}

func (ch *clientHandle) receiveMessage() {
	for {
		msg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		// print message to console
		fmt.Printf("%s : %s \n", msg.Name, msg.Body)
	}
}
