package cuapi

import (
	"fmt"
	"free-adventure-go/main.go/clickup"
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
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
					if m.options[0] == "Get Access Token" {
						m.response = "Fetching Access Token"
						return m, getAccessToken
					} else {
						m.response = "Starting OAuth Flow"
						return m, oAuthStart
					}
				case 1:
					m.response = "Fetching Workspaces"
					return m, getWorkspaces
				case 2:
					m.response = "Fetching Authorized User"
					return m, getAuthorizedUser
				}
			}
		}

	case respMsg:
		switch msg {
		case "Waiting for code.":
			m.options[0] = "Get Access Token"
			delete(m.selected, 0)
		}
		m.response = string(msg)
	}
	return m, nil
}

type respMsg string

func oAuthStart() tea.Msg {

	err := godotenv.Load()
	if err != nil {
		return respMsg(err.Error())
	}
	client_id := os.Getenv("CLICKUP_CLIENT_ID")
	redirect_uri := os.Getenv("CLICKUP_REDIRECT_URI")

	var url string = fmt.Sprintf("https://app.clickup.com/api?client_id=%s&redirect_uri=%s", client_id, redirect_uri)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	case "windows": // Windows
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default: // Linux and others
		cmd = exec.Command("xdg-open", url)
	}

	browseErr := cmd.Start()
	if browseErr != nil {
		return respMsg(browseErr.Error())
	}
	return respMsg("Waiting for code.")
}

func getAccessToken() tea.Msg {
	err := godotenv.Load()
	if err != nil {
		return respMsg(err.Error())
	}
	client_id := os.Getenv("CLICKUP_CLIENT_ID")
	client_secret := os.Getenv("CLICKUP_CLIENT_SECRET")

	body, err := clickup.GetAccessToken(client_id, client_secret)
	if err != nil {
		return respMsg(err.Error())
	}

	return respMsg(string(body))

}

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
