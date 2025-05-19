package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

// Main
func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
			os.Exit(1)
	}
}

// View
func (m model) View() string {
	return fmt.Sprintf("Debug: %s\nCurrent song: %s - %s\n",m.Debug, m.CurrentSong.Artist, m.CurrentSong.Title)
}
