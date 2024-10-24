package auth

import (
	"context"

	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClientKey struct{}

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with the chitty-chat server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		host := cmd.Flag("host").Value.String()
		conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		client := pb.NewAuthServiceClient(conn)

		ctx := context.WithValue(cmd.Context(), AuthClientKey{}, client)
		cmd.SetContext(ctx)

		return nil
	},
}

func init() {
	AuthCmd.AddCommand(
		loginCmd,
		logoutCmd,
		tokenCmd,
	)
}
