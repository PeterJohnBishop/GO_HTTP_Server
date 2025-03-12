package server

import (
	"io"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

func (m ServerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
				m.response = "Fetching Server"
				return m, testAPI
			}
		}

	case testMsg:
		m.response = string(msg)
	}
	return m, nil
}

type testMsg string

func testAPI() tea.Msg {
	url := "http://localhost:8080" // Example API

	resp, err := http.Get(url)
	if err != nil {
		return testMsg(err.Error())
	}
	defer resp.Body.Close() // Ensure response body is closed

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return testMsg(err.Error())
	}

	return testMsg(string(body))
}
