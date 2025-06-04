package views

import(
	"m/utils"
	"m/styles"
	"m/lists"
	"github.com/charmbracelet/lipgloss"
	"os"
	"golang.org/x/term"
	"github.com/charmbracelet/bubbles/list"
)

func NewMenuList() list.Model {
	utils.Log("NML")
	items := []list.Item{
		lists.BaseListItem{
			Name:   "Play",
			Action: "NO_ACTION",
		},
		lists.BaseListItem{
			Name:   "Pause",
			Action: "NO_ACTION",
		},
		lists.BaseListItem{
			Name:   "Example",
			Action: "NO_ACTION",
		},
	}
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		utils.Log("NSL: " + err.Error())
	}
	l := list.New(items, lists.MenuDelegate(), width, height-4)
	l.Title = "Options:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	title_style := lipgloss.NewStyle().MarginLeft(0).Foreground(styles.Light)
	l.Styles.Title = title_style
	return l
}


