package javascript

import (
	"aten039/scaff/internal/components/menu"
	"aten039/scaff/internal/msg"
)

var jsPooOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.ChangeToParadigmMsg{}},
	{Description: "Hexagonal Arquitecture", Msg: msg.ChangeToParadigmMsg{}},
	{Description: "Model View Controller", Msg: msg.ChangeToParadigmMsg{}},
	{Description: "Modular Clean", Msg: msg.ChangeToParadigmMsg{}},
}

func InitialJsPooModel() menu.Model {

	return menu.InitialModel(jsPooOptions)
}
