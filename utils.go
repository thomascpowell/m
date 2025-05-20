package main

import (
	"bytes"
	"os/exec"
	"strings"
	"strconv"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)


/**
* Async wrappers for sync. utility functions.
*	Returns Cmd funtions that eventually return a Msg.
*/

// Wraps sync functions into Cmd.
// Sends CmdResultMsg for error reporting.
// Used for control commands, not data fetching.
func RunAsCmd(name string, fn func() error) tea.Cmd {
	return func() tea.Msg {
		err := fn()
		if err != nil {
			Log(err.Error())
		}
		// CmdResultMessage is currently ignored in Update()
		return CmdResultMsg {}
	}
}
type CmdResultMsg struct {}

// Handles data fetching operations.
// Sends StateMsg with updated results.
func RefreshStateCmd() tea.Cmd {
	return func() tea.Msg {
		playing, _ := IsPlaying()
		song := GetCurrentSongObject()
		albums, _ := GetAlbums()
		playlists, _ := GetPlaylists()
		return stateMsg { 
			IsPlaying:  playing,
			CurrentSong: song,
			Albums:     albums,
			Playlists:  playlists,
		}
	}
}
// Contains updated state.
type stateMsg struct {	
	Albums      []string
	Playlists   []string
	CurrentSong	Song
	IsPlaying		bool
}

// Handles data fetching for playlist or album views
// Sends List struct to update the CurrentList
func UpdateListCmd(source Source, name string) tea.Cmd {
	return func() tea.Msg {
		var songs []Song
		var owner string
		var err error

		switch source {
		case Album:
			songs, err = GetSongsFromSource("album", name)
			owner = GetAlbumArtist(name) 
		case Playlist:
			songs, err = GetSongsFromSource("playlist", name)
			owner = "You"
		default:
			err = nil
			songs = nil
			owner = ""
		}
		if err != nil {
			Log(err.Error())
		}
		return ListMsg {
			Name: name,
			Owner: owner,
			Songs: songs,
		}
	}
}
// Wrapper type to indicate this List is sent as a Msg
type ListMsg List


/**
* Sync. functions that execute applescript.
* Used to fetch data and control Apple Music.
*/

// os/exec call
func run(command string) (string, error) {
	cmd := exec.Command("osascript", "-e", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		Log(err.Error())
		return "", err
	}
	return strings.TrimRight(out.String(), "\r\n"), nil
}
// controls
func TogglePlayPause() error {
	_, err := run(`tell application "Music" to playpause`)
	return err
}
func NextTrack() error {
	_, err := run(`tell application "Music" to next track`)
	return err
}
func PreviousTrack() error {
	_, err := run(`tell application "Music" to previous track`)
	return err
}
// song info
func GetCurrentSongTitle() (string, error) {
	return run(`tell application "Music" to get name of current track`)
}
func GetCurrentArtist() (string, error) {
	return run(`tell application "Music" to get artist of current track`)
}
func GetCurrentAlbum() (string, error) {
	return run(`tell application "Music" to get album of current track`)
}
func GetCurrentDuration() (string, error) {
	out, err := run(`tell application "Music" to get duration of current track`)
	if err != nil {
		return "?", err
	}
	return ParseDuration(out), nil
}
func IsPlaying() (bool, error) {
	state, err := run(`tell application "Music" to get player state`)
	if err != nil {
		return false, err
	}
	return state == "playing", nil
}
func IsPlayingToString(isPlaying bool) string {
	if isPlaying {
		return "playing"
	}
	return "stopped"
} 
func GetCurrentSongObject() (Song) {
	title, _ := GetCurrentSongTitle()
	artist, _ := GetCurrentArtist()
	duration, _ := GetCurrentDuration()
	res := Song {
		Title: title,
		Artist: artist,
		Duration: duration,
	}
	return res
}
// library info
func GetPlaylists() ([]string, error) {
	raw, err := run(`tell application "Music" to get name of playlists`)
	if err != nil {
		return nil, err
	}
	return strings.Split(raw, ", "), nil
}
func GetAlbums() ([]string, error) {
	raw, err := run(`tell application "Music" to get album of every track`)
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	var albums []string
	for album := range strings.SplitSeq(raw, ", ") {
		if !seen[album] && album != "" {
			seen[album] = true
			albums = append(albums, album)
		}
	}
	return albums, nil
}
func GetAlbumArtist(albumName string) string {
	script := fmt.Sprintf(`tell application "Music" to get artist of album "%s"`, albumName)
	artist, err := run(script)
	if err != nil {
		Log(err.Error())
		return ""
	}
	return artist
}
func GetSongsFromSource(sourceType, sourceName string) ([]Song, error) {
    // get {name, artist, duration} of every track of the source
    script := fmt.Sprintf(`
        tell application "Music"
            if "%s" = "playlist" then
                set trackInfo to get {name, artist, duration} of every track of playlist "%s"
            else if "%s" = "album" then
                set trackInfo to get {name, artist, duration} of every track whose album is "%s"
            else
                set trackInfo to {}
            end if
            return trackInfo
        end tell
    `, sourceType, sourceName, sourceType, sourceName)

    raw, err := run(script)
    if err != nil {
			Log(err.Error())
      return nil, err
    }

    raw = strings.Trim(raw, "{}")
    parts := strings.Split(raw, ", ")
		total := len(parts) 
		if total % 3 != 0 {
			return nil, fmt.Errorf("unexpected number of items: %d", total)
    }
		n := total / 3
		names := parts[:n]
    artists := parts[n : 2*n]
    durations := parts[2*n:]
    var songs []Song
    for i := 0; i < n; i++ {
        name := strings.Trim(names[i], `"`)
        artist := strings.Trim(artists[i], `"`)
        duration := ParseDuration(durations[i])
        songs = append(songs, Song{
            Title:     name,
            Artist:   artist,
            Duration: duration,
        })
    }
    return songs, nil
}
func ParseDuration(duration string) string {
	duration = strings.TrimSpace(duration)
	out, err := strconv.ParseFloat(duration, 32)
	if err != nil {
		Log(err.Error())
		return "?"
	}
	s := int(out)
	mins := s / 60
	secs := s % 60
	return fmt.Sprintf("%d:%02d", mins, secs)
}
func GetSongs() ([]string, error) {
	raw, err := run(`tell application "Music" to get name of every track of library playlist 1`)
	if err != nil {
		return nil, err
	}
	songs := strings.Split(raw, ", ")
	return songs, nil
}
// library playback
func PlayPlaylist(name string) error {
	_, err := run(`tell application "Music" to play playlist "` + name + `"`)
	return err
}
func PlayAlbum(name string) error {
	_, err := run(`
		tell application "Music"
			set theTracks to every track whose album is "` + name + `"
			if (count of theTracks) > 0 then
				play item 1 of theTracks
			end if
		end tell`)
	return err
}
