package scripts

import "os/exec"

func GetScript(selection ScriptApp) *exec.Cmd {
	return exec.Command("cmd", "/C", "echo", "¡Hola, Bubbletea en Windows!")

}
