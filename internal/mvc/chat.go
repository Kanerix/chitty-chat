package mvc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type ChatView struct {
	viewport.Model
	messages []*ChatMessage
}

type ChatMessage struct {
	Timestamp uint64
	Username  string
	Message   string
}

var (
	usernameStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#36D136"))
	timestampStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EB34C6"))
)

func NewChatView() ChatView {
	view := ChatView{
		Model:    viewport.New(Width, 10),
		messages: []*ChatMessage{},
	}

	return view
}

func (cm *ChatView) AppendMessage(message *ChatMessage) {
	cm.messages = append(cm.messages, message)
}

func (cm *ChatView) chatMessagesToString() string {
	start := len(cm.messages) - 10
	if start < 0 {
		start = 0
	}

	message := []string{}
	for _, msg := range cm.messages[start:] {
		message = append(message, msg.String())
	}

	return strings.Join(message, "\n")
}

func (cm *ChatView) RenderMessages() {
	messages := cm.chatMessagesToString()
	cm.SetContent(messages)
}

func (cm ChatMessage) String() string {
	timestamp := fmt.Sprintf("L[%s]", strconv.FormatUint(cm.Timestamp, 10))
	username := cm.Username
	message := cm.Message

	return timestampStyle.Render(timestamp) + " @ " + usernameStyle.Render(username) + ": " + message
}
