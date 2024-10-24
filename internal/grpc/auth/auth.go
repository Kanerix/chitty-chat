package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	SessionStore *session.InMemorySessionStore
}

var NonAuthRoutes = []string{
	"/chitty_chat.AuthService/Login",
	"/grpc.health.v1.Health/",
	"/grpc.reflection.v1.ServerReflection/",
}

type streamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func AuthUnaryInterceptor(
	sessionStore *session.InMemorySessionStore,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		for _, route := range NonAuthRoutes {
			if strings.HasPrefix(info.FullMethod, route) {
				return handler(ctx, req)
			}
		}

		session_token, err := isValidContext(ctx, sessionStore)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		ctx = context.WithValue(ctx, session.SessionKey{}, session_token)

		return handler(ctx, req)
	}
}

func AuthStreamInterceptor(
	sessionStore *session.InMemorySessionStore,
) grpc.StreamServerInterceptor {
	return func(
		srv any,
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
		session_token, err := isValidContext(ctx, sessionStore)
		if err != nil {
			return status.Error(codes.Unauthenticated, err.Error())
		}

		ctx = context.WithValue(ctx, session.SessionKey{}, session_token)

		wss := &streamWrapper{ss, ctx}
		return handler(srv, wss)
	}
}

func isValidContext(
	ctx context.Context,
	sessionStore *session.InMemorySessionStore,
) (*session.Session, error) {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrUnauthorized
	}

	token, exsists := metadata["authorization"]
	if !exsists {
		return nil, ErrUnauthorized
	}

	session, err := isValidToken(token, sessionStore)
	if err != nil {
		return nil, ErrUnauthorized
	}

	return session, nil
}

func isValidToken(
	authorization []string,
	sessionStore *session.InMemorySessionStore,
) (*session.Session, error) {
	if len(authorization) != 1 {
		return nil, ErrUnauthorized
	}

	token := strings.Trim(authorization[0], " ")
	sessionFromString, err := session.StringToSession(token)
	if err != nil {
		return nil, err
	}

	session, err := sessionStore.Get(sessionFromString.Username)
	if err != nil {
		return nil, err
	}

	return session, nil
}

var (
	ErrUnauthorized = errors.New("unauthorized for resource")
)
