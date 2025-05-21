package utils

import (
	"fmt"
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
			Log("RunAsCmd - " + name + ": " +  err.Error())
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
	CurrentSong	Song
	IsPlaying		bool
}


// Handles initial loading of music sources.
func FetchSourcesCmd() tea.Cmd {
	return func() tea.Msg {
		albums, playlists := FetchSources()
		return SourcesMsg {
			Albums: albums,
			Playlists: playlists,
		}
	}
}
// Contains albums and playlists
type SourcesMsg struct {
	Albums []Source
	Playlists []Source
}


// Handles data fetching for playlist or album views.
// Sends List struct to update the CurrentList.
// Updates CurrentList to specified album or playlist.
func UpdateListCmd(source SourceType, name string, desc string) tea.Cmd {
	Log("Call to UpdateListCmd.")
	return func() tea.Msg {
		var songs []Song
		var owner string
		var err error
		switch source {
		case Album:
			songs, err = GetSongsFromSource("album", name)
			if err != nil {
				Log(fmt.Sprintf("UpdateListCmd: failed to get songs from %s: %v", name, err))
			}
			owner = desc
		case Playlist:
			songs, err = GetSongsFromSource("playlist", name)
			owner = "you"
		default:
			err = nil
			songs = nil
			owner = ""
		}
		if err != nil {
			Log("UpdateListCmd: " + err.Error())
		}
		return ListMsg {
			Name: name,
			Owner: owner,
			Songs: songs,
		}
	}
}
// Wrapper type to indicate this List is sent as a Msg
// When this arrives, view is updated to show data in UIList.
type ListMsg List
