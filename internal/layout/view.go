package layout

import "fmt"

func (m AppModel) View() string {

	switch m.state {
	case RunningScript:
		return fmt.Sprintf("\n%s configuring project...", m.Spinner.View())

	case ScriptError:
		return fmt.Sprintf("\n%s  \nPress q to quit.\n", m.ScriptError)

	case ScriptFinished:
		return fmt.Sprintf("\nfinished successfully\n%s  \nPress q to quit.\n", m.ScriptFinished)

	case LanguageModel:
		return m.Language.View()

	case ParadigmModel:
		return m.Paradigm.View()

	case JsPoo:
		return m.JsPoo.View()

	case JsFuntional:
		return m.JsFuntional.View()

	default:
		return "unknown error"
	}
}
