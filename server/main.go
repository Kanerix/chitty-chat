package main

import (
	"net"
	"os"
	"time"

	"github.com/kanerix/chitty-chat/pb"
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

	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	reflection.Register(s)

	pb.RegisterChatServer(s, &chatServer{})
	pb.RegisterAuthServer(s, &authServer{})
	if err := s.Serve(listener); err != nil {
		logger.Fatal().Msgf("failed to serve: %s", err)
	}

	s.Serve(listener)
}
