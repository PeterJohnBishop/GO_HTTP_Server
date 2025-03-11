package boba

import "fmt"

// Create a view string with the model data!

func (m model) View() string {

	s := "\n\nHey there!\n\n"

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	if m.response == "" {
		s += "\nWiating for an API response\n"
	} else {
		s += fmt.Sprintf("\n%s\n", m.response)
	}

	s += "\nPress q to quit.\n"

	return s

}
