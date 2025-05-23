package views

import (
	"m/utils"
	"github.com/charmbracelet/bubbles/list"
)

/**
* Sources view.
* Used for AlbumsView and PlaylistsView.
*/

func NewSourceList(sources []utils.Source, name string) list.Model {
	items := make([]list.Item, len(sources))
	for i, source := range sources {
			items[i] = ListItem {
					Name: 	source.Title,
					Desc:   source.Artist,
			}
	}
	const width = 50
	const height = 30
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = name
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)
	return l
}


// Returns the UIList view.
// m.CurrentList stores data used here.
func ShowSourcesView(m Model) string {
	return m.UIList.View()
}
