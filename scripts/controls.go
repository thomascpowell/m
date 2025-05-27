package scripts

import (
	"fmt"
	"m/utils"
	"strings"
)

/**
* AppleScript wrappers for controlling Apple Music.
*/


func TogglePlayPause() error {
	_, err := Run(`tell application "Music" to playpause`)
	return err
}

func NextTrack() error {
	_, err := Run(`tell application "Music" to next track`)
	return err
}

func PreviousTrack() error {
	_, err := Run(`tell application "Music" to previous track`)
	return err
}

func PlayPlaylist(name string) error {
	_, err := Run(`tell application "Music" to play playlist "` + name + `"`)
	return err
}

func PlaySongList(songs []utils.Song) error {
	utils.Log(fmt.Sprintf("songs count: %d", len(songs)))
	var script strings.Builder
	script.WriteString(`
	tell application "Music"
		if exists playlist "temporary" then
				delete playlist "temporary"
		end if
		set temporary to make new playlist with properties {name:"temporary"}`)
	for _, song := range songs {
		script.WriteString(fmt.Sprintf("\n" + `duplicate (some track whose persistent ID is "%s") to temporary`, song.SongId))
	}
	script.WriteString(`
		play temporary
	end tell`)
	utils.Log(script.String())
	_, err := Run(script.String())
	return err
}

func SelectTrack(id string) error {
	script := fmt.Sprintf(`
		tell application "Music"
			set t to (some track whose persistent ID is "%s")
			play t
		end tell
	`, id)
	_, err := Run(script)	
	return err
}

func ToggleShuffle(enable bool) error {
	state := "false"
	if enable {
		state = "true"
	}
	_, err := Run(`
		tell application "Music"
			set shuffle to ` + state + `
			set shuffle mode to songs
		end tell`)
	return err
}

