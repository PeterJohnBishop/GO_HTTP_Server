package swticher

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SwitchModel struct {
	message  string
	options  []string
	cursor   int
	selected map[int]struct{}
}

func InitialSwitchModel() SwitchModel {
	return SwitchModel{
		options:  []string{"test", "login", "exit"},
		selected: make(map[int]struct{}),
	}
}

func (sw SwitchModel) Init() tea.Cmd {
	return nil
}
