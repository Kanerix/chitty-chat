package chat

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type contextKey string

const ChatClientContextKey = contextKey("chat_client")

var ChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with other users on the chitty-chat server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.PersistentPreRun != nil {
			cmd.PersistentPreRun(cmd, args)
		}

		host := cmd.Flag("host").Value.String()
		conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		client := pb.NewChatServiceClient(conn)

		ctx := context.WithValue(cmd.Context(), ChatClientContextKey, client)
		cmd.SetContext(ctx)
	},
}

func init() {
	ChatCmd.AddCommand(
		joinCmd,
		listCmd,
	)
}
