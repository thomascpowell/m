package main

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	switch m.CurrentView {
	case BaseView:
		return baseView(m)
	case AlbumsView:
		return sourcesView("albums", m)
	case PlaylistsView:
		return sourcesView("playlists", m)
	case SourceDetailView:
		return sourceDetailView(m)
	default:
		return "unknown"
	}
}

// song info, controls
func baseView(m model) string {
	if m.CurrentSong.Title == "" {
		return "not playing"
	}
	return fmt.Sprintf(
		"%s — %s\n(%v)\n\n(a: albums, p: playlists)",
		m.CurrentSong.Title,
		m.CurrentSong.Artist,
		IsPlayingToString(m.IsPlaying),
	)
}

// 
func sourcesView(title string, m model) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s:\n\n", title))
	var items []string
	if title == "albums" {
		items = m.Albums
	} else {
		items = m.Playlists
	}
	if len(items) == 0 {
		b.WriteString(fmt.Sprintf("no %s loaded", title))
	} else {
		for i, item := range items {
			b.WriteString(fmt.Sprintf("  %d. %s\n", i+1, item))
		}
	}
	b.WriteString("\n(b: back)")
	return b.String()
}

// shows the detail view of whatever is in CurrentList
// this will be different (album, playlist) depending
// on what the user seleted from the sourceView
func sourceDetailView(m model) string {
	var b strings.Builder
	if len(m.CurrentList.Songs) == 0 {
		b.WriteString(fmt.Sprintf("no songs in \"%s\"", m.CurrentList.Name))
	} else {
		for i, song := range m.CurrentList.Songs {
			b.WriteString(fmt.Sprintf("  %d. %s — %s (%s)\n", i+1, song.Title, song.Artist, song.Duration))
		}
	}
	b.WriteString("\n(unfinished: instructions and binding)\n")
	return b.String()
}

