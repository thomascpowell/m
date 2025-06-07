package lists

/**
* bubbles/list item definitions.
*/


// Represents a m.UIList item.
// Used in Albums, Playlists and Detail views.
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

// Represents a m.UIList item.
// Used in the Base view.
type BaseListItem struct {
	Name string
	Action string
}
func (a BaseListItem) Title() string {
	return a.Name
}
func (a BaseListItem) Description() string { return a.Action }
func (a BaseListItem) FilterValue() string {
	return a.Name
}


