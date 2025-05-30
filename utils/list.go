package utils

/**
* bubbles/list configuration.
*/

// Represents a m.UIList item
// Either a playlist or Album
type ListItem struct {
	Name string
	Desc string
	Id	 string
}
func (a ListItem) Title() string {
	return a.Name
}
func (a ListItem) Description() string {
	return a.Desc
}
func (a ListItem) FilterValue() string {
	return a.Name
}
