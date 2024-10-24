package cmd

import (
	"github.com/spf13/cobra"
)

type TokenKey struct{}

var Hostname string

var RootCmd = &cobra.Command{
	Use:   "chitty",
	Short: "The Chitty-Chat chat client!",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	RootCmd.PersistentFlags().StringVarP(
		&Hostname,
		"host",
		"H",
		"localhost:8080",
		"The host of the chitty-chat server",
	)

	RootCmd.AddCommand(
		chatCmd,
	)
}
