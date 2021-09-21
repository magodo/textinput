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
				m.matchingWords = []string{v}
				for _, candidate := range m.CandidateWords {
					if strings.HasPrefix(candidate, v) {
						m.matchingWords = append(m.matchingWords, candidate)
					}
				}

				m.ranges, m.rangeIndex = m.buildRanges()

				// The first time entering matching mode will stop here, the user needs to press the tab the second time to get the first completion.
				break
			}
			if len(m.matchingWords) == 1 {
				break
			}

			if v == m.matchingWords[m.index] {
				if msg.Type == tea.KeyTab {
					m.index = (m.index + 1) % len(m.matchingWords)
					if m.index == m.ranges[m.rangeIndex].eindex || m.index == 0 {
						m.rangeIndex = (m.rangeIndex + 1) % len(m.ranges)
					}
				} else {
					m.index = (m.index - 1 + len(m.matchingWords)) % len(m.matchingWords)
					if m.index < m.ranges[m.rangeIndex].sindex || m.index == len(m.matchingWords)-1 {
						m.rangeIndex = (m.rangeIndex - 1 + len(m.ranges)) % len(m.ranges)
					}
				}
			}

			m.SetValue(m.matchingWords[m.index])
			m.SetCursor(len(m.matchingWords[m.index]))
		default:
			m.matchingWords = nil
			m.index = 0
			m.rangeIndex = 0
			m.ranges = nil
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.ranges, m.rangeIndex = m.buildRanges()
	}
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	out := m.Model.View()
	if len(m.matchingWords) <= 1 || m.CandidateViewMode == CandidateViewNone {
		return out
	}
	if m.CandidateViewMode == CandidateViewVertical {
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

	sindex, eindex := m.ranges[m.rangeIndex].sindex, m.ranges[m.rangeIndex].eindex
	mlist := make([]string, 0, eindex-sindex)
	for idx, w := range m.matchingWords[sindex:eindex] {
		if m.index == sindex+idx {
			w = m.StyleMatching.Render(w)
		} else {
			w = m.StyleCandidate.Render(w)
		}
		mlist = append(mlist, w)
	}
	return out + "\n" + strings.Join(mlist, " ")
}

// buildRanges build the ranges and the range index based on current index and the window size.
func (m Model) buildRanges() ([]matchingRange, int) {
	var ranges []matchingRange
	var width int
	var sindex int
	for i := 0; i < len(m.matchingWords); i++ {
		if width == 0 {
			sindex = i
			width = len(m.matchingWords[i])
			continue
		}

		width += len(m.matchingWords[i]) + 1
		if width > m.Width {
			ranges = append(ranges, matchingRange{sindex: sindex, eindex: i})
			width = 0
			i--
			continue
		}
	}
	if width != 0 {
		ranges = append(ranges, matchingRange{sindex: sindex, eindex: len(m.matchingWords)})
	}

	rangeIndex := 0
	for i := 0; i < m.index; i++ {
		if i == ranges[rangeIndex].eindex {
			rangeIndex++
		}
	}

	return ranges, rangeIndex
}
