package app

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"m/lists"
	"m/scripts"
	"m/utils"
	"m/views"
	"time"
)

/**
* Contains the update loop and associated functions.
 */

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case scripts.LibraryMsg:
		return m.handleLibraryMsg(msg)
	case scripts.StateMsg:
		return m.handleStateMsg(msg)
	case tickMsg:
		return m.handleTickMsg()
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)
	case scripts.ChangeViewMsg:
		return m.handleChangeViewMsg(msg)
	}
	var cmd tea.Cmd
	m.UIList, cmd = m.UIList.Update(msg)
	return m, cmd
}

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	if !m.Loaded {
		return *m, nil
	}
	if m.CurrentView == utils.Menu {
		m.UIList.SetSize(msg.Width, msg.Height-5)
		return *m, nil
	}
	m.UIList.SetSize(msg.Width, msg.Height)
	return *m, nil
}

func (m *Model) handleChangeViewMsg(msg scripts.ChangeViewMsg) (tea.Model, tea.Cmd) {
	m.CurrentView = msg.View
	m.UIList = msg.List
	m.Loaded = true
	return *m, nil
}

func TickCmd() tea.Cmd {
	return tea.Tick(time.Second*5, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

type tickMsg struct{}

func (m Model) handleStateMsg(msg scripts.StateMsg) (tea.Model, tea.Cmd) {
	m.CurrentSong = msg.CurrentSong
	m.IsPlaying = msg.IsPlaying
	return m, nil
}

func (m Model) handleTickMsg() (tea.Model, tea.Cmd) {
	return m, tea.Batch(TickCmd(), scripts.RefreshStateCmd())
}

func (m Model) handleLibraryMsg(msg scripts.LibraryMsg) (tea.Model, tea.Cmd) {
	m.Library = utils.Library{
		Songs:     msg.Songs,
		Albums:    msg.Albums,
		Playlists: msg.Playlists,
	}
	utils.Log(fmt.Sprintf("handleLibraryMsg: songs=%d albums=%d playlists=%d",
		len(msg.Songs), len(msg.Albums), len(msg.Playlists)))
	return m, nil
}

func containsUIList(view utils.View) bool {
	// currently, all views use UIList
	return true
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.UIList.FilterState() == list.Filtering {
		var cmd tea.Cmd
		m.UIList, cmd = m.UIList.Update(msg)
		return m, cmd
	}
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "b", "x":
		return m, scripts.ChangeViewCmd(utils.Menu, views.NewMenuList())
	case "enter":
		if m.CurrentView == utils.Menu {
			return m.handleMenuSelect()
		}
		return m.handleSelect()
	}
	var cmd tea.Cmd
	m.UIList, cmd = m.UIList.Update(msg)
	return m, cmd
}

func (m Model) handleSelect() (tea.Model, tea.Cmd) {
	selected := m.UIList.SelectedItem()
	item, ok := selected.(lists.ListItem)
	if !ok {
		return m, nil
	}
	title := item.Name
	desc := item.Desc
	switch m.CurrentView {
	case utils.Albums:
		songs := scripts.GetSongsFromSource(utils.Album, title, m.Library)
		return m, scripts.ChangeViewCmd(utils.AlbumDetail, views.NewAlbumDetailList(songs, title, item.Desc))
	case utils.Playlists:
		songs := scripts.GetSongsFromSource(utils.Playlist, title, m.Library)
		return m, scripts.ChangeViewCmd(utils.PlaylistDetail, views.NewPlaylistDetailList(songs, title, item.Desc))
	case utils.PlaylistDetail:
		if item.Id == "PLAY_ALL" {
			scripts.PlayPlaylist(desc)
		} else {
			scripts.SelectTrack(item.Id)
		}
	case utils.AlbumDetail:
		if item.Id == "PLAY_ALL" {
			scripts.PlaySongList(scripts.GetSongsFromSource(utils.Album, desc, m.Library))
		} else {
			scripts.SelectTrack(item.Id)
		}
	}
	return m, nil
}

func (m Model) handleMenuSelect() (tea.Model, tea.Cmd) {
	selected := m.UIList.SelectedItem()
	item, ok := selected.(lists.BaseListItem)
	if !ok {
		return m, nil
	}
	action := item.Action
	utils.Log("MENUACTION: " + action)
	switch action {
	case "PLAY_PAUSE":
		m.IsPlaying = !m.IsPlaying
		return m, scripts.RunAsCmd("PLAY_PAUSE", scripts.TogglePlayPause)
	case "SKIP":
		return m, scripts.RunAsCmd("SKIP", scripts.NextTrack)
	case "SHOW_ALBUMS":
		return m, scripts.ChangeViewCmd(utils.Albums, views.NewAlbumList(m.Library.Albums))
	case "SHOW_PLAYLISTS":
		return m, scripts.ChangeViewCmd(utils.Playlists, views.NewPlaylistList(m.Library.Playlists))
	}
	return m, nil
}
