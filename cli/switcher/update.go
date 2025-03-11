package swticher

import (
	"io"
	"net/http"

	"free-adventure-go/main.go/cli/api"

	tea "github.com/charmbracelet/bubbletea"
)

func (sw SwitchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return sw, tea.Quit
		case "up", "k":
			if sw.cursor > 0 {
				sw.cursor--
			}
		case "down", "j":
			if sw.cursor < len(sw.options)-1 {
				sw.cursor++
			}
		case "enter", " ":
			_, ok := sw.selected[sw.cursor]
			if ok {
				delete(sw.selected, sw.cursor)
				if sw.cursor == 0 {
					sw.message = "Ready to test, again!\n"
				}
			} else {
				sw.selected[sw.cursor] = struct{}{}
				switch sw.cursor {
				case 0:
					return sw, test
				case 1:
					return api.InitialapiModel(), nil
				case 2:
					return sw, tea.Quit
				}
			}
		}

	case testMsg:
		sw.message = string(msg)
	}
	return sw, nil
}

type testMsg string

func test() tea.Msg {
	url := "http://localhost:8080/" // Example API

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
