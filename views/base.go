package views

import (
	"m/scripts"
	"fmt"

)
/**
* Base View.
* Shows the current player state.
*/

func ShowBaseView(m Model) string {
	if m.CurrentSong.Title == "" {
		return "not playing"
	}
	return fmt.Sprintf(
		"%s — %s\n(%s)\n\n(a: albums, p: playlists)",
		m.CurrentSong.Title,
		m.CurrentSong.Artist,
		scripts.IsPlayingToString(m.IsPlaying),
	)
}

