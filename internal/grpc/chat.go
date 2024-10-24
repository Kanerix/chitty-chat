package grpc

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/kanerix/chitty-chat/internal/lamport"
	pb "github.com/kanerix/chitty-chat/proto"
)

type ChatServer struct {
	pb.UnimplementedChatServer
	mu      sync.Mutex
	clients map[string]pb.Chat_BroadcastServer
	clock   lamport.Clock
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

		if err := s.process(req, stream); err != nil {
			switch err {
			case ErrInvalidUsername:

			}
		}
	}
}

func (s *ChatServer) process(req *pb.ChatEvent, stream pb.Chat_BroadcastServer) error {
	switch evt := req.Event.(type) {
	case *pb.ChatEvent_Join:
		return s.userJoin(evt, stream)
	case *pb.ChatEvent_Leave:
		return s.userLeave(evt)
	case *pb.ChatEvent_Message:
		return s.chatMessage(evt)
	default:
		return ErrUnknownChatEvent
	}
}

func (s *ChatServer) userJoin(event *pb.ChatEvent_Join, stream pb.Chat_BroadcastServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.clients[event.Join.Username]
	if ok {
		return ErrUsernameExists
	}

	username := event.Join.Username

	if len(username) > 16 || strings.ToLower(username) == "server" {
		return ErrInvalidUsername
	}

	s.clients[username] = stream

	if err := s.broadcast(&pb.ChatEvent_ChatMessage{
		Username: "Server",
		Message:  fmt.Sprintf("Participant %s joined Chitty-Chat", username),
	}); err != nil {
		return err
	}

	return nil
}

func (s *ChatServer) userLeave(event *pb.ChatEvent_Leave) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	username := event.Leave.Username

	_, ok := s.clients[username]
	if !ok {
		return ErrUsernameNotFound
	}

	delete(s.clients, username)

	if err := s.broadcast(&pb.ChatEvent_ChatMessage{
		Username: "SERVER",
		Message:  fmt.Sprintf("Participant %s left Chitty-Chat", username),
	}); err != nil {
		return err
	}

	return nil
}

func (s *ChatServer) chatMessage(event *pb.ChatEvent_Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	message := event.Message.Message

	if len(message) > 128 {
		return ErrMessageTooLong
	}

	s.broadcast(event.Message)

	return nil
}

func (s *ChatServer) broadcast(message *pb.ChatEvent_ChatMessage) error {
	s.clock.Step()

	for _, stream := range s.clients {
		if err := stream.Send(&pb.ChatMessage{
			Timestamp: s.clock.Now(),
			Username:  message.Username,
			Message:   message.Message,
		}); err != nil {
			if err == io.EOF {
				s.userLeave(&pb.ChatEvent_Leave{
					Leave: &pb.ChatEvent_UserLeave{
						Username: message.Username,
					},
				})

				return nil
			}

			return err
		}
	}

	return nil
}

var (
	ErrInvalidUsername  = errors.New("the username is invalid")
	ErrUsernameNotFound = errors.New("the username is not found")
	ErrUsernameExists   = errors.New("the username already exists")
	ErrUnknownChatEvent = errors.New("the chat event is unknown")
	ErrMessageTooLong   = errors.New("the message must only be 128 charaters")
)
