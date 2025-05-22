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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case utils.SourcesMsg:
		return m.handleSourcesMsg(msg)
	case utils.StateMsg:
		return m.handleStateMsg(msg)
	case tickMsg:
		return m.handleTickMsg(msg)
	case utils.ListMsg:
		return m.handleListMsg(msg)
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	default:
		return m, nil
	}
}

func (m model) handleSourcesMsg(msg utils.SourcesMsg) (tea.Model, tea.Cmd) {
	m.Albums = msg.Albums
	m.Playlists = msg.Playlists
	return m, nil
}

func (m model) handleStateMsg(msg utils.StateMsg) (tea.Model, tea.Cmd) {
	m.CurrentSong = msg.CurrentSong
	m.IsPlaying = msg.IsPlaying
	return m, nil
}

func (m model) handleTickMsg(msg tickMsg) (tea.Model, tea.Cmd) {
	return m, tea.Batch(tickCmd(), utils.RefreshStateCmd())
}

func (m model) handleListMsg(msg utils.ListMsg) (tea.Model, tea.Cmd) {
	m.CurrentList = utils.List(msg)
	m.UIList = NewDetailList(m.CurrentList.Songs, m.CurrentList.Name, m.CurrentList.Owner)
	m.CurrentView = SourceDetailView
	return m, nil
}

func (m model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.UIList.FilterState() == list.Filtering {
		utils.Log("Setting Filter")
		var cmd tea.Cmd
		m.UIList, cmd = m.UIList.Update(msg)
		return m, cmd
	}
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
		item, ok := selected.(ListItem)
		if !ok {
			break
		}
		switch m.CurrentView {
		case SourceDetailView:
			utils.Log("PlayTrack triggered for: " + item.Name)
			return m, utils.RunAsCmd("play track", func() error {
				return utils.PlayTrack(item.Name)
			})
		case AlbumsView:
			return m, utils.UpdateListCmd(utils.Album, item.Name, item.Desc)
		case PlaylistsView:
			return m, utils.UpdateListCmd(utils.Playlist, item.Name, item.Desc)
		default:
			// return m, nil
			break
		}
	}
	var cmd tea.Cmd
	m.UIList, cmd = m.UIList.Update(msg)
	return m, cmd
}



