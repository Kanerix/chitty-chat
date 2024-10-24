package grpc

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	pb "github.com/kanerix/chitty-chat/proto"
)

type ChatServer struct {
	pb.UnimplementedChatServer
	mu      sync.Mutex
	clients map[string]pb.Chat_BroadcastServer
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[string]pb.Chat_BroadcastServer),
	}
}

func (s *ChatServer) Broadcast(stream pb.Chat_BroadcastServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		log.Println("Recived message")

		switch evt := req.Event.(type) {
		case *pb.ChatEvent_Join:
			s.userJoin(evt, stream)
		case *pb.ChatEvent_Leave:
			s.userLeave(evt)
		case *pb.ChatEvent_Message:
			s.chatMessage(evt)
		default:
			return errors.New("unkown event for broadcast")
		}
	}
}

func (s *ChatServer) userJoin(event *pb.ChatEvent_Join, stream pb.Chat_BroadcastServer) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.clients[event.Join.Username]
	if ok {
		log.Fatalln("The username is already taken")
		// TODO: Return ErrUsernameTaken
		return
	}

	username := event.Join.Username

	if strings.ToLower(username) == "server" {
		log.Fatalln("The username is invalid")
		// TODO: Return ErrInvalidUsername
		return
	}

	log.Println("Stored username in clients")
	s.clients[username] = stream

	s.broadcast(&pb.ChatEvent_ChatMessage{
		Username: "Server",
		Message:  fmt.Sprintf("Participant %s joined Chitty-Chat", event.Join.Username),
	})
}

func (s *ChatServer) userLeave(event *pb.ChatEvent_Leave) {
	s.mu.Lock()
	defer s.mu.Unlock()

	username := event.Leave.Username

	_, ok := s.clients[username]
	if !ok {
		log.Fatalln("Username not found in clients")
		// TODO: Return ErrUserNameNotFound
		return
	}

	log.Println("Deleted username from clients")
	delete(s.clients, username)

	s.broadcast(&pb.ChatEvent_ChatMessage{
		Username: "Server",
		Message:  fmt.Sprintf("Participant %s left Chitty-Chat", username),
	})
}

func (s *ChatServer) chatMessage(event *pb.ChatEvent_Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.broadcast(event.Message)
}

func (s *ChatServer) broadcast(message *pb.ChatEvent_ChatMessage) {
	for username, stream := range s.clients {
		log.Printf("Sending message to %s", username)
		if err := stream.Send(&pb.ChatMessage{
			Timestamp: 0, // TODO: Lamport
			Username:  message.Username,
			Message:   message.Message,
		}); err != nil {
			log.Printf("Error sending message to %s: %v", username, err)
		}
	}
}
