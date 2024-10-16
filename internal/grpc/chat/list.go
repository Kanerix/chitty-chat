package chat

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ChatServer) ListMembers(ctx context.Context, in *pb.ListMembersRequest) (*pb.ListMembersResponse, error) {
	usernames := s.SessionStore.List(int(in.Page))

	if len(usernames) == 0 {
		return nil, status.Error(codes.NotFound, "No members found on this page")
	}

	return &pb.ListMembersResponse{
		Members: usernames,
	}, nil
}
