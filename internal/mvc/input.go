package mvc

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InputArea struct {
	textarea.Model
}

func NewInputArea() InputArea {
	area := InputArea{textarea.New()}
	area.Placeholder = "Send a message..."
	area.Focus()

	area.Prompt = ">>> "
	area.CharLimit = 128

	area.SetWidth(Width)
	area.SetHeight(1)

	area.FocusedStyle.CursorLine = lipgloss.NewStyle()
	area.ShowLineNumbers = false
	area.KeyMap.InsertNewline.SetEnabled(false)

	return area
}

func (ia InputArea) Init() tea.Cmd {
	return textarea.Blink
}

func (ia InputArea) Update(msg tea.Msg) (InputArea, tea.Cmd) {
	area, cmd := ia.Model.Update(msg)
	ia.Model = area
	return ia, cmd
}
