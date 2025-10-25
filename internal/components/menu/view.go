package menu

import "fmt"

func (m Model) View() string {
	// The header
	s := "\nselect your language?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// // Is this choice selected?
		// checked := " " // not selected
		// if _, ok := m.selected[i]; ok {
		// 	checked = "x" // selected!
		// }
		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor /*checked,*/, choice.Description)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
