package auth

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type contextKey string

const AuthClientContextKey = contextKey("auth_client")

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with the chitty-chat server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()
		conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		client := pb.NewAuthServiceClient(conn)

		ctx := context.WithValue(cmd.Context(), AuthClientContextKey, client)
		cmd.SetContext(ctx)
	},
}

func init() {
	AuthCmd.AddCommand(
		loginCmd,
		logoutCmd,
		tokenCmd,
	)
}
