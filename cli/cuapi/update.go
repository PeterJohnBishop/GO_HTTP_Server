package cuapi

import (
	"free-adventure-go/main.go/clickup"

	tea "github.com/charmbracelet/bubbletea"
)

func (m CUAPIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
				switch m.cursor {
				case 0:
					m.response = "Fetching Authorized User"
					return m, getAuthorizedUser
				case 1:
					m.response = "Fetching Workspaces"
					return m, getWorkspaces
				}
			}
		}

	case respMsg:
		m.response = string(msg)
	}
	return m, nil
}

type respMsg string

func getAuthorizedUser() tea.Msg {

	body, err := clickup.GetAuthorizedUser()
	if err != nil {
		return respMsg(err.Error())
	}

	return respMsg(string(body))
}

func getWorkspaces() tea.Msg {

	body, err := clickup.GetWorkspaces()
	if err != nil {
		return respMsg(err.Error())
	}

	return respMsg(string(body))
}
