package main

import (
	"flag"
	"fmt"

	"github.com/kanerix/chitty-chat/pb"
	"github.com/spf13/cobra"
)

type Client struct {
	pb.ChatClient
}

func (c *Client) Connect(user *pb.User) (*pb.Message, error) {
	return nil, nil
}

var rootCmd = &cobra.Command{
	Use:   "chitty-chat <username>",
	Short: "A simple chat application using gRPC",
	Long: `Chitty-chat is a chat service that allows users
	to connect, send messages, and leave a chat room.
		
	Chitty-chat is built using gRPC and Protocol Buffers.`,
}

func main() {
	user := flag.String("user", "Anonymous", "The username of the user to connect as.")
	flag.Parse()
	fmt.Println("Hello,", *user)
}
