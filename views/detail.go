package views

import(
	"m/utils"
	"m/styles"
	"github.com/charmbracelet/bubbles/list"
	"os"
	"golang.org/x/term"
)

/**
* Detail view.
* Shows the content of a playlist or album.
*/


func NewDetailList(songs []utils.Song, name string, artist string, source utils.SourceType) list.Model {
	// define the "play all" command
	// see views/update for how this is used
	source_type_string := "PLAYLIST"
	if source == utils.Album {
		source_type_string = "ALBUM"
	}
	items := make([]list.Item, len(songs)+1)
	items[0] = utils.ListItem {
		Name: name,
		Desc: "Play All",
		Id:   "PLAY_" + source_type_string,
	}
	
	// add songs to the list
	for i, song := range songs {
		items[i+1] = utils.ListItem {
			Name: 	song.Title,
			Desc:   song.Artist + " • " + song.Duration,
			Id:			song.SongId, // Id used for playing with PID
		}
	}

	// define styles
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NDL: " + err.Error())
	}
	l := list.New(items,  styles.ListDelegate(), width, height)
	l.Title = name + " • " + artist
	title := l.Styles.Title
	title = title.Foreground(styles.Dark).Background(styles.Light)
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
