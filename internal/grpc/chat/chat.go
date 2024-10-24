package chat

import (
	"fmt"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	SessionStore *session.InMemorySessionStore
	participants []*pb.ChatService_ChatServer
}

func (s *ChatServer) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		ctx := stream.Context()
		session := ctx.Value(session.SessionKey{}).(*session.Session)

		name := session.Username
		if session.Anonymous {
			name = "Anonymous"
		}

		message := fmt.Sprintf("%s - %s @ %s", "TIMESTAMP", name, msg.Message)
		err = stream.Send(&pb.Message{Message: message})
		if err != nil {
			return err
		}
	}
}

func (s *ChatServer) Broadcast(msg string) error {
	s.SessionStore.RLock()
	defer s.SessionStore.RUnlock()

	return nil
}
