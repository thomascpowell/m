package app

import (
	"m/scripts"
	"m/utils"
	"time"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
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
		if containsUIList(m.CurrentView) && m.Loaded {
			m.UIList.SetSize(msg.Width, msg.Height-4)
		}
		return m, nil
	case scripts.ChangeViewMsg:
		return m.handleChangeViewMsg(msg)
	}

	




	return m, nil
}

func (m Model) handleChangeViewMsg(msg scripts.ChangeViewMsg) (tea.Model, tea.Cmd) {
	m.CurrentView = msg.View
	m.UIList = msg.List
	return m, nil
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
	m.Library = utils.Library {
		Songs: msg.Songs,
		Albums: msg.Albums,
		Playlists: msg.Playlists,
	}
	utils.Log(fmt.Sprintf("handleLibraryMsg: songs=%d albums=%d playlists=%d", 
		len(msg.Songs), len(msg.Albums), len(msg.Playlists)))
	return m, nil
}

func containsUIList(view utils.View) bool {
	return true
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// do stuff
	return m, nil
}


