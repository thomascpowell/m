package app

import(
	"fmt"
	"m/scripts"
	"m/utils"
	"github.com/charmbracelet/lipgloss"
)

/**
* tea View() function.
*/


func (m Model) View() string {
	switch m.CurrentView {
	case utils.Albums, utils.Playlists, utils.PlaylistDetail, utils.AlbumDetail:
		return ShowView(m)
	case utils.Menu:
		return ShowMenuView(m)
	default:
		return ""
	}
}

func ShowView(m Model) string {
	return m.UIList.View()
}

func ShowMenuView(m Model) string {
	text := fmt.Sprintf(
			"\n%s — %s\n(%s)\n\n",
			m.CurrentSong.Title,
			m.CurrentSong.Artist,
			scripts.IsPlayingToString(m.IsPlaying))
	if m.CurrentSong.Title == "" {
		text = "\nNot Playing\n\n\n"
	}
	now_playing := lipgloss.NewStyle().PaddingLeft(2).Render(text)
	return lipgloss.JoinVertical(lipgloss.Left, now_playing, m.UIList.View())
}

