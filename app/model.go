package app

import (
	"m/utils"
	"github.com/charmbracelet/bubbles/list"
)

/**
* Contains the application Model.
*/


type Model struct {
	Loaded					bool
	Library					utils.Library
	CurrentSong			utils.Song
	IsPlaying				bool
	CurrentView			utils.View
	UIList 					list.Model
}


