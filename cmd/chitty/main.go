package main

import (
	"os"

	"github.com/kanerix/chitty-chat/internal/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
