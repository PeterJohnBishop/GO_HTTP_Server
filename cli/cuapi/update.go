package cuapi

import (
	"encoding/json"
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
			if !ok {
				m.selected[m.cursor] = struct{}{}
				switch m.cursor {
				case 0:
					m.cursor = 0
					m.selected[m.cursor] = struct{}{}
					return m, oAuthStart
				case 1:
					m.cursor = 0
					m.selected[m.cursor] = struct{}{}
					return m, getAccessToken

				}
			}
		}

	case oauthMsg:
		m.response = "Ready!"
		m.options = append(m.options, string(msg))
		return m, nil
	case tokenMsg:
		m.response = "Your OAuth token is saved! Select a Workspace."
		return m, getWorkspaces
	case wkspcMsg:
		m.response = "Workspaces Found!"
		m.selected = make(map[int]struct{})
		m.options = []string{}
		for _, team := range msg {
			m.options = append(m.options, team.ID+" "+team.Name)
		}
	}
	return m, nil
}

type errMsg string
type oauthMsg string
type tokenMsg string
type wkspcMsg []Team
type usrMsg string

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Teams []Team `json:"teams"`
}

func oAuthStart() tea.Msg {

	err := godotenv.Load()
	if err != nil {
		return errMsg(err.Error())
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

	err = cmd.Start()
	if err != nil {
		return errMsg(err.Error())
	}

	return oauthMsg("Find my Authorized Workspaces!")
}

func getAccessToken() tea.Msg {
	err := godotenv.Load()
	if err != nil {
		return errMsg(err.Error())
	}
	client_id := os.Getenv("CLICKUP_CLIENT_ID")
	client_secret := os.Getenv("CLICKUP_CLIENT_SECRET")

	_, err = clickup.GetAccessToken(client_id, client_secret)
	if err != nil {
		return errMsg(err.Error())
	}

	return tokenMsg("Get Workspaces")
}

func getAuthorizedUser() tea.Msg {

	body, err := clickup.GetAuthorizedUser()
	if err != nil {
		return errMsg(err.Error())
	}

	return usrMsg(string(body))
}

func getWorkspaces() tea.Msg {

	body, err := clickup.GetWorkspaces()
	if err != nil {
		return errMsg(err.Error())
	}

	var res Response

	err = json.Unmarshal([]byte(body), &res)

	return wkspcMsg(res.Teams)
}
