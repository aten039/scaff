package javascript

import (
	"aten039/scaff/internal/components/menu"
	"aten039/scaff/internal/msg"
	"aten039/scaff/internal/scripts"
)

var jsfuntionalOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalClean}},
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalHexa}},
	{Description: "Model View Controller", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalMvc}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalMC}},
}

func InitialjsfuntionalModel() menu.Model {

	return menu.InitialModel(jsfuntionalOptions)
}
