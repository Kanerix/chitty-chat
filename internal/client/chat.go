package client

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kanerix/chitty-chat/internal/lamport"
	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc"
)

type ChatClient struct {
	pb.ChatClient
}

type BroadcastStream struct {
	pb.Chat_BroadcastClient
	clock lamport.Clock
}

type MessageRecvEvent struct {
	Timestamp uint64
	Username  string
	Message   string
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
		lamport.Clock{},
	}, nil
}

func (bs *BroadcastStream) JoinChat(username string) error {
	bs.clock.Step()

	if err := bs.Send(&pb.ChatEvent{
		Timestamp: bs.clock.Now(),
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
	bs.clock.Step()

	if err := bs.Send(&pb.ChatEvent{
		Timestamp: bs.clock.Now(),
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
	bs.clock.Step()

	if err := bs.Send(&pb.ChatEvent{
		Timestamp: bs.clock.Now(),
		Event: &pb.ChatEvent_Message{
			Message: &pb.ChatEvent_ChatMessage{
				Username: username,
				Message:  message,
			},
		},
	}); err != nil {
		return err
	}

	return nil
}

func (bs *BroadcastStream) MessageListener(p *tea.Program) error {
	for {
		req, err := bs.Recv()
		if err != nil {
			log.Fatal(err.Error())
		}

		p.Send(MessageRecvEvent{
			Timestamp: req.Timestamp,
			Username:  req.Username,
			Message:   req.Message,
		})
	}
}
