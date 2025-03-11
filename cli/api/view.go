package api

import "fmt"

// Create a view string with the model data!

func (a ApiModel) View() string {

	s := "\n\nHey there!\n\n"

	for i, choice := range a.options {
		cursor := " "
		if a.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := a.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	if a.response == "" {
		s += "\nVerifying server connection...\n"
	} else {
		s += fmt.Sprintf("\n%s\n", a.response)
	}

	s += "\nPress q to quit.\n"

	return s

}
