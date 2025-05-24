package views

import (
	"m/scripts"
	"m/utils"

	"time"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

/**
* Contains the application Model and update loop.
*/

type View int
const (
	BaseView View = iota
	AlbumsView
	PlaylistsView
	SourceDetailView
)

type Model struct {
	Library				utils.Library
	CurrentSong		utils.Song
	IsPlaying			bool
	CurrentView		View
	CurrentList		utils.List
	UIList 				list.Model
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(scripts.GetLibraryCmd(), scripts.RefreshStateCmd(), tickCmd())
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*3, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
type tickMsg struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case scripts.LibraryMsg:
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


func (m Model) handleLibraryMsg(msg scripts.LibraryMsg) (tea.Model, tea.Cmd) {
	m.Library = utils.Library {
		Songs: msg.Songs,
		Albums: msg.Albums,
		Playlists: msg.Playlists,
	}
	return m, nil
}

func (m Model) handleStateMsg(msg scripts.StateMsg) (tea.Model, tea.Cmd) {
	m.CurrentSong = msg.CurrentSong
	m.IsPlaying = msg.IsPlaying
	return m, nil
}

func (m Model) handleTickMsg(msg tickMsg) (tea.Model, tea.Cmd) {
	return m, tea.Batch(tickCmd(), scripts.RefreshStateCmd())
}

func (m Model) handleListMsg(msg scripts.ListMsg) (tea.Model, tea.Cmd) {
	m.CurrentList = utils.List(msg)
	m.UIList = NewDetailList(m.CurrentList.Songs, m.CurrentList.Name, m.CurrentList.Owner)
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
	case "b":
		m.CurrentView = BaseView
		return m, nil
	case "enter":
		selected := m.UIList.SelectedItem()
		item, ok := selected.(utils.ListItem)
		if !ok {
			break
		}
		s := utils.Source {
			Title: item.Name,
			Artist: item.Desc,
		}
		switch m.CurrentView {
		case SourceDetailView:
			return m, scripts.RunAsCmd("play track", func() error {
				utils.Log(item.Id)
				return scripts.SelectTrack(item.Id)
			})
		case AlbumsView:
			return m, scripts.UpdateListCmd(utils.Album, s, m.Library)
		case PlaylistsView:
			return m, scripts.UpdateListCmd(utils.Playlist, s, m.Library)
		default:
			break
		}
	}
	var cmd tea.Cmd
	m.UIList, cmd = m.UIList.Update(msg)
	return m, cmd
}

func containsUIList(view View) bool {
	return view == AlbumsView || 
	view == PlaylistsView || 
	view == SourceDetailView 
}


