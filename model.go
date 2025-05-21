package main

import (
	"m/utils"

	"time"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Contains the application model and update loop.
*/

type View int
const (
	BaseView View = iota
	AlbumsView
	PlaylistsView
	SourceDetailView
)

// Represents the application state.
type model struct {
	Albums				[]utils.Source
	Playlists			[]utils.Source
	CurrentSong		utils.Song
	IsPlaying			bool
	CurrentView		View
	CurrentList		utils.List
	UIList 				list.Model
}

// starts the event loop
func (m model) Init() tea.Cmd {
	return tea.Batch(utils.FetchSourcesCmd(), utils.RefreshStateCmd(), tickCmd())
}

// sends a tickMsg after 3 seconds
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
type tickMsg struct{}

// update loop
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// initial sources load
	case utils.SourcesMsg:
		m.Albums = msg.Albums
		m.Playlists = msg.Playlists
		return m, nil

	// state update
	case utils.StateMsg:
		m.CurrentSong = msg.CurrentSong
		m.IsPlaying = msg.IsPlaying
		return m, nil

	// tick occurs
	case tickMsg:
		return m, tea.Batch(tickCmd(), utils.RefreshStateCmd())

	// Pressing enter on a Source sends a Cmd that returns a ListMsg.
	// When it arrives, the view changes to detail view with that data.
	case utils.ListMsg:
		m.CurrentList = utils.List(msg)
		m.UIList = NewDetailList(m.CurrentList.Songs, m.CurrentList.Name, m.CurrentList.Owner)
		m.CurrentView = SourceDetailView
		return m, nil

	// keypress
	case tea.KeyMsg:

		// user types a key but is in a filter
		if m.UIList.FilterState() == list.Filtering {
			var cmd tea.Cmd
			m.UIList, cmd = m.UIList.Update(msg)
			return m, cmd
		}

		// user types a key in any other context
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			m.IsPlaying = !m.IsPlaying 
			return m, utils.RunAsCmd("toggle", utils.TogglePlayPause)
		case "a":
			m.UIList = NewSourceList(m.Albums, "albums")
			m.CurrentView = AlbumsView
			return m, nil
		case "p":
			m.UIList = NewSourceList(m.Playlists, "playlists")
			m.CurrentView = PlaylistsView
			return m, nil
		case "b":
			m.CurrentView = BaseView
			return m, nil
		case "enter":
			selected := m.UIList.SelectedItem()
			if item, ok := selected.(ListItem); ok {
					switch m.CurrentView {
					case SourceDetailView:
							// user selected a track, play it
							utils.Log("PlayTrack triggered for: " + item.Name)
							return m, utils.RunAsCmd("play track", func() error {
									return utils.PlayTrack(item.Name)
							})
					case AlbumsView:
							// user selected an album, show album tracks
							return m, utils.UpdateListCmd(utils.Album, item.Name, item.Desc)
					case PlaylistsView:
							// user selected a playlist, show playlist tracks
							return m, utils.UpdateListCmd(utils.Playlist, item.Name, item.Desc)
					default:
							return m, nil
					}
			}
		}
	}
	// send keypresses to UIList
  if m.CurrentView == AlbumsView || 
	m.CurrentView == PlaylistsView ||
	m.CurrentView == SourceDetailView {
		var cmd tea.Cmd
		m.UIList, cmd = m.UIList.Update(msg)
		return m, cmd
  }
	// default
	return m, nil
}

