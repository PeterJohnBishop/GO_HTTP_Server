package cli

import (
	"fmt"
	"os"

	"free-adventure-go/main.go/cli/cuapi"

	tea "github.com/charmbracelet/bubbletea"
)

func StartCLI() {
	p := tea.NewProgram(cuapi.InitOAuthModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
