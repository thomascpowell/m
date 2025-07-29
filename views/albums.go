package views

import (
	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/term"
	"m/lists"
	"m/styles"
	"m/utils"
	"os"
)

func NewAlbumList(sources []utils.Source) list.Model {
	items := make([]list.Item, len(sources))
	for i, source := range sources {
		items[i] = lists.ListItem{
			Name: source.Title,
			Desc: source.Artist,
		}
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NSL: " + err.Error())
	}
	l := list.New(items, lists.ListDelegate(), width, height)
	l.Title = "Albums:"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	title := l.Styles.Title
	title = title.Foreground(styles.Dark).Background(styles.Light)
	l.Styles.Title = title
	return l
}
