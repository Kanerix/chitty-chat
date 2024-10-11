package grpc

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatServer struct {
	pb.UnimplementedChatServer
}

func (s *ChatServer) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	if len(in.Message) > 128 {
		return nil, status.Error(codes.InvalidArgument, "message is too long")
	}

	return &pb.MessageResponse{}, nil
}
