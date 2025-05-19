package main

import (
	"fmt"
	"os"
	"time"
	tea "github.com/charmbracelet/bubbletea"
	// bubbles "github.com/charmbracelet/bubbles"
)

// Main
func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
	}
}

type Song struct {
	Title			string
	Artist 		string
	Duration	float32
}
type model struct {
	Albums				[]string
	Lists					[]string
	CurrentSource	string
	CurrentSong		Song
	IsPlaying			bool
	// CurrentView		View
}


// Init
func (m model) Init() tea.Cmd {
	m.refreshState()
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*3, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
type tickMsg struct{}


// Update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Tick
	case tickMsg:
		m.refreshState()
		return m, tickCmd()

	// Keypress
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "space":
			// TODO: Toggle play
		}
	}
	// Default
	return m, nil
}
func (m *model) refreshState() tea.Cmd {
	if albums, err := GetAlbums(); err == nil {
			m.Albums = albums
	}
	if lists, err := GetPlaylists(); err == nil {
			m.Lists = lists
	}
	if playing, err := IsPlaying(); err == nil {
			m.IsPlaying = playing
	}
	m.CurrentSong = GetCurrentSongObject()
	return nil
}

// View
func (m model) View() string {
    status := "Stopped"
    if m.IsPlaying {
        status = "Playing"
    }
    return fmt.Sprintf("Status: %s\nCurrent song: %s - %s\n", status, m.CurrentSong.Artist, m.CurrentSong.Title)
}
