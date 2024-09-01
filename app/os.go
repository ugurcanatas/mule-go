package app

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item string

type itemDelegate struct{}

func (i item) FilterValue() string { return "" }

// Height implements list.ItemDelegate.
func (d itemDelegate) Height() int {
	return 1
}

// Render implements list.ItemDelegate.
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

// Spacing implements list.ItemDelegate.
func (d itemDelegate) Spacing() int {
	return 1
}

// Update implements list.ItemDelegate.
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

const IOS string = "iOS"
const ANDROID string = "Android"

type model struct {
	list   list.Model
	choice string
	isQuit bool
}

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)


func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Selected ->", m.choice))
	}
	if m.isQuit {
		return quitTextStyle.Render("Bye 👋🏽")
	}
	return "\n" + m.list.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.isQuit = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// Setup initial list model
func InitialModel() model {
	items := []list.Item{
		item(IOS),
		item(ANDROID),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select OS:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{
		list: l,
	}
}
