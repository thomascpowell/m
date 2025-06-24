package app

import(
	"m/scripts"
	"m/views"
	"m/utils"
	tea "github.com/charmbracelet/bubbletea"
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
