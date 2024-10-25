package mvc

import "github.com/charmbracelet/bubbles/viewport"

type ChatView struct {
	viewport.Model
	messages []string
}

func NewChatView() ChatView {
	view := ChatView{
		Model:    viewport.New(Width, 5),
		messages: make([]string, 10),
	}

	view.SetContent("Welcome to Chitty Chat!")

	return view
}
