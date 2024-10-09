package main

import (
	"context"

	"github.com/kanerix/chitty-chat/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type chatServer struct {
	pb.UnimplementedChatServer
}

func (s *chatServer) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	if len(in.Message) > 128 {
		return nil, status.Error(codes.InvalidArgument, "message is too long")
	}

	return &pb.MessageResponse{}, nil
}
