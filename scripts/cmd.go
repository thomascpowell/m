package scripts

import (
	"m/utils"
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
		playing, _ := IsPlaying()
		song := GetCurrentSongObject()
		return StateMsg { 
			IsPlaying:  playing,
			CurrentSong: song,
		}
	}
}
// Contains updated state.
type StateMsg struct {	
	CurrentSong	utils.Song
	IsPlaying		bool
}


// Handles initial loading of music sources.
func GetLibraryCmd() tea.Cmd {
	return func() tea.Msg {
		library, err := GetLibraryData()
		if err != nil {
			utils.Log("GetLibraryData:" + err.Error())
		}
		return LibraryMsg(library)
	}
}
type LibraryMsg utils.Library


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
