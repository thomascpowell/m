package utils

// Represents a source.
// Can be an album or playlist.
// Does not list the contents.
// Used for displaying options.
type Source struct {
	Title 		string
	Artist 		string
}

// Represents a Song.
type Song struct {
	Title			string
	Artist 		string
	Duration	string
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

