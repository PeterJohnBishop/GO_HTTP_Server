package boba

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	options  []string
	cursor   int
	selected map[int]struct{}
	spinner  spinner.Model
}

func InitialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{
		options:  []string{"Users", "Items", "Orders", "Invoices", "Payments"},
		selected: make(map[int]struct{}),
		spinner:  s,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick)
}
