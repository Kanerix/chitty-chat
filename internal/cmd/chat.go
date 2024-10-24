package cmd

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var username string

var chatCmd = &cobra.Command{
	Use:     "chat",
	Short:   "Join the chat on the Chitty-chat server",
	Example: "chitty chat -H chitty.lerpz.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.NewClient(Hostname, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		defer conn.Close()

		client := pb.NewChatClient(conn)

		stream, err := client.Broadcast(context.Background())
		if err != nil {
			return err
		}

		joinChat(stream)

		go messageListener(stream)

		<-stream.Context().Done()
		return nil
	},
}

func init() {
	chatCmd.Flags().StringVarP(
		&username,
		"username",
		"u",
		"",
		"The username to use",
	)
	chatCmd.MarkFlagRequired("username")
}

func joinChat(stream pb.Chat_BroadcastClient) {
	stream.Send(&pb.ChatEvent{
		Event: &pb.ChatEvent_Join{
			Join: &pb.ChatEvent_UserJoin{
				Username: username,
			},
		},
	})
}

func messageListener(stream pb.Chat_BroadcastClient) {
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Printf("%d @ %s - %s\n", req.Timestamp, req.Username, req.Message)
	}
}
