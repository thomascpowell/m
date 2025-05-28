package styles

import (
	"github.com/charmbracelet/bubbles/list"
)

/*
* Defines styling for a list item.
*/

func ListDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	
	selectedTitle := d.Styles.SelectedTitle
	selectedDesc := d.Styles.SelectedDesc
	normalTitle := d.Styles.NormalTitle
	normalDesc := d.Styles.NormalDesc

	selectedTitle = selectedTitle.Foreground(Light).BorderForeground(Border)
	selectedDesc = selectedDesc.Foreground(Dim).BorderForeground(Border)
	normalTitle = normalTitle.Foreground(Light)
	normalDesc = normalDesc.Foreground(Dim)

	d.Styles.SelectedTitle = selectedTitle
	d.Styles.SelectedDesc = selectedDesc
	d.Styles.NormalTitle = normalTitle
	d.Styles.NormalDesc = normalDesc

	return d
}

