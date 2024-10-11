package auth

import (
	"context"
	"errors"
	"strings"

	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	SessionStorage *InMemorySessionStore
	pb.UnimplementedAuthServiceServer
}

type contextKey string

const SessionContextKey = contextKey("session")

var NonAuthRoutes = []string{
	"/chitty_chat.AuthService/",
	"/grpc.health.v1.Health/",
	"/grpc.reflection.v1.ServerReflection/",
}

func AuthUnaryInterceptor(
	sessionStore *InMemorySessionStore,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		for _, route := range NonAuthRoutes {
			if strings.HasPrefix(info.FullMethod, route) {
				return handler(ctx, req)
			}
		}

		session, err := isValidContext(ctx, sessionStore)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		ctx = context.WithValue(ctx, SessionContextKey, session)

		return handler(ctx, req)
	}
}

func AuthStreamInterceptor(
	sessionStore *InMemorySessionStore,
) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		for _, route := range NonAuthRoutes {
			if strings.HasPrefix(info.FullMethod, route) {
				return handler(srv, ss)
			}
		}

		ctx := ss.Context()
		_, err := isValidContext(ctx, sessionStore)
		if err != nil {
			return err
		}

		return handler(srv, ss)
	}
}

func isValidContext(ctx context.Context, sessionStore *InMemorySessionStore) (*Session, error) {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	token, exsists := metadata["authorization"]
	if !exsists {
		return nil, errors.New("session token not found")
	}

	session, err := isValidToken(token, sessionStore)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func isValidToken(authorization []string, sessionStore *InMemorySessionStore) (*Session, error) {
	if len(authorization) != 1 {
		return nil, errors.New("invalid authorization header")
	}

	token := strings.TrimPrefix(authorization[0], "Bearer ")
	sessionFromString, err := StringToSession(token)
	if err != nil {
		return nil, err
	}

	session, err := sessionStore.Get(sessionFromString.Username)
	if err != nil {
		return nil, err
	}

	return session, nil
}
