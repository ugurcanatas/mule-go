package constants

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const IOS string = "iOS"
const ANDROID string = "Android"

const OsTitle string = "Select the OS Type:"
const IOSRuntimeTitle string = "Select a runtime:"
const IOSDevicesTitle string = "Select a device:"

const DefaultListHeight int = 14
const DefaultListWidth int = 100

var DefaultIOSCommands []string = []string{"Boot", "Erase", "Send Link", "Shutdown"}

var (
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
