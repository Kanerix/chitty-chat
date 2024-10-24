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
	Example: "chitty chat list",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := cmd.Context().Value(ChatClientKey{}).(pb.ChatServiceClient)

		token := cmd.Context().Value(session.SessionKey{}).(string)
		md := metadata.Pairs("authorization", token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		page, err := cmd.Flags().GetInt32("page")
		if err != nil {
			return err
		}

		res, err := client.ListMembers(ctx, &pb.ListMembersRequest{Page: page})
		if err != nil {
			return err
		}

		for _, member := range res.Members {
			cmd.Println(member)
		}

		return nil
	},
}

func init() {
	listCmd.Flags().Int32P("page", "p", 1, "The page number of the users")
}
