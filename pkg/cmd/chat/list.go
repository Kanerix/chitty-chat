package chat

import (
	"context"
	"time"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all members currently on the chitty-chat server",
	Example: "chitty chat list 1",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		client, ok := cmd.Context().Value(ChatClientContextKey).(pb.ChatServiceClient)
		if !ok {
			cmd.PrintErr("could not get chat client")
			return
		}

		token := cmd.Context().Value(session.SessionContextKey).(string)
		md := metadata.Pairs("authorization", token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		client.ListMembers(ctx, &pb.ListMembersRequest{
			Page: 0,
		})
	},
}

func init() {
	listCmd.Flags().StringP("page", "p", "0", "The page number of the users")
}
