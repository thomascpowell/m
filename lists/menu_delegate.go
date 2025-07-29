package lists

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"m/styles"
	"strings"
)

/*
* Defines styling for a list item.
* Used in Menu view.
 */

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(2).Foreground(styles.Dim)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(0).Foreground(styles.Light)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(BaseListItem)
	if !ok {
		return
	}
	str := fmt.Sprintf("%s", i.Name)
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}
	fmt.Fprint(w, fn(str))
}

func MenuDelegate() itemDelegate {
	return itemDelegate{}
}
