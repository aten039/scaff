package menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

type OptionsInfo struct {
	Description string
	Msg         tea.Msg
}

type Model struct {
	choices []OptionsInfo
	cursor  int
}

func InitialModel(options []OptionsInfo) Model {

	return Model{
		choices: options,
		cursor:  0,
	}
}
