package chat

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ChatServer) ListMembers(ctx context.Context, in *pb.ListMembersRequest) (*pb.ListMembersResponse, error) {
	if in.Page < 1 {
		return nil, status.Error(codes.InvalidArgument, "page can't be less than 0")
	}

	s.SessionStore.RLock()
	usernames := s.SessionStore.List(int(in.Page))
	s.SessionStore.RUnlock()

	if len(usernames) == 0 {
		return nil, status.Error(codes.NotFound, "no members found on this page")
	}

	return &pb.ListMembersResponse{
		Members: usernames,
	}, nil
}
