package api

import (
	"io"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

// When an event (Msg) occurs, do this.
func (a ApiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return a, tea.Quit
		case "up", "k":
			if a.cursor > 0 {
				a.cursor--
			}
		case "down", "j":
			if a.cursor < len(a.options)-1 {
				a.cursor++
			}
		case "enter", " ":
			_, ok := a.selected[a.cursor]
			if ok {
				delete(a.selected, a.cursor)
			} else {
				a.selected[a.cursor] = struct{}{}
			}
		}

	case initMsg:
		a.response = string(msg)
	}
	return a, nil
}

type initMsg string

func testAPI() tea.Msg {
	url := "http://localhost:8080/" // Example API

	resp, err := http.Get(url)
	if err != nil {
		return initMsg(err.Error())
	}
	defer resp.Body.Close() // Ensure response body is closed

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return initMsg(err.Error())
	}

	return initMsg(string(body))
}
