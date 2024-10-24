package main

import (
	"os"

	"github.com/kanerix/chitty-chat/internal/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableTraverseRunHooks = true
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
