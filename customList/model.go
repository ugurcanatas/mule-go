package customList

import (
	"mule-go/constants"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	Name       string
	Identifier string
}

type Item item

// FilterValue implements list.Item.
func (i Item) FilterValue() string { return i.Identifier }
func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Identifier }

func CreateNewList(items []list.Item, title string) list.Model {
	const defaultWidth = 100
	// TODO: Replace with this after learning about custom itemDelegate implementation
	// l := list.New(items, itemDelegate{}, defaultWidth, constants.DEFAULT_LIST_HEIGHT)

	newDefaultDelegate := list.NewDefaultDelegate()
	c := lipgloss.Color("#6f03fc")
	newDefaultDelegate.Styles.SelectedTitle = newDefaultDelegate.Styles.SelectedTitle.Foreground(c).BorderLeftForeground(c)
	newDefaultDelegate.ShowDescription = false

	list := list.New(items, newDefaultDelegate, defaultWidth, constants.DEFAULT_LIST_HEIGHT)
	list.Title = title
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)
	list.ShowTitle()
	list.Styles.Title = constants.TitleStyle
	list.Styles.PaginationStyle = constants.PaginationStyle
	list.Styles.HelpStyle = constants.HelpStyle

	return list
}
