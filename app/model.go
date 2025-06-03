package app

import (
	"m/utils"
	"github.com/charmbracelet/bubbles/list"
)

/**
* Contains the application Model.
*/


type View int
const (
	BaseView View = iota
	AlbumsView
	PlaylistsView
	SourceDetailView
)

type Model struct {
	Loaded					bool

	Library					utils.Library
	CurrentSong			utils.Song
	IsPlaying				bool
	CurrentView			View

	UIList 					list.Model
	CurrentList			utils.List // UIList Contents

	DetailSource		utils.SourceType // in SourceDetailView this will contain a SourceType
}


