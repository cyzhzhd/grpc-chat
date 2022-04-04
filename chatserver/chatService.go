package chatserver

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

type User struct {
	Css Services_ChatServiceServer
	Uid int
}
type ChatServer struct {
	UnimplementedServicesServer
	Users   []User
	usersMu sync.Mutex
}

func (cs *ChatServer) ChatService(css Services_ChatServiceServer) error {
	fmt.Println("ChatServce user connected")
	newUser := User{Css: css, Uid: rand.Intn(1e6)}

	cs.usersMu.Lock()
	cs.Users = append(cs.Users, newUser)
	cs.usersMu.Unlock()

	errch := make(chan error)

	go receiveFromStream(cs, &newUser, errch)

	return <-errch
}

func receiveFromStream(cs *ChatServer, user *User, errch_ chan error) {
	for {
		msg, err := user.Css.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
			errch_ <- err
		} else {

			for _, user := range cs.Users {
				user.Css.Send(&FromServer{Name: msg.Name, Body: msg.Body})
			}
		}
	}
}
