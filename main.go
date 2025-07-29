package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"m/app"
	"m/scripts"
	"os"
	"strings"
)

/**
* Starts the application.
 */

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		handleArgs(args)
		os.Exit(0)
	}
	p := tea.NewProgram(app.Model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

func handleArgs(args []string) {
	usage := "Usage: m [play|skip|prev]"
	cmd := strings.ToLower(args[0])
	switch cmd {
	case "play", "pause", "p":
		_ = scripts.RunAsCli("Toggling playback", "✓", scripts.TogglePlayPause)
		return
	case "skip":
		_ = scripts.RunAsCli("Skipping", "✓", scripts.NextTrack)
		return
	case "prev":
		_ = scripts.RunAsCli("Rewinding", "✓", scripts.PreviousTrack)
		return
	}
	println(usage)
	return
}
