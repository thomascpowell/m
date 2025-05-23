package scripts

import (
	"fmt"
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

func PlayAlbum(name string) error {
	_, err := Run(`
		tell application "Music"
			set theTracks to every track whose album is "` + name + `"
			if (count of theTracks) > 0 then
				play item 1 of theTracks
			end if
		end tell`)
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


