package cmd

import (
	"context"

	"github.com/kanerix/chitty-chat/pkg/cmd/auth"
	"github.com/kanerix/chitty-chat/pkg/cmd/chat"
	"github.com/kanerix/chitty-chat/pkg/session"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:              "chitty",
	Short:            "The Chitty-Chat chat client!",
	Args:             cobra.MinimumNArgs(1),
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		session_token, err := getToken(cmd)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		ctx := context.WithValue(cmd.Context(), session.SessionContextKey, session_token)
		cmd.SetContext(ctx)
	},
}

func init() {
	RootCmd.AddCommand(
		auth.AuthCmd,
		chat.ChatCmd,
	)
	RootCmd.PersistentFlags().StringP("token", "t", "", "The token used for authentication")
	RootCmd.PersistentFlags().StringP("host", "H", "localhost:8080", "The host of the chitty-chat server")
}

func getToken(cmd *cobra.Command) (string, error) {
	token, err := cmd.Flags().GetString("token")
	if err != nil || token == "" {
		token, err := session.GetSessionFileContent()
		if err != nil {
			return "", err
		}

		return string(token), nil
	}

	return token, nil
}
