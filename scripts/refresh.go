package scripts

import (
	"m/utils"
	"strings"
)

/**
* Functions to get the current state of the music player.
*/


func GetPlayerState() (utils.Song, bool) {
	raw, err := Run(`tell application "Music"
		set t to name of current track
		set a to artist of current track
		set d to duration of current track
		set p to player state
		return t & "|||" & a & "|||" & d & "|||" & p
	end tell`)
	if err != nil {
		utils.Log("GetCurrentSongObject: AppleScript Error - " + err.Error())
		return utils.Song{}, false
	}
	parts := strings.Split(raw, "|||")
	if len(parts) != 4 {
		utils.Log("GetCurrentSongObject: Malformed Response")
		return utils.Song{}, false
	}
	current_song := utils.Song{
		Title:    strings.TrimSpace(parts[0]),
		Artist:   strings.TrimSpace(parts[1]),
		Duration: ParseDuration(strings.TrimSpace(parts[2])),
	}
	player_state := parts[3] == "playing"
	return current_song, player_state
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
