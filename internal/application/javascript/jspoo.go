package javascript

import (
	"aten039/scaff/internal/components/menu"
	"aten039/scaff/internal/msg"
	"aten039/scaff/internal/scripts"
)

var jsPooOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsPooClean}},
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsPooHexa}},
	{Description: "Model View Controller", Msg: msg.RunningScriptMsg{Script: scripts.JsPooMvc}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.JsPooMc}},
}

func InitialJsPooModel() menu.Model {

	return menu.InitialModel(jsPooOptions)
}
