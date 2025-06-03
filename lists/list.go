package lists

import(
	"m/utils"
	"m/styles"
	"os"
	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/term"
)

/**
* bubbles/list configuration.
*/


// Represents a m.UIList item.
// Used in Albums, Playlists and Detail views.
type ListItem struct {
	Name string
	Desc string
	Id	 string
}
func (a ListItem) Title() string {
	return a.Name
}
func (a ListItem) Description() string {
	return a.Desc
}
func (a ListItem) FilterValue() string {
	return a.Name
}

func NewDetailList(songs []utils.Song, name string, artist string, source utils.SourceType) list.Model {
	// define the "play all" command
	source_type_string := "PLAYLIST"
	if source == utils.Album {
		source_type_string = "ALBUM"
	}

	items := make([]list.Item, len(songs)+1)
	items[0] = ListItem {
		Name: name,
		Desc: "Play All",
		Id:   "PLAY_" + source_type_string,
	}
	// add songs to the list
	for i, song := range songs {
		items[i+1] = ListItem {
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
	l := list.New(items,  ListDelegate(), width, height)
	l.Title = name + " • " + artist
	title := l.Styles.Title
	title = title.Foreground(styles.Dark).Background(styles.Light)
	l.Styles.Title = title
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
  return l
}

func NewSourceList(sources []utils.Source, name string) list.Model {
	items := make([]list.Item, len(sources))
	for i, source := range sources {
		items[i] = ListItem {
			Name: 	source.Title,
			Desc:   source.Artist,
		}
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NSL: " + err.Error())
	}
	l := list.New(items, ListDelegate(), width, height)
	l.Title = name
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	title := l.Styles.Title
	title = title.Foreground(styles.Dark).Background(styles.Light)
	l.Styles.Title = title
	return l
}
