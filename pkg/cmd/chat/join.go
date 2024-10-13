package chat

import (
	"context"
	"time"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var joinCmd = &cobra.Command{
	Use:     "join",
	Short:   "Join the chat on the chitty-chat server",
	Example: "chitty chat join -H localhost:8080",
	Run: func(cmd *cobra.Command, args []string) {
		client, ok := cmd.Context().Value(ChatClientContextKey).(pb.ChatServiceClient)
		if !ok {
			cmd.PrintErr("could not get chat client")
			return
		}

		token, ok := cmd.Context().Value(session.SessionContextKey).(string)
		if !ok {
			cmd.PrintErr("could not get session token")
			return
		}

		md := metadata.Pairs("authorization", token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		recv, err := client.Chat(ctx)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		for {
			msg, err := recv.Recv()
			if err != nil {
				cmd.PrintErr(err)
				return
			}

			cmd.Println(msg)
		}
	},
}
