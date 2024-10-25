package mvc

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kanerix/chitty-chat/internal/client"
)

type model struct {
	chat     ChatView
	input    InputArea
	notify   NotifyView
	stream   *client.BroadcastStream
	username string
}

func NewChatModel(stream *client.BroadcastStream, username string) model {
	stream.JoinChat(username)

	model := model{
		chat:     NewChatView(),
		input:    NewInputArea(),
		notify:   NewNotifyView(),
		username: username,
		stream:   stream,
	}

	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.chat.Width = msg.Width
		m.input.SetWidth(msg.Width)
		m.notify.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.stream.LeaveChat(m.username)
			return m, tea.Quit

		case tea.KeyEnter:
			if err := m.sendMessage(); err != nil {
				m.notify.NotifyErr(err)
			} else {
				m.input.Reset()
			}
			return m, nil

		default:
			var cmd tea.Cmd
			m.input.Model, cmd = m.input.Model.Update(msg)
			return m, cmd
		}

	case cursor.BlinkMsg:
		var cmd tea.Cmd
		m.input.Model, cmd = m.input.Model.Update(msg)
		return m, cmd

	case MessageRecvEvent:
		m.chat.AppendMessage(&msg.Message)
		m.chat.RenderMessages()
		return m, nil

	default:
		return m, nil
	}
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n%s",
		m.chat.View(),
		m.input.View(),
		m.notify.View(),
	)
}

func (m *model) sendMessage() error {
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

	return nil
}
