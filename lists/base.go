package lists

import(
	"m/utils"
	"m/styles"
	// "m/scripts"
	// "fmt"
	"github.com/charmbracelet/lipgloss"
	"os"
	"golang.org/x/term"
	"github.com/charmbracelet/bubbles/list"
)

// Represents a m.UIList item.
// Used in the Base view.
type BaseListItem struct {
	Name string
	Action string
}
func (a BaseListItem) Title() string {
	return a.Name
}
func (i BaseListItem) Description() string { return "" }
func (a BaseListItem) FilterValue() string {
	return a.Name
}

func NewBaseList() list.Model {
	items := []list.Item{
		BaseListItem{
			Name:   "Play",
			Action: "NO_ACTION",
		},
		BaseListItem{
			Name:   "Pause",
			Action: "NO_ACTION",
		},
		BaseListItem{
			Name:   "Example",
			Action: "NO_ACTION",
		},
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NSL: " + err.Error())
	}
	l := list.New(items, BaseDelegate(), width, height-4)
	l.Title = "Options:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	title_style := lipgloss.NewStyle().MarginLeft(0).Foreground(styles.Light)
	l.Styles.Title = title_style
	return l
}
