package platform

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"mule-go/android"
	"mule-go/constants"
	"mule-go/customList"
	"mule-go/ios"
	"mule-go/sharedState"
)

func (m OSModel) Init() tea.Cmd {
	return nil
}

func (m OSModel) View() string {
	if sharedState.QuitProgram {
		return constants.QuitTextStyle.Render("Bye 👋🏽")
	}
	return "\n" + m.list.View()
}

func (m OSModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.QuitProgram()
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(customList.Record)

			if !ok {
				return m, tea.Quit
			}

			if i.Identifier == constants.IOS {
				iosModel := ios.IOSModel{}
				iosModel.InitialModel()
				return m.ChangeViews(&iosModel)
			}
			if i.Identifier == constants.ANDROID {
				androidModel := android.AndroidModel{}
				androidModel.InitialModel()
				return m.ChangeViews(&androidModel)
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// Create the initial os model
func InitialModel() OSModel {
	items := []list.Item{
		customList.Record{Name: constants.IOS, Identifier: constants.IOS},
		customList.Record{Name: constants.ANDROID, Identifier: constants.ANDROID},
	}

	list := customList.CreateNewList(items, constants.OsTitle)

	return OSModel{
		list: list,
	}
}

func (m OSModel) ChangeViews(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init()
}
