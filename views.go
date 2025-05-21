package main

import (
	"m/utils"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)


/**
* bubbles/list configuration.
*/

// Represents a m.UIList item
// Either a playlist or Album
// Implements? required bubbles/list functions
type ListItem struct {
    Name string
    Desc string
}
func (a ListItem) Title() string {
    return a.Name
}
func (a ListItem) Description() string {
    return a.Desc
}
func (a ListItem) FilterValue() string {
    return a.Name
}


/**
* tea View() function.
*/

func (m model) View() string {
	switch m.CurrentView {
	case BaseView:
		return baseView(m)
	case AlbumsView:
		return sourcesView(m)
	case PlaylistsView:
		return sourcesView(m)
	case SourceDetailView:
		return sourceDetailView(m)
	default:
		return "unknown"
	}
}


/**
* Base View
*/

func baseView(m model) string {
	if m.CurrentSong.Title == "" {
		return "not playing"
	}
	return fmt.Sprintf(
		"%s — %s\n(%v)\n\n(a: albums, p: playlists)",
		m.CurrentSong.Title,
		m.CurrentSong.Artist,
		utils.IsPlayingToString(m.IsPlaying),
	)
}


/**
* Sources view.
* Used for AlbumsView and PlaylistsView.
*/

func NewSourceList(sources []utils.Source, name string) list.Model {
	items := make([]list.Item, len(sources))
	for i, source := range sources {
			items[i] = ListItem {
					Name: 	source.Title,
					Desc:   source.Artist,
			}
	}
	const width = 50
	const height = 30
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = name
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	return l
}

// Returns the UIList view.
// m.CurrentList stores data used here.
func sourcesView(m model) string {
	return m.UIList.View()
}


/**
* Source detail view.
* Shows the content of a playlist or album.
*/


func NewDetailList(songs []utils.Song, name string, artist string) list.Model {
	items := make([]list.Item, len(songs))
	for i, source := range songs {
			items[i] = ListItem {
					Name: 	source.Title,
					Desc:   source.Artist + " • " + source.Duration,
			}
	}
	const width = 50
	const height = 30
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = name + " • " + artist
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
  return l
}

// shows the detail view of whatever is in CurrentList
// this will be different (album, playlist) depending
// on what the user seleted from the sourceView
func sourceDetailView(m model) string {
	return m.UIList.View()
}

