package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chitty-chat <username>",
	Short: "A simple chat application using gRPC",
	Long: `Chitty-chat is a chat service that allows users
	to connect, send messages, and leave a chat room.
		
	Chitty-chat is built using gRPC and Protocol Buffers.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
