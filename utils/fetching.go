package utils

import (
	"bytes"
	"os/exec"
	"strings"
	"strconv"
	"fmt"
)

/**
* Sync. functions that execute AppleScript.
* Used to fetch data and control Apple Music.
*/

// Generic os call
func run(command string) (string, error) {
	cmd := exec.Command("osascript", "-e", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(out.String(), "\r\n"), nil
}

// Functions for controlling music.
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
func PlayTrack(trackName string) error {
	script := fmt.Sprintf(`
		tell application "Music"
			set t to first track whose name is "%s"
			play t
		end tell
	`, trackName)
	_, err := run(script)
	return err
}

// Functions for fetching song info.
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

// Returns Song struct.
// Uses song info functions.
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

// Functions for getting playlists and albums.
func GetPlaylists() ([]string, error) {
	raw, err := run(`tell application "Music" to get name of playlists`)
	if err != nil {
		return nil, err
	}
	return strings.Split(raw, ", "), nil
}
func ParseDuration(duration string) string {
	duration = strings.TrimSpace(duration)
	out, err := strconv.ParseFloat(duration, 32)
	if err != nil {
		Log("ParseDuration:" + err.Error())
		return "?"
	}
	s := int(out)
	mins := s / 60
	secs := s % 60
	return fmt.Sprintf("%d:%02d", mins, secs)
}

// Functions for fetching music sources.
func FetchSources() ([]Source, []Source) {
	albumScript := `
	tell application "Music"
  set albumData to {}
  repeat with t in tracks
    set albumName to album of t
    set artistName to artist of t
    if albumName is not "" then
      copy (albumName & "||" & artistName) to end of albumData
    end if
  end repeat
  set AppleScript's text item delimiters to "\n"
  return albumData as string
	end tell`
	rawAlbums, err := run(albumScript)
	if err != nil {
		return nil, nil
	}
	albumLines := strings.Split(strings.Trim(rawAlbums, "{}"), "\n")
	albumMap := map[string]string{}
	for _, line := range albumLines {
		parts := strings.Split(line, "||")
		name := parts[0]
		artist := parts[1]
		albumMap[name] = artist
	}
	albums := make([]Source, 0, len(albumMap))
	for name, artist := range albumMap {
		albums = append(albums, Source{
			Title:   name,
			Artist: artist,
		})
	}
	playlistScript := `tell application "Music" to get name of playlists`
	rawPlaylists, err := run(playlistScript)
	if err != nil {
		Log("error in getting playlists")
		return nil, nil
	}
	playlistLines := strings.Split(rawPlaylists, ", ")
	playlists := make([]Source, 0, len(playlistLines))
	for _, name := range playlistLines {
		playlists = append(playlists, Source{
			Title:   name,
			Artist: "you", 
		})
	}
	return albums, playlists
}
func GetSongsFromSource(sourceType, sourceName string) ([]Song, error) {
	var script string
	switch sourceType {
	case "playlist":
		script = fmt.Sprintf(`
			tell application "Music"
				set songData to {}
				if exists playlist "%s" then
					repeat with t in tracks of playlist "%s"
						set titleName to name of t
						set artistName to artist of t
						set durationVal to duration of t
						copy (titleName & "||" & artistName & "||" & durationVal) to end of songData
					end repeat
				end if
				set AppleScript's text item delimiters to "\n"
				return songData as string
			end tell
		`, sourceName, sourceName)
	case "album":
		script = fmt.Sprintf(`
			tell application "Music"
				set songData to {}
				repeat with t in tracks
					if album of t as string is "%s" then
						set titleName to name of t
						set artistName to artist of t
						set durationVal to duration of t
						copy (titleName & "||" & artistName & "||" & durationVal) to end of songData
					end if
				end repeat
				set AppleScript's text item delimiters to "\n"
				return songData as string
			end tell
		`, sourceName)
	default:
		return nil, fmt.Errorf("unknown source type: %s", sourceType)
	}
	raw, err := run(script)
	if err != nil {
		Log("AppleScript error: " + err.Error())
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	var songs []Song
	for _, line := range lines {
		parts := strings.Split(line, "||")
		if len(parts) != 3 {
			Log("Malformed line: " + line)
			continue
		}
		title := strings.TrimSpace(parts[0])
		artist := strings.TrimSpace(parts[1])
		duration := ParseDuration(parts[2])
		songs = append(songs, Song{
			Title:    title,
			Artist:   artist,
			Duration: duration,
		})
	}
	return songs, nil
}
