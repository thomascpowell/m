package views

import (
	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/term"
	"m/lists"
	"m/styles"
	"m/utils"
	"os"
)

func NewAlbumDetailList(songs []utils.Song, name string, artist string) list.Model {
	items := make([]list.Item, len(songs)+1)
	items[0] = lists.ListItem{
		Name: "Play All",
		Desc: name,
		Id:   "PLAY_ALL",
	}
	for i, song := range songs {
		items[i+1] = lists.ListItem{
			Name: song.Title,
			Desc: song.Artist + " • " + song.Duration,
			Id:   song.SongId,
		}
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NDL: " + err.Error())
	}
	l := list.New(items, lists.ListDelegate(), width, height)
	l.Title = name + " • " + artist
	title := l.Styles.Title
	title = title.Foreground(styles.Dark).Background(styles.Light)
	l.Styles.Title = title
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	return l
}
