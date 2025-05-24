package views

import (
	"m/utils"
	"m/colors"
	"m/delegates"
	"os"
	"golang.org/x/term"
	"github.com/charmbracelet/bubbles/list"
	// "github.com/charmbracelet/lipgloss"
)

/**
* Sources view.
* Used for AlbumsView and PlaylistsView.
*/

func NewSourceList(sources []utils.Source, name string) list.Model {
	items := make([]list.Item, len(sources))
	for i, source := range sources {
			items[i] = utils.ListItem {
					Name: 	source.Title,
					Desc:   source.Artist,
			}
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NSL: " + err.Error())
	}

	l := list.New(items, delegates.ListDelegate(), width, height)
	l.Title = name
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	title := l.Styles.Title
	title = title.Foreground(colors.Dark).Background(colors.Light)
	l.Styles.Title = title

	return l
}


// Returns the UIList view.
// m.CurrentList stores data used here.
func ShowSourcesView(m Model) string {
	return m.UIList.View()
}
