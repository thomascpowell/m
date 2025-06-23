package main

import (
	"os"
	"strings"
	"m/scripts"
	"m/app"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Starts the application.
*/

func main() {

	args := os.Args[1:]
	
	if len(args) != 0 {
		print(handleArgs(args))
		os.Exit(0)
	}

	p := tea.NewProgram(app.Model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
			os.Exit(1)
	}
}


func handleArgs(args []string) (string) {
	usage := `
	Usage: 
		m [play|pause|skip]
		m [album|playlist] "<name>"`

	cmd := strings.ToLower(args[0])
	args_count := len(args)

	// single arg commands
	switch cmd {
	case "play", "pause":
		scripts.TogglePlayPause()
		return "Attempting to " + cmd
	case "skip":
		scripts.NextTrack()
		return "Attempting to skip track..."
		// can we get action handler type thing? to show result?
	case "prev":
		scripts.PreviousTrack()
	}
	
	// all other valid commands are 2 arg
		
	if args_count != 2 {
		return usage
	}

	value := strings.ToLower(args[1])


	if 




	return ""
}
