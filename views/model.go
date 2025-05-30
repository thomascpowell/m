package views

import (
	"m/scripts"
	"m/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Contains the application Model and update loop.
*/

type View int
const (
	BaseView View = iota
	AlbumsView
	PlaylistsView
	SourceDetailView
)

type Model struct {
	Library					utils.Library
	CurrentSong			utils.Song
	IsPlaying				bool
	CurrentView			View

	UIList 					list.Model
	CurrentList			utils.List // UIList Contents
	DetailSource		utils.SourceType // in SourceDetailView this will contain a SourceType
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		scripts.GetLibraryCmd(), 
		scripts.RefreshStateCmd(),
		scripts.RefreshLibraryCmd(),
		TickCmd(),
	)
}

