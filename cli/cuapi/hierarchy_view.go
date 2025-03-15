package cuapi

import (
	"fmt"
)

func (m HierarchyModel) View() string {

	var levels = []string{"Spaces", "Folders", "Folderless Lists", "Lists", "Tasks"}

	s := fmt.Sprintf("\n%d %s\n\n", m.user.ID, m.user.Username)

	switch m.lvl {
	case 0: // Spaces
		s += fmt.Sprintf("\n%s in %s/\n", levels[m.lvl], m.team.Name)
	case 1: // Folders
		s += fmt.Sprintf("\n%s in %s/%s/\n", levels[m.lvl], m.team.Name, m.space.Name)
	case 2: // Folderless Lists
		s += fmt.Sprintf("\n%s in %s/%s/\n", levels[m.lvl], m.team.Name, m.space.Name)
	case 3: // Lists
		s += fmt.Sprintf("\n%s in %s/%s/%s/\n", levels[m.lvl], m.team.Name, m.space.Name, m.folder.Name)
	case 4: // Tasks
		s += fmt.Sprintf("\n%s in %s/%s/%s/%s\n", levels[m.lvl], m.team.Name, m.space.Name, m.folder.Name, m.list.Name)
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
