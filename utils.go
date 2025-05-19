package main

import (
	"bytes"
	"os/exec"
	"strings"
	"strconv"
)

// run command
func run(command string) (string, error) {
	cmd := exec.Command("osascript", "-e", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// controls
func Play() error {
	_, err := run(`tell application "Music" to play`)
	return err
}
func Pause() error {
	_, err := run(`tell application "Music" to pause`)
	return err
}
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
func GetCurrentDuration() (float32, error) {
	out, err := run(`tell application "Music" to get duration of current track`)
	if err != nil {
		return 0, err
	}
	out = strings.TrimSpace(out)
	parsed, err := strconv.ParseFloat(out, 32)
	if err != nil {
		return 0, err
	}
	return float32(parsed), nil
}
func IsPlaying() (bool, error) {
	state, err := run(`tell application "Music" to get player state`)
	if err != nil {
		return false, err
	}
	return state == "playing", nil
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
