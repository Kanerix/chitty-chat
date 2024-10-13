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
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		token, ok := ctx.Value(session.SessionKey{}).(string)
		if !ok {
			cmd.Println("No token found")
			return
		}

		cmd.Println(token)
	},
}
