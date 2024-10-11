package chat

import (
	"fmt"

	"github.com/kanerix/chitty-chat/internal/grpc/auth"
	pb "github.com/kanerix/chitty-chat/proto"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServer) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		ctx := stream.Context()
		session := ctx.Value(auth.SessionContextKey).(auth.Session)
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
