package main

import (
	"context"
	"strings"

	"github.com/kanerix/chitty-chat/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type authServer struct {
	pb.UnimplementedAuthServer
}

func (s *authServer) JoinChannel(ctx context.Context, in *pb.JoinRequest) (*pb.JoinResponse, error) {
	return &pb.JoinResponse{
		SessionToken: "valid-token",
	}, nil
}

func isValid(authorization []string) bool {
	if len(authorization) != 1 {
		return false
	}

	token := strings.TrimPrefix(authorization[0], "Bearer ")

	return token == "valid-token"
}

func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if strings.HasPrefix(info.FullMethod, "/chitty_chat.Auth/") {
		return handler(ctx, req)
	}

	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Missing metadata")
	}

	token, exsists := metadata["authorization"]
	if !exsists || !isValid(token) {
		return nil, status.Error(codes.Unauthenticated, "Invalid session token")
	}

	return handler(ctx, req)
}
