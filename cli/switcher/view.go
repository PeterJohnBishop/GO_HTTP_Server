package swticher

import "fmt"

func (sw SwitchModel) View() string {
	s := "\n\nHey there!\n"

	if sw.message == "" {
		s += "\nReady to Test!\n\n"
	} else {
		s += fmt.Sprintf("\n%s\n", sw.message)
	}

	for i, choice := range sw.options {
		cursor := " "
		if sw.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := sw.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s

}
