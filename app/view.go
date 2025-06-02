package app

/**
* tea View() function.
*/


func (m Model) View() string {
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
	return m.UIList.View()
}

func ShowDetailView(m Model) string {
	return m.UIList.View()
}
