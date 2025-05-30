package scripts

import (
	"m/utils"
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* tea.Cmd functions.
*/


// Wraps sync functions.
// Used for control commands, not data fetching.
func RunAsCmd(name string, fn func() error) tea.Cmd {
	return func() tea.Msg {
		err := fn()
		if err != nil {
			utils.Log("RunAsCmd - " + name + ": " +  err.Error())
		}
		return CmdResultMsg {}
	}
}
type CmdResultMsg struct {}


// Handles state refesh ticks.
// Sends StateMsg with updated results.
func RefreshStateCmd() tea.Cmd {
	return func() tea.Msg {
		current_song, player_state := GetPlayerState()
		return StateMsg { 
			IsPlaying:  player_state,
			CurrentSong: current_song,
		}
	}
}
// Contains updated state.
type StateMsg struct {	
	CurrentSong	utils.Song
	IsPlaying		bool
}


// handles initial loading of music sources
// loads from cache initially
func GetLibraryCmd() tea.Cmd {
	CACHE_PATH := utils.GetGlobalCachePath()
	return func() tea.Msg {
		var(
			library *utils.Library
			err error
		)
		if utils.FileExists(CACHE_PATH) {
			library, err = LoadLibrary(CACHE_PATH)
		} else {
			library, err = GetLibraryData()
				_ = SaveLibrary(library, CACHE_PATH)
		}
		if err != nil {
			utils.Log("GetLibraryData:" + err.Error())
			return LibraryMsg(utils.Library{})
		}
		return LibraryMsg(*library)
	}
}
type LibraryMsg utils.Library

// handles refreshing the library
// (initial loading is from cache)
func RefreshLibraryCmd() tea.Cmd {
	CACHE_PATH := utils.GetGlobalCachePath()
	return func() tea.Msg {
		time.Sleep(10 * time.Second)
		library, err := GetLibraryData()
		if err != nil {
				utils.Log("RefreshLibraryCmd:" + err.Error())
				return nil
		}
		_ = SaveLibrary(library, CACHE_PATH)
		return LibraryMsg(*library)
	}
}

// Handles data fetching for playlist or album views.
// Sends List struct to update the CurrentList.
// Updates CurrentList to specified album or playlist.
func UpdateListCmd(kind utils.SourceType, source utils.Source, library utils.Library) tea.Cmd {
	return func() tea.Msg {
		songs := GetSongsFromSource(kind, source, library)
		return ListMsg {
			Name: source.Title,
			Owner: source.Artist,
			Songs: songs,
		}
	}
}
type ListMsg utils.List
