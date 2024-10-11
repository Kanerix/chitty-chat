package auth

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	session, err := s.SessionStorage.Save(in.Username, in.Anonymous)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	return &pb.LoginResponse{
		SessionToken: session.String(),
	}, nil
}
