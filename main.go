package main

import (
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Entry.
*/

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
			os.Exit(1)
	}
}

