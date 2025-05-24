
package delegates

import (
	"github.com/charmbracelet/bubbles/list"
	// "github.com/charmbracelet/lipgloss"
	"m/colors"

)

/*
* Defines styping for a list item.
*/

func ListDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	
	selectedTitle := d.Styles.SelectedTitle
	selectedDesc := d.Styles.SelectedDesc
	normalTitle := d.Styles.NormalTitle
	normalDesc := d.Styles.NormalDesc

	selectedTitle = selectedTitle.Foreground(colors.Light).BorderForeground(colors.Border)
	selectedDesc = selectedDesc.Foreground(colors.Dim).BorderForeground(colors.Border)
	normalTitle = normalTitle.Foreground(colors.Light)
	normalDesc = normalDesc.Foreground(colors.Dim)

	d.Styles.SelectedTitle = selectedTitle
	d.Styles.SelectedDesc = selectedDesc
	d.Styles.NormalTitle = normalTitle
	d.Styles.NormalDesc = normalDesc

	return d
}

