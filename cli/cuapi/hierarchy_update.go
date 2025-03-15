package cuapi

import (
	"encoding/json"
	"free-adventure-go/main.go/clickup"

	tea "github.com/charmbracelet/bubbletea"
)

func (m HierarchyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if !ok {
				m.selected[m.cursor] = struct{}{}
			} else {
				delete(m.selected, m.cursor)
			}
		}

	case spcMsg:
		m.spaces = spcMsg(msg)
		for _, s := range m.spaces {
			m.options = append(m.options, s.ID+" "+s.Name)
		}
	}

	return m, nil
}

type spcMsg []Space
type fldMsg []Folder
type fldlsMsg []List
type lstMsg []List
type tskMsg []Task

func getSpaces(id string) tea.Cmd {

	return func() tea.Msg {
		body, err := clickup.GetSpaces(id)
		if err != nil {
			return errMsg(err.Error())
		}

		var res SpaceResponse

		err = json.Unmarshal([]byte(body), &res)
		if err != nil {
			return errMsg(err.Error())
		}

		return spcMsg(res.Spaces)
	}
}

func getFolders(id string) tea.Cmd {
	return func() tea.Msg {
		body, err := clickup.GetFolders(id)
		if err != nil {
			return errMsg(err.Error())
		}

		var res FolderResponse

		err = json.Unmarshal([]byte(body), &res)
		if err != nil {
			return errMsg(err.Error())
		}

		return fldMsg(res.Folders)
	}
}

func getFolderlessLists(id string) tea.Cmd {
	return func() tea.Msg {
		body, err := clickup.GetFolderlessLists(id)
		if err != nil {
			return errMsg(err.Error())
		}

		var res ListResponse

		err = json.Unmarshal([]byte(body), &res)
		if err != nil {
			return errMsg(err.Error())
		}

		return fldlsMsg(res.Lists)
	}
}

func getLists(id string) tea.Cmd {
	return func() tea.Msg {
		body, err := clickup.GetLists(id)
		if err != nil {
			return errMsg(err.Error())
		}

		var res ListResponse

		err = json.Unmarshal([]byte(body), &res)
		if err != nil {
			return errMsg(err.Error())
		}

		return lstMsg(res.Lists)
	}
}

func getTasks(id string) tea.Cmd {
	return func() tea.Msg {
		body, err := clickup.GetTasks(id)
		if err != nil {
			return errMsg(err.Error())
		}

		var res TaskResponse

		err = json.Unmarshal([]byte(body), &res)
		if err != nil {
			return errMsg(err.Error())
		}

		return tskMsg(res.Tasks)
	}
}
