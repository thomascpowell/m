package scripts

import (
	"m/utils"
)

/**
* Functions to get the current state of the music player.
*/


func GetCurrentSongObject() (utils.Song) {
	title, _ := GetCurrentSongTitle()
	artist, _ := GetCurrentArtist()
	duration, _ := GetCurrentDuration()
	res := utils.Song {
		Title: title,
		Artist: artist,
		Duration: duration,
	}
	return res
}

func GetCurrentSongTitle() (string, error) {
	return Run(`tell application "Music" to get name of current track`)
}

func GetCurrentArtist() (string, error) {
	return Run(`tell application "Music" to get artist of current track`)
}

func GetCurrentAlbum() (string, error) {
	return Run(`tell application "Music" to get album of current track`)
}

func GetCurrentDuration() (string, error) {
	out, err := Run(`tell application "Music" to get duration of current track`)
	if err != nil {
		return "?", err
	}
	return ParseDuration(out), nil
}

func IsPlaying() (bool, error) {
	state, err := Run(`tell application "Music" to get player state`)
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
