package auth

import (
	"context"
	"errors"
	"unicode"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := IsValidUsername(in.Username); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	session, err := s.SessionStore.Save(in.Username, in.Anonymous)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	return &pb.LoginResponse{
		SessionToken: session.String(),
	}, nil
}

func IsValidUsername(username string) error {
	for _, l := range username {
		if !unicode.IsLetter(l) {
			return errors.New("invalid username")
		}
	}

	return nil
}
