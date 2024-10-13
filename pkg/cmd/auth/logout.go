package auth

import (
	"context"
	"time"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:     "logout",
	Short:   "Logout of the chitty-chat server",
	Example: "chitty auth logout",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		client, ok := cmd.Context().Value(AuthClientContextKey).(pb.AuthServiceClient)
		if !ok {
			cmd.PrintErr("could not get auth client")
			return
		}

		token, ok := cmd.Context().Value(session.SessionContextKey).(string)
		if !ok {
			cmd.PrintErr("could not get session token")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := client.Logout(ctx, &pb.LogoutRequest{
			SessionToken: token,
		})
		if err != nil {
			cmd.PrintErr(err)
			return
		}
	},
}
