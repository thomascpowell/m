package views

import(
	"m/utils"
	"m/colors"
	"m/delegates"
	"github.com/charmbracelet/bubbles/list"
	"os"
	"golang.org/x/term"
)

/**
* Detail view.
* Shows the content of a playlist or album.
*/


func NewDetailList(songs []utils.Song, name string, artist string) list.Model {
	items := make([]list.Item, len(songs))
	for i, source := range songs {
			items[i] = utils.ListItem {
					Name: 	source.Title,
					Desc:   source.Artist + " • " + source.Duration,
					Id:			source.SongId, // Id used for playing with PID
			}
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NSL: " + err.Error())
	}
	l := list.New(items,  delegates.ListDelegate(), width, height)
	l.Title = name + " • " + artist


	title := l.Styles.Title
	title = title.Foreground(colors.Dark).Background(colors.Light)
	l.Styles.Title = title
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
