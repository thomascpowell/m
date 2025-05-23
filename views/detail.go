package views

import(
	"m/utils"
	"github.com/charmbracelet/bubbles/list"
)

/**
* Detail view.
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
func ShowDetailView(m Model) string {
	return m.UIList.View()
}
