package lists

import(
	"m/utils"
	"m/styles"
	// "m/scripts"
	// "fmt"
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
			Name:   "Hello",
			Action: "NO_ACTION",
		},
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NSL: " + err.Error())
	}
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Options:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	title := l.Styles.Title
	title = title.Foreground(styles.Dark).Background(styles.Light)
	l.Styles.Title = title
	return l
}
