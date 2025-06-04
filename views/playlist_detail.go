package views


import(
	"m/utils"
	"m/styles"
	"m/lists"
	"os"
	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/term"
)

func NewPlaylistDetailList(songs []utils.Song, name string, artist string) list.Model {
	source_type_string := "PLAYLIST"
	items := make([]list.Item, len(songs)+1)

	items[0] = lists.ListItem {
		Name: name,
		Desc: "Play All",
		Id:   "PLAY_" + source_type_string,
	}

	for i, song := range songs {
		items[i+1] = lists.ListItem {
			Name: 	song.Title,
			Desc:   song.Artist + " • " + song.Duration,
			Id:			song.SongId, 
		}
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NDL: " + err.Error())
	}

	l := list.New(items,  lists.ListDelegate(), width, height)
	l.Title = name + " • " + artist

	title := l.Styles.Title
	title = title.Foreground(styles.Dark).Background(styles.Light)
	l.Styles.Title = title
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)

  return l
}


