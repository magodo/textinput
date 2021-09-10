package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	EchoMode   = textinput.EchoMode
	CursorMode = textinput.CursorMode
)

const (
	EchoNormal   = textinput.EchoNormal
	EchoPassword = textinput.EchoPassword
	EchoNone     = textinput.EchoNone

	CursorBlink  = textinput.CursorBlink
	CursorStatic = textinput.CursorStatic
	CursorHide   = textinput.CursorHide
)

type Model struct {
	textinput.Model
	CandidateWords []string
	matchingWords  []string
	index          int
}

func NewModel() Model {
	return Model{Model: textinput.NewModel()}
}

func Blink() tea.Msg {
	return textinput.Blink()
}

func Paste() tea.Msg {
	return textinput.Paste()
}