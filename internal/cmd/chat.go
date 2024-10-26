package cmd

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kanerix/chitty-chat/internal/client"
	"github.com/kanerix/chitty-chat/internal/mvc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var username string

var chatCmd = &cobra.Command{
	Use:     "chat",
	Short:   "Join the chat on the Chitty-chat server",
	Example: "chitty chat -H chitty.lerpz.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.SetFlags(0)

		conn, err := grpc.NewClient(Hostname, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		client := client.NewChatClient(conn)
		stream, err := client.Stream(context.Background())
		if err != nil {
			log.Fatalln(err.Error())
		}

		p := tea.NewProgram(mvc.NewChatModel(stream, username), tea.WithAltScreen())
		go stream.MessageListener(p)
		if _, err := p.Run(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	chatCmd.Flags().StringVarP(&username, "username", "u", "", "Username for Chitty-Chat")
	chatCmd.MarkFlagRequired("test")
}
