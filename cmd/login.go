package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login <username>",
	Args:  cobra.ExactArgs(1),
	Short: "Login as an user with a username",
	Long: `Login with a username to start chatting with others.
The username must be unique and not already taken.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
