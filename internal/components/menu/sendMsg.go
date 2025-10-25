package menu

import tea "github.com/charmbracelet/bubbletea"

func (m Model) SendMsg() tea.Cmd {

	return func() tea.Msg {
		return m.choices[m.cursor].Msg
	}
}
