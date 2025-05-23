package main

import (
	"os"
	"m/views"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Entry.
*/

func main() {
	p := tea.NewProgram(views.Model{})
	if _, err := p.Run(); err != nil {
			os.Exit(1)
	}
}

