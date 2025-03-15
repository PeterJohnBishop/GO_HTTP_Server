package cuapi

import "fmt"

func (m OAuthModel) View() string {

	s := "\n"

	if m.user.ID > 1 {
		s = fmt.Sprintf("\n%d %s\n\n", m.user.ID, m.user.Username)
	}

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

	s += fmt.Sprintf("\n%s\n", m.message)

	s += "\nPress q to quit.\n"

	return s

}
