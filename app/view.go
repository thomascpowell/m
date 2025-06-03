package app

import(
	"fmt"
	"m/scripts"
	"github.com/charmbracelet/lipgloss"
)

/**
* tea View() function.
*/


func (m Model) View() string {
	if !m.Loaded {
		return ""
	}
	switch m.CurrentView {
	case BaseView:
		return ShowBaseView(m)
	case AlbumsView, PlaylistsView:
		return ShowSourcesView(m)
	case SourceDetailView:
		return ShowDetailView(m)
	default:
		return ""
	}
}

func ShowSourcesView(m Model) string {
	return m.UIList.View()
}

func ShowBaseView(m Model) string {
	text := fmt.Sprintf(
			"%s — %s\n(%s)\n\n",
			m.CurrentSong.Title,
			m.CurrentSong.Artist,
			scripts.IsPlayingToString(m.IsPlaying))
	now_playing := lipgloss.NewStyle().Render(text)
	return lipgloss.JoinVertical(lipgloss.Left, now_playing, m.UIList.View())
}

func ShowDetailView(m Model) string {
	return m.UIList.View()
}
