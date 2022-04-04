package chatserver

import (
	"fmt"
)

type ChatServer struct {
	UnimplementedServicesServer
}

func (cs *ChatServer) ChatService(css Services_ChatServiceServer) error {
	fmt.Println("ChatServce user connected")

	errch := make(chan error)

	return <-errch
}
