package constants

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const IOS string = "iOS"
const ANDROID string = "Android"

const OS_TITLE string = "Select the OS Type:"
const IOS_RUNTIME_TITLE string = "Select a runtime:"
const IOS_DEVICES_TITLE string = "Select a device:"

const DEFAULT_LIST_HEIGHT int = 14

var (
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
