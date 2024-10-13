package auth

import (
	"context"
	"time"

	"github.com/kanerix/chitty-chat/pkg/session"
	pb "github.com/kanerix/chitty-chat/proto"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:     "login -u [username]",
	Short:   "Login to the chitty-chat server",
	Example: "chitty auth login -u username",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		client, ok := cmd.Context().Value(AuthClientKey{}).(pb.AuthServiceClient)
		if !ok {
			cmd.PrintErr("could not get auth client")
			return
		}

		username, _ := cmd.Flags().GetString("username")
		anonymous, _ := cmd.Flags().GetBool("anonymous")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		res, err := client.Login(ctx, &pb.LoginRequest{
			Anonymous: anonymous,
			Username:  username,
		})
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		if err := session.SaveSessionToken(res.SessionToken); err != nil {
			cmd.PrintErr(err)
			return
		}

		if show, _ := cmd.Flags().GetBool("show"); show {
			cmd.Println(res.SessionToken)
		}
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Username for login")
	loginCmd.Flags().BoolP("show", "s", false, "Output the token after the login succeeds")
	loginCmd.Flags().BoolP("anonymous", "A", false, "Login as an anonymous user")
	loginCmd.MarkFlagRequired("username")
}
