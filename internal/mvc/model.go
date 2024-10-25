package mvc

import (
	"context"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kanerix/chitty-chat/internal/client"
	"google.golang.org/grpc"
)

type model struct {
	chat     ChatView
	input    InputArea
	notify   NotifyView
	username string
	stream   *client.BroadcastStream
}

func NewChatModel(conn *grpc.ClientConn, username string) model {
	client := client.NewChatClient(conn)
	stream, err := client.Stream(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	return model{
		chat:     NewChatView(),
		input:    NewInputArea(),
		notify:   NewNotifyView(),
		username: username,
		stream:   stream,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.chat.Width = msg.Width
		m.input.SetWidth(msg.Width)
		m.notify.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if err := m.sendMessage(); err != nil {
				m.notify.NotifyErr(err)
			}

			return m, nil

		default:
			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			return m, cmd
		}

	case cursor.BlinkMsg:
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n%s\n\n",
		m.chat.View(),
		m.input.View(),
		m.notify.View(),
	)
}

func (m model) sendMessage() error {
	message := m.input.Value()

	if message == "" || len(message) > 128 {
		return ErrInvalidMessage
	}

	if m.stream == nil {
		return ErrStreamClosed
	}

	if err := m.stream.SendMessage(m.username, message); err != nil {
		return err
	}

	m.chat.SetContent("test")
	m.input.Reset()
	m.chat.GotoBottom()

	return nil
}
