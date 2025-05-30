package views

import (
	"m/scripts"
	"m/utils"
	"time"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Contains the update loop and associated functions.
*/


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case scripts.LibraryMsg:		
		utils.Log("Recieved a LibraryMsg")
		return m.handleLibraryMsg(msg)
	case scripts.StateMsg:
		return m.handleStateMsg(msg)
	case tickMsg:
		return m.handleTickMsg(msg)
	case scripts.ListMsg:
		return m.handleListMsg(msg)
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case tea.WindowSizeMsg:
		if containsUIList(m.CurrentView) {
			m.UIList.SetSize(msg.Width, msg.Height)
		}
	}
	if containsUIList(m.CurrentView) {
		var cmd tea.Cmd
		m.UIList, cmd = m.UIList.Update(msg)
		return m, cmd
	}
	return m, nil
}

func TickCmd() tea.Cmd {
	return tea.Tick(time.Second*3, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
type tickMsg struct{}

func (m Model) handleLibraryMsg(msg scripts.LibraryMsg) (tea.Model, tea.Cmd) {
	m.Library = utils.Library {
		Songs: msg.Songs,
		Albums: msg.Albums,
		Playlists: msg.Playlists,
	}
	utils.Log(fmt.Sprintf("handleLibraryMsg: songs=%d albums=%d playlists=%d", 
		len(msg.Songs), len(msg.Albums), len(msg.Playlists)))
	return m, nil
}

func (m Model) handleStateMsg(msg scripts.StateMsg) (tea.Model, tea.Cmd) {
	m.CurrentSong = msg.CurrentSong
	m.IsPlaying = msg.IsPlaying
	return m, nil
}

func (m Model) handleTickMsg(msg tickMsg) (tea.Model, tea.Cmd) {
	return m, tea.Batch(TickCmd(), scripts.RefreshStateCmd())
}

func (m Model) handleListMsg(msg scripts.ListMsg) (tea.Model, tea.Cmd) {
	m.CurrentList = utils.List(msg)
	m.UIList = NewDetailList(m.CurrentList.Songs, m.CurrentList.Name, m.CurrentList.Owner, m.DetailSource)
	m.CurrentView = SourceDetailView
	return m, nil
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.UIList.FilterState() == list.Filtering {
		var cmd tea.Cmd
		m.UIList, cmd = m.UIList.Update(msg)
		return m, cmd
	}
	switch msg.String() {
	// generic keybinds
	case "ctrl+c", "q":
		return m, tea.Quit
	case " ":
		m.IsPlaying = !m.IsPlaying 
		return m, scripts.RunAsCmd("toggle", scripts.TogglePlayPause)
	case "a":
		m.UIList = NewSourceList(m.Library.Albums, "albums")
		m.CurrentView = AlbumsView
		return m, nil
	case "p":
		m.UIList = NewSourceList(m.Library.Playlists, "playlists")
		m.CurrentView = PlaylistsView
		return m, nil
	// "back" or "base"
	case "b":
		m.CurrentView = BaseView
		return m, nil
	// select item
	case "enter":
		return m.handleSelect()
	}
	// pass other keypresses to UIList
	var cmd tea.Cmd
	m.UIList, cmd = m.UIList.Update(msg)
	return m, cmd
}

func (m Model) handleSelect() (tea.Model, tea.Cmd) {
	selected := m.UIList.SelectedItem()
	item, ok := selected.(utils.ListItem)
	if !ok {
		return m, nil
	}
	s := utils.Source {
		Title: item.Name,
		Artist: item.Desc,
	}
	switch m.CurrentView {
	case SourceDetailView:
		return m, scripts.RunAsCmd("play track", func() error {
			utils.Log(item.Id)
			if item.Id == "PLAY_PLAYLIST" {
				return scripts.PlayPlaylist(s.Title)
			} else if item.Id == "PLAY_ALBUM" {
				return scripts.PlaySongList(scripts.GetSongsFromSource(utils.Album, s, m.Library))
			} else {
				return scripts.SelectTrack(item.Id)
			}
		})
	case AlbumsView:
		m.DetailSource = utils.Album
		return m, scripts.UpdateListCmd(utils.Album, s, m.Library)
	case PlaylistsView:
		m.DetailSource = utils.Playlist
		return m, scripts.UpdateListCmd(utils.Playlist, s, m.Library)
	default:
		break
	}
	return m, nil
}



func containsUIList(view View) bool {
	return view == AlbumsView || 
	view == PlaylistsView || 
	view == SourceDetailView 
}


