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

func (m OAuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
				if len(m.teams) == 0 {
					switch m.cursor {

					case 0:
						m.cursor = 0
						m.selected[m.cursor] = struct{}{}
						return m, oAuthStart
					case 1:
						m.cursor = 0
						m.selected[m.cursor] = struct{}{}
						return m, tea.Batch(getAccessToken, getAuthorizedUser)
					}
				} else {
					for index := range m.teams {
						if m.cursor == index {
							if ok {
								delete(m.selected, m.cursor)
							} else {
								delete(m.selected, m.cursor)
								// m.selected[m.cursor] = struct{}{}
								return InitHierarchyModel(m.user, m.teams[m.cursor]), nil
							}
						}
					}
				}
			}
		}

	case oauthMsg:
		m.message = "Ready!"
		m.options = append(m.options, string(msg))
		return m, nil
	case tokenMsg:
		m.message = "Your OAuth token is saved! Select a Workspace."
		return m, getWorkspaces
	case wkspcMsg:
		m.message = "Workspaces Found!"
		m.selected = make(map[int]struct{})
		m.options = []string{}
		for _, team := range msg {
			m.options = append(m.options, team.ID+" "+team.Name)
			m.teams = append(m.teams, team)
		}
	case usrMsg:
		m.user.ID = msg.ID
		m.user.Username = msg.Username
		m.user.Email = msg.Email
		m.user.Color = msg.Color
		m.user.Initials = msg.Initials
		m.user.Timezone = msg.Timezone
		m.user.WeekStartDay = msg.WeekStartDay
		m.user.ProfilePicture = msg.ProfilePicture
	}
	return m, nil
}

type errMsg string
type oauthMsg string
type tokenMsg string
type wkspcMsg []Team
type usrMsg User

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

func getWorkspaces() tea.Msg {

	body, err := clickup.GetWorkspaces()
	if err != nil {
		return errMsg(err.Error())
	}

	var res TeamResponse

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return errMsg(err.Error())
	}

	return wkspcMsg(res.Teams)
}

func getAuthorizedUser() tea.Msg {

	body, err := clickup.GetAuthorizedUser()
	if err != nil {
		return errMsg(err.Error())
	}

	var res UserResponse

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return errMsg(err.Error())
	}
	fmt.Println(res.User)
	return usrMsg(res.User)
}
