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
	RunE: func(cmd *cobra.Command, args []string) error {
		client := cmd.Context().Value(AuthClientKey{}).(pb.AuthServiceClient)

		token, ok := cmd.Context().Value(session.SessionKey{}).(string)
		if !ok {
			return session.ErrSessionKeyNotFound
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := client.Logout(ctx, &pb.LogoutRequest{
			SessionToken: token,
		})
		if err != nil {
			return err
		}

		return nil
	},
}
