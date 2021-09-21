package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type CandidateViewMode int

const (
	CandidateViewVertical CandidateViewMode = iota
	CandidateViewHorizental
	CandidateViewNone
)

const (
	defaultWidth = 40
)

type matchingRange struct {
	sindex int
	eindex int
}

type Model struct {
	textinput.Model
	CandidateWords []string
	matchingWords  []string
	index          int

	CandidateViewMode CandidateViewMode
	Width             int
	ranges            []matchingRange
	rangeIndex        int

	StyleMatching  lipgloss.Style
	StyleCandidate lipgloss.Style
}

func NewModel() Model {
	return Model{
		Model: textinput.NewModel(),
		Width: defaultWidth,
	}
}

func Blink() tea.Msg {
	return textinput.Blink()
}

func Paste() tea.Msg {
	return textinput.Paste()
}
