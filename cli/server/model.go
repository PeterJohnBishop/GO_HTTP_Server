package server

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ServerModel struct {
	response string
	options  []string
	cursor   int
	selected map[int]struct{}
}

func InitServerModel() ServerModel {
	return ServerModel{
		response: "",
		options:  []string{"Fetch Test Response"},
		selected: make(map[int]struct{}),
	}
}

func (m ServerModel) Init() tea.Cmd {
	return nil
}
