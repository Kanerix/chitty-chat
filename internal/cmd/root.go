package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var Hostname string

var RootCmd = &cobra.Command{
	Use:   "chitty",
	Short: "The Chitty-Chat chat client!",
	Args:  cobra.MinimumNArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer f.Close()

		log.SetFlags(0)
		log.SetOutput(f)

		return nil
	},
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
