package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var joinCmd = &cobra.Command{
	Use:     "join",
	Short:   "Join the chat on the chitty-chat server",
	Example: "chitty chat join -H localhost:8080",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := cmd.Context().Value(ChatClientKey{}).(pb.ChatServiceClient)

		token, ok := cmd.Context().Value(session.SessionKey{}).(string)
		if !ok {
			return session.ErrSessionKeyNotFound
		}

		md := metadata.Pairs("authorization", token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		stream, err := client.Chat(ctx)
		if err != nil {
			return err
		}

		stream.Send(&pb.Message{
			Message: "I joined the channel",
		})

		for {
			log.Println("Waiting for message")
			msg, err := stream.Recv()
			if err != nil {
				return err
			}

			stream.Send(&pb.Message{
				Message: "I received your message",
			})

			fmt.Println(msg)
		}
	},
}