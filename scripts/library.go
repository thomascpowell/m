package scripts

import (
	"strings"
	"strconv"
	"m/utils"
	"fmt"
	"os"
	"encoding/gob"
)

/**
* Functions for fetching Library data.
*/


// For initial fetching of library data
func GetLibraryData() (*utils.Library, error) {
	lib := &utils.Library{}
	songs, err := getSongs()
	if err != nil {
		return nil, err
	}
	albums, err := computeAlbums(songs)
	if err != nil {
		return nil, err
	}
	playlists, err := getPlaylists()
	if err != nil {
		return nil, err
	}
	lib.Songs = songs
	lib.Albums = albums
	lib.Playlists = playlists
	return lib, err
}

// Get all songs in library.
func getSongs() ([]utils.Song, error) {
	raw, err := Run(`
	tell application "Music"
		set output to {}
		repeat with t in tracks
			set end of output to (persistent ID of t as string) & " ||| " & (name of t as string) & " ||| " & (artist of t as string) & " ||| " & (duration of t as string) & " ||| " & (album of t as string)
		end repeat
		set AppleScript's text item delimiters to linefeed
		return output as string
	end tell`)
	var songs []utils.Song
	if err != nil {
		utils.Log("getSongs: AS error")
		return songs, err
	}
	lines := strings.SplitSeq(strings.TrimSpace(raw), "\n")
	for line := range lines {
		parts := strings.Split(line, " ||| ")
		if len(parts) != 5 {
			utils.Log("getSongs: Malformed response detected.")
			continue
		}
		song := utils.Song{
			SongId: parts[0],
			Title: parts[1],
			Artist: parts[2],
			Duration: ParseDuration(parts[3]),
			Album: parts[4],
		}
		songs = append(songs, song)
	}
	return songs, nil
}

// Get all playlists in library.
func getPlaylists() ([]utils.Source, error) {
	raw, err := Run(`
	tell application "Music"
			set output to {}
			repeat with p in playlists
					set end of output to (name of p as string)
			end repeat
			set AppleScript's text item delimiters to linefeed
			return output as string
	end tell`)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(raw, "\n")
  res := make([]utils.Source, 0, len(lines))
	for _, l := range lines {
		res = append(res, utils.Source {
			Title: l,
			Artist: "You",
		})
	}
	return res, nil
}

// Finds all unique albums based on songs in the library.
// NOTE: Caching all songs makes this faster than using AppleScript here
func computeAlbums(songs []utils.Song) ([]utils.Source, error) {
	albumToArtist := make(map[string]string) // title to artist
	for _, song := range songs {
			if _, exists := albumToArtist[song.Album]; !exists {
					albumToArtist[song.Album] = song.Artist
			}
	}
	albums := make([]utils.Source, 0, len(albumToArtist))
	for album, artist := range albumToArtist {
			albums = append(albums, utils.Source{
					Title:   album,
					Artist: artist,
			})
	}
	return albums, nil
}

func GetSongsFromSource(kind utils.SourceType, source utils.Source, library utils.Library) []utils.Song {
	var result []utils.Song
	switch kind {
	case utils.Album:
		for _, song := range library.Songs {
			if song.Album == source.Title /* && song.Artist == source.Artist */ {
				result = append(result, song)
			}
		}
	case utils.Playlist:
		songs, err := getSongsFromPlaylist(source.Title, library)
		if err == nil {
			result = songs
		} else {
			utils.Log("failed to fetch playlist songs: " + err.Error())
		}
	}
	return result
}

func getSongsFromPlaylist(playlistName string, library utils.Library) ([]utils.Song, error) {
	script := fmt.Sprintf(`
	tell application "Music"
		if not (exists playlist "%s") then return ""
		set output to {}
		repeat with t in tracks of playlist "%s"
			set end of output to persistent ID of t as string
		end repeat
		set AppleScript's text item delimiters to linefeed
		return output as string
	end tell`, playlistName, playlistName)
	raw, err := Run(script)
	if err != nil {
		return nil, err
	}
	pids := strings.Split(strings.TrimSpace(raw), "\n")
	pidSet := make(map[string]struct{}, len(pids))
	for _, pid := range pids {
		pidSet[pid] = struct{}{}
	}
	var result []utils.Song
	for _, song := range library.Songs {
		if _, exists := pidSet[song.SongId]; exists {
			result = append(result, song)
		}
	}
	return result, nil
}

// Converts a string of digits to a readable format.
func ParseDuration(duration string) string {
	duration = strings.TrimSpace(duration)
	out, err := strconv.ParseFloat(duration, 32)
	if err != nil {
		utils.Log("ParseDuration:" + err.Error())
		return "?"
	}
	s := int(out)
	mins := s / 60
	secs := s % 60
	return fmt.Sprintf("%d:%02d", mins, secs)
}

// saves the Library object to improve load time
func SaveLibrary(library *utils.Library, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(library)
}

// loads library object from disk
func LoadLibrary(path string) (*utils.Library, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	var lib utils.Library
	err = decoder.Decode(&lib)
	if err != nil {
		return nil, err
	}
	return &lib, nil
}
