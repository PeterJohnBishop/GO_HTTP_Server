package cuapi

import tea "github.com/charmbracelet/bubbletea"

type CUAPIModel struct {
	response string
	options  []string
	cursor   int
	selected map[int]struct{}
}

func InitCUAPIModel() CUAPIModel {
	return CUAPIModel{
		response: "",
		options:  []string{"Start OAuth"},
		selected: make(map[int]struct{}),
	}
}

func (m CUAPIModel) Init() tea.Cmd {
	return nil
}
