package lists

import (
	"github.com/charmbracelet/bubbles/list"
	"m/styles"
)

/*
* Defines styling for a list item.
* Used in all views but the base.
 */

func ListDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	selectedTitle := d.Styles.SelectedTitle
	selectedDesc := d.Styles.SelectedDesc
	normalTitle := d.Styles.NormalTitle
	normalDesc := d.Styles.NormalDesc

	selectedTitle = selectedTitle.Foreground(styles.Light).BorderForeground(styles.Border)
	selectedDesc = selectedDesc.Foreground(styles.Dim).BorderForeground(styles.Border)
	normalTitle = normalTitle.Foreground(styles.Light)
	normalDesc = normalDesc.Foreground(styles.Dim)

	d.Styles.SelectedTitle = selectedTitle
	d.Styles.SelectedDesc = selectedDesc
	d.Styles.NormalTitle = normalTitle
	d.Styles.NormalDesc = normalDesc

	return d
}
