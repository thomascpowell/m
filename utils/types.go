package utils

import (
	// "github.com/charmbracelet/bubbles/list"
)

/**
* Contains types used throughout.
*/



// Represents a source.
// Can be an album or playlist.
// Does not list the contents.
type Source struct {
	Title 		string
	Artist 		string
}

// Represents a Song.
type Song struct {
	Title			string
	Artist 		string
	Duration	string
	Album			string
	SongId		string
}

// Used in the model to store all fetched information.
type Library struct {
    Songs     []Song
    Albums    []Source
    Playlists []Source
}

// Represents a list of Songs
// Can be an album or a playlist
// Used when viewing one of the above.
type List struct {
	Name 		string
	Owner		string
	Songs 	[]Song
}

// Represents the possible sources
// (Playlist and Album)
type SourceType int
const (
	Album = iota
	Playlist
)

// Represents the possible UI Views
type View int
const (
	Blank View = iota
	Menu
	Albums
	Playlists
	AlbumDetail
	PlaylistDetail
)

