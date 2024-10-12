package cmd

import (
	"context"
	"strings"

	"github.com/kanerix/chitty-chat/pkg/cmd/auth"
	"github.com/kanerix/chitty-chat/pkg/cmd/chat"
	"github.com/kanerix/chitty-chat/pkg/util"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:              "chitty",
	Short:            "The Chitty-Chat chat client!",
	Args:             cobra.MinimumNArgs(1),
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		content, err := util.GetSessionFileContent()
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		session_token := strings.ToValidUTF8(string(content), "")

		ctx := context.WithValue(cmd.Context(), util.SessionContextKey, session_token)
		cmd.Println("key", ctx.Value(util.SessionContextKey))
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
