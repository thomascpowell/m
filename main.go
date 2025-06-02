package main

import (
	"os"
	"m/app"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Starts the application.
*/

func main() {
	p := tea.NewProgram(app.Model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
			os.Exit(1)
	}
}

