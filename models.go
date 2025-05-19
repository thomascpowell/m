package main

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

type Song struct {
	Title			string
	Artist 		string
	Duration	float32
}

type model struct {
	Albums				[]string
	Playlists			[]string
	CurrentSong		Song
	IsPlaying			bool
	Debug					string
	// CurrentView		View
}

type stateMsg struct {	
	Albums      []string
	Playlists   []string
	CurrentSong		Song
	IsPlaying			bool
}


func (m model) Init() tea.Cmd {
	return tea.Batch(tickCmd(), refreshStateCmd())
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*3, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
type tickMsg struct{}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case stateMsg:
		m.Albums = msg.Albums
		m.Playlists = msg.Playlists
		m.CurrentSong = msg.CurrentSong
		m.IsPlaying = msg.IsPlaying
		return m, nil

	case tickMsg:
		return m, tea.Batch(tickCmd(), refreshStateCmd())

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			return m, RunAsCmd("toggle", TogglePlayPause)
		}

	case CmdResultMsg:
		m.Debug = msg.Name + "ran"
    if msg.Err != nil {
			m.Debug += "(error)"
    }
    return m, nil
	}
	// default
	return m, nil
}

func refreshStateCmd() tea.Cmd {
	return func() tea.Msg {
		playing, _ := IsPlaying()
		song := GetCurrentSongObject()
		albums, _ := GetAlbums()
		playlists, _ := GetPlaylists()
		return stateMsg{
			IsPlaying:  playing,
			CurrentSong: song,
			Albums:     albums,
			Playlists:  playlists,
		}
	}
}
