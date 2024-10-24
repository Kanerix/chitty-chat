package auth

import (
	"github.com/kanerix/chitty-chat/pkg/session"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:     "token",
	Short:   "Shows the current session token",
	Example: "chitty auth token",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		token, ok := ctx.Value(session.SessionKey{}).(string)
		if !ok {
			return session.ErrSessionKeyNotFound
		}

		cmd.Println(token)

		return nil
	},
}
