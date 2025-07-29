package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"m/scripts"
	"m/utils"
	"m/views"
)

/**
* Initializes the app.
 */

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		scripts.GetLibraryCmd(),
		scripts.RefreshStateCmd(),
		scripts.RefreshLibraryCmd(),
		scripts.ChangeViewCmd(utils.Menu, views.NewMenuList()),
		TickCmd(),
	)
}
