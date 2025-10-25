package layout

import (
	"aten039/scaff/internal/components/menu"
	msgs "aten039/scaff/internal/msg"
	"aten039/scaff/internal/scripts"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			if m.state != RunningScript {
				return m, tea.Quit
			}

		}
	case spinner.TickMsg:
		NewSpinner, cmd := m.Spinner.Update(msg)
		m.Spinner = NewSpinner
		return m, cmd

	case msgs.RunningScriptMsg:
		m.state = RunningScript
		cmd := scripts.GetScript(msg.Script)
		return m, tea.Batch(m.runScript(cmd), m.Spinner.Tick)

	case msgs.ScriptErrorMsg:
		m.state = ScriptError
		m.ScriptError = msg.Err
		return m, nil

	case msgs.ScriptFinishedMsg:
		m.state = ScriptFinished
		m.ScriptFinished = msg.Msg
		return m, nil

	case msgs.ChangeToLanguageMsg:
		m.state = LanguageModel
		return m, nil

	case msgs.ChangeToParadigmMsg:
		m.state = ParadigmModel
		return m, nil

	case msgs.ChangeToJsPooMsg:
		m.state = JsPoo
		return m, nil

	case msgs.ChangeToJsFuntionalMsg:
		m.state = JsFuntional
		return m, nil
	}

	switch m.state {

	case LanguageModel:
		newLanguage, cmd := m.Language.Update(msg)
		m.Language = newLanguage.(menu.Model)
		return m, cmd

	case ParadigmModel:
		newParadigm, cmd := m.Paradigm.Update(msg)
		m.Paradigm = newParadigm.(menu.Model)
		return m, cmd

	case JsPoo:
		newJsPoo, cmd := m.JsPoo.Update(msg)
		m.JsPoo = newJsPoo.(menu.Model)
		return m, cmd

	case JsFuntional:
		newJsFuntional, cmd := m.JsFuntional.Update(msg)
		m.JsFuntional = newJsFuntional.(menu.Model)
		return m, cmd
	}

	return m, nil
}
