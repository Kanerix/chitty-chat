package main

import (
	"fmt"
	"os"

	"github.com/kanerix/chitty-chat/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableTraverseRunHooks = true
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
