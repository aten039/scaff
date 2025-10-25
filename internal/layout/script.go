package layout

import (
	"aten039/scaff/internal/msg"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *AppModel) runScript(script *exec.Cmd) tea.Cmd {

	return func() tea.Msg {

		time.Sleep(3 * time.Second)
		output, err := script.CombinedOutput()

		if err != nil {
			return msg.ScriptErrorMsg{Err: err.Error() + ": " + string(output)}
		}

		return msg.ScriptFinishedMsg{Msg: string(output)}
	}
}
