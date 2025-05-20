package main

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

type View int
const (
	BaseView View = iota
	AlbumsView
	PlaylistsView
	SourceDetailView
)

// Represents a Song.
type Song struct {
	Title			string
	Artist 		string
	Duration	string
}

// Represents a list of Songs
// Can be an album or a playlist
// Used when viewing one of the above.
type List struct {
	Name 		string
	Owner		string
	Songs 	[]Song
}

// Represents the possible sources
// (Playlist and Album)
type Source int
const (
	Album = iota
	Playlist
)

// Represents the application state.
type model struct {
	Albums				[]string
	Playlists			[]string
	CurrentSong		Song
	IsPlaying			bool
	CurrentView		View
	CurrentList		List
}

// starts the event loop
func (m model) Init() tea.Cmd {
	return tea.Batch(tickCmd(), RefreshStateCmd())
}

// sends a tickMsg after 3 seconds
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*3, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
type tickMsg struct{}

// update loop
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// state update
	case stateMsg:
		m.Albums = msg.Albums
		m.Playlists = msg.Playlists
		m.CurrentSong = msg.CurrentSong
		m.IsPlaying = msg.IsPlaying
		return m, nil

	// tick occurs
	case tickMsg:
		return m, tea.Batch(tickCmd(), RefreshStateCmd())

	// new list 
	case ListMsg:
		m.CurrentList = List(msg)
		return m, nil

	// keypress
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			m.IsPlaying = false 
			return m, RunAsCmd("toggle", TogglePlayPause)
		case "a":
			m.CurrentView = AlbumsView
			return m, nil
		case "p":
			m.CurrentView = PlaylistsView
			return m, nil
		case "b":
			m.CurrentView = BaseView
			return m, nil
		case "0": 
			// for testing sourceDetailView
			// test playlist name
			cmd := UpdateListCmd(Playlist, "study")
			m.CurrentView = SourceDetailView
			return m, cmd
		default:
			return m, nil
		}
	}
	// default
	return m, nil
}

