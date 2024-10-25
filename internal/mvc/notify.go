package mvc

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NotifyView struct {
	viewport.Model
}

type Notify struct {
	kind    NotifyKind
	message string
}

type NotifyKind int

const (
	Info NotifyKind = iota + 1
	Warn
	Err
)

var (
	InfoStyle = lipgloss.NewStyle().Background(lipgloss.Color("#3486EB")).Padding(0, 1)
	WarStyle  = lipgloss.NewStyle().Background(lipgloss.Color("#FCBA03")).Padding(0, 1)
	ErrStyle  = lipgloss.NewStyle().Background(lipgloss.Color("#EB3434")).Padding(0, 1)
)

func NewNotifyView() NotifyView {
	return NotifyView{viewport.New(Width, 1)}
}

func (nv NotifyView) Update(msg tea.Msg) (NotifyView, tea.Cmd) {
	view, cmd := nv.Model.Update(msg)
	nv.Model = view
	return nv, cmd
}

func (nd NotifyKind) String() string {
	return [...]string{"INFO", "WARN", "ERROR"}[nd-1]
}

func (nd NotifyKind) Style() lipgloss.Style {
	return [...]lipgloss.Style{InfoStyle, WarStyle, ErrStyle}[nd-1]
}

func (nv *NotifyView) Notify(notify Notify) {
	nv.SetContent(fmt.Sprintf("%s: %s", notify.kind, notify.message))
	nv.Style = notify.kind.Style()
}

func (nv *NotifyView) NotifyInfo(message string) {
	nv.Notify(Notify{Info, message})
}

func (nv *NotifyView) NotifyWarn(message string) {
	nv.Notify(Notify{Warn, message})
}

func (nv *NotifyView) NotifyErr(err error) {
	nv.Notify(Notify{Err, err.Error()})
}
