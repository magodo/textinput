package textinput

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab:
			v := m.Value()
			if len(v) == 0 {
				break
			}
			if m.matchingWords == nil {
				m.matchingWords = []string{}
				for _, candidate := range m.CandidateWords {
					if strings.HasPrefix(candidate, v) {
						m.matchingWords = append(m.matchingWords, candidate)
					}
				}
			}
			if len(m.matchingWords) == 0 {
				break
			}

			// For the first time entering the auto-completion loop, this condition check ensures we start from the first item
			// (as long as that item is not the same as what is currently entered)
			if v == m.matchingWords[m.index] {
				if msg.Type == tea.KeyTab {
					m.index = (m.index + 1) % len(m.matchingWords)
				} else {
					m.index = (m.index - 1 + len(m.matchingWords)) % len(m.matchingWords)
				}
			}

			m.SetValue(m.matchingWords[m.index])
			m.SetCursor(len(m.matchingWords[m.index]))
		default:
			m.matchingWords = nil
			m.index = 0
		}
	}
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	out := m.Model.View()
	if len(m.matchingWords) == 0 || !m.ShowCandidate {
		return out
	}
	mlist := []string{}
	for idx, w := range m.matchingWords {
		if idx == m.index {
			w = m.StyleMatching.Render(w)
		} else {
			w = m.StyleCandidate.Render(w)
		}
		mlist = append(mlist, w)
	}
	return out + "\n" + strings.Join(mlist, "\n")
}
