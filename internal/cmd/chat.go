package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
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
		conn, err := grpc.NewClient(Hostname, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		p := tea.NewProgram(mvc.NewChatModel(conn, username), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			return err
		}

		return nil
	},
}
