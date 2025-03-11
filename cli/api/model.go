package api

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Define the data structure.
type ApiModel struct {
	response string
	options  []string
	cursor   int
	selected map[int]struct{}
}

// Initialize the data if needed.
func InitialapiModel() ApiModel {
	return ApiModel{
		response: "",
		options:  []string{"Fetch Test Response"},
		selected: make(map[int]struct{}),
	}
}

// Send an initial CMD when the app starts.
func (a ApiModel) Init() tea.Cmd {
	return tea.Batch(tea.SetWindowTitle("Run List"), testAPI)
}
