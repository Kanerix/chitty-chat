package auth

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
)

func (s *AuthServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	s.SessionStorage.Delete(in.SessionToken)
	return &pb.LogoutResponse{}, nil
}
