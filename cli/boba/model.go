package boba

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Define the data structure.
type model struct {
	response string
	options  []string
	cursor   int
	selected map[int]struct{}
}

// Initialize the data if needed.
func InitialModel() model {
	return model{
		response: "",
		options:  []string{"Fetch Test Response"},
		selected: make(map[int]struct{}),
	}
}

// Send an initial CMD when the app starts.
func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Run List")
}
