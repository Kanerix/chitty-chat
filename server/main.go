package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/kanerix/chitty-chat/pb"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedChatServer
}

func (s *Server) SendMessage(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	if len(in.Message) > 128 {
		return nil, status.Error(codes.InvalidArgument, "message is too long")
	}

	return &pb.Message{
		Message: in.Message,
	}, nil
}

func main() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatal().Msgf("failed to listen: %s", err)
	}

	logger.Info().Msgf("server is listening on %s", listener.Addr().String())

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterChatServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		logger.Fatal().Msgf("failed to serve: %s", err)
	}

	s.Serve(listener)
}
