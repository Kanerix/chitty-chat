package mvc

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
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

func NewNotifyView() NotifyView {
	view := NotifyView{
		Model: viewport.New(Width, 1),
	}

	view.SetContent("test")

	return view
}

func (d NotifyKind) String() string {
	return [...]string{"Info", "Warn", "Err"}[d-1]
}

func (nv NotifyView) Notify(notify Notify) {
	nv.SetContent(fmt.Sprintf("%d: %s", notify.kind, notify.message))
}

func (nv NotifyView) NotifyInfo(message string) {
	nv.Notify(Notify{Info, message})
}

func (nv NotifyView) NotifyWarn(message string) {
	nv.Notify(Notify{Warn, message})
}

func (nv NotifyView) NotifyErr(err error) {
	nv.Notify(Notify{Err, err.Error()})
}
