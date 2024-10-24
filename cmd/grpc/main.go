package main

import (
	"log"
	"net"

	grpcAuth "github.com/kanerix/chitty-chat/internal/grpc/auth"
	grpcChat "github.com/kanerix/chitty-chat/internal/grpc/chat"
	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("failed to listen:", err)
	}
	defer listener.Close()

	log.Println("server is listening on", listener.Addr().String())

	var sessionStore = session.NewInMemorySessionStore()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcAuth.AuthUnaryInterceptor(sessionStore)),
		grpc.StreamInterceptor(grpcAuth.AuthStreamInterceptor(sessionStore)),
	)
	reflection.Register(s)

	pb.RegisterAuthServiceServer(s, grpcAuth.NewAuthServer(sessionStore))
	pb.RegisterChatServiceServer(s, grpcChat.NewChatServer(sessionStore))

	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve:", err)
	}
}
