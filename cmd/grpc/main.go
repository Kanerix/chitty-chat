package main

import (
	"log"
	"net"

	grpcChat "github.com/kanerix/chitty-chat/internal/grpc"
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

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterChatServer(s, grpcChat.NewChatServer())

	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve:", err)
	}
}
