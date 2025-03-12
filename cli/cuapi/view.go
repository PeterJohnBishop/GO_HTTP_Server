package cuapi

import "fmt"

func (m CUAPIModel) View() string {

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

	s += fmt.Sprintf("\n%s\n", m.response)

	s += "\nPress q to quit.\n"

	return s

}
