package main

import (
	"net"
	"os"
	"time"

	grpcAuth "github.com/kanerix/chitty-chat/internal/grpc/auth"
	grpcChat "github.com/kanerix/chitty-chat/internal/grpc/chat"
	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatal().Msgf("failed to listen: %s", err)
	}
	defer listener.Close()

	logger.Info().Msgf("server is listening on %s", listener.Addr().String())

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcAuth.AuthInterceptor))
	reflection.Register(s)

	pb.RegisterAuthServer(s, &grpcAuth.AuthServer{})
	pb.RegisterChatServer(s, &grpcChat.ChatServer{})

	if err := s.Serve(listener); err != nil {
		logger.Fatal().Msgf("failed to serve: %s", err)
	}
}
