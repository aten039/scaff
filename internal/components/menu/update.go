package menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "down", "right", "up", "left":
			return m.moveCursor(msg), nil

		case "enter":

			return m, m.SendMsg()

		default:
			return m, nil
		}

	}
	return m, nil
}
