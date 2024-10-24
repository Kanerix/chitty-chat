package chat

import (
	"fmt"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	SessionStore  *session.InMemorySessionStore
	broadcastChan chan *pb.Message
	clients       map[string]*pb.Message
}

func NewChatServer(sessionStore *session.InMemorySessionStore) *ChatServer {
	return &ChatServer{
		SessionStore:  sessionStore,
		broadcastChan: make(chan *pb.Message, 10),
		clients:       make(map[string]*pb.Message),
	}
}

func (s *ChatServer) Chat(server pb.ChatService_ChatServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			return err
		}

		session, ok := session.FromContext(server.Context())
		if !ok {
			return status.Error(codes.Unauthenticated, "no active session")
		}

		name := session.Username
		if session.Anonymous {
			name = "Anonymous"
		}

		fmt.Println(name, "-", request.Message)
	}
}

func (s *ChatServer) Broadcast(message string) {
	for _, conn := range s.clients {
		fmt.Println(conn.Message)
	}
}
