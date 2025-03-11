package boba

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	greeting string
}

func InitialModel() model {
	return model{
		greeting: "Hello, World!",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
