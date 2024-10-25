package client

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc"
)

type ChatClient struct {
	pb.ChatClient
}

type BroadcastStream struct {
	pb.Chat_BroadcastClient
}

func NewChatClient(conn *grpc.ClientConn) *ChatClient {
	return &ChatClient{
		pb.NewChatClient(conn),
	}
}

func (cc *ChatClient) Stream(ctx context.Context) (*BroadcastStream, error) {
	stream, err := cc.Broadcast(ctx)
	if err != nil {
		return nil, err
	}

	return &BroadcastStream{
		stream,
	}, nil
}

func (bs *BroadcastStream) JoinChat(username string) error {
	if err := bs.Send(&pb.ChatEvent{
		Event: &pb.ChatEvent_Join{
			Join: &pb.ChatEvent_UserJoin{
				Username: username,
			},
		},
	}); err != nil {
		return err
	}

	return nil
}

func (bs *BroadcastStream) LeaveChat(username string) error {
	if err := bs.Send(&pb.ChatEvent{
		Event: &pb.ChatEvent_Leave{
			Leave: &pb.ChatEvent_UserLeave{
				Username: username,
			},
		},
	}); err != nil {
		return err
	}

	return nil
}

func (bs *BroadcastStream) SendMessage(username string, message string) error {
	if err := bs.Send(&pb.ChatEvent{
		Event: &pb.ChatEvent_Message{
			Message: &pb.ChatEvent_ChatMessage{
				Username: username,
			},
		},
	}); err != nil {
		return err
	}

	return nil
}
