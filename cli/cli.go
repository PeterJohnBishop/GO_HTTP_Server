package cli

import (
	"fmt"
	"os"

	swticher "free-adventure-go/main.go/cli/switcher"

	tea "github.com/charmbracelet/bubbletea"
)

func StartCLI() {
	p := tea.NewProgram(swticher.InitialSwitchModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
