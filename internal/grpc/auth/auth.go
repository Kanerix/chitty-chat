package grpc

import (
	"context"
	"strings"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var storage = &InMemorySessionStore{}

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func AuthInterceptor(
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

func isValid(authorization []string) bool {
	if len(authorization) != 1 {
		return false
	}

	token := strings.TrimPrefix(authorization[0], "Bearer ")

	return token == "valid-token"
}
