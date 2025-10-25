package menu

import tea "github.com/charmbracelet/bubbletea"

func (m Model) moveCursor(msg tea.KeyMsg) Model {

	switch msg.String() {
	case "up", "left":

		if m.cursor > 0 {
			m.cursor--
		} else if m.cursor == 0 {
			m.cursor = len(m.choices) - 1
		}
	case "down", "right":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		} else if m.cursor == len(m.choices)-1 {
			m.cursor = 0
		}
	}
	return m
}
