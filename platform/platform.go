package platform

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"mule-go/constants"
	"mule-go/customList"
	"mule-go/ios"
)

func (m OSModel) Init() tea.Cmd {
	return nil
}

func (m OSModel) View() string {
	if m.isQuit {
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
			m.SetIsQuit(true)
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(customList.Item)

			if ok {
				if m.currentStepIndex == 0 {
					m.SetChoice(Record{
						Name:       i.Name,
						Identifier: i.Identifier,
					})
				}
				if m.currentStepIndex == 1 {
					m.SetRuntimeChoice(Record{
						Name:       i.Name,
						Identifier: i.Identifier,
					})
				}
				if m.currentStepIndex == 2 {
					m.SetDeviceChoice(Record{
						Name:       i.Name,
						Identifier: i.Identifier,
					})
				}
				if m.currentStepIndex == 3 {
					m.SetActionChoice(Record{
						Name:       i.Name,
						Identifier: i.Identifier,
					})
				}

				m.IncrementCurrentStepIndex()
			}

			if m.choice.Identifier == constants.IOS { // IOS is selected Update the list
				if m.currentStepIndex == 1 {
					m.createRuntimesList()
				}
				if m.currentStepIndex == 2 {
					m.createDevicesListByRuntimeIdentifier(m.runtimeChoice.Identifier)
				}
				if m.currentStepIndex == 3 {
					m.createCommandsList()
				}
				if m.currentStepIndex == 4 {
					d := ios.CurrentDevices.DeviceByDeviceUDID(m.deviceChoice.Identifier)
					ios.RunAppleScript(m.actionChoice.Name, d.Udid, d.State)
					m.SetIsQuit(true)
					return m, tea.Quit
				}
			}
			if m.choice.Identifier == constants.ANDROID {
				if m.currentStepIndex == 1 {
					m.createAvdDevicesList()
				}
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// Create the initial os model
func InitialModel() OSModel {
	items := []list.Item{
		customList.Item{Name: constants.IOS, Identifier: constants.IOS},
		customList.Item{Name: constants.ANDROID, Identifier: constants.ANDROID},
	}

	list := customList.CreateNewList(items, constants.OsTitle)

	// set default ios commands
	commandsSlice := []Record{}
	for key := range constants.DefaultIOSCommands {
		commandsSlice = append(commandsSlice, Record{
			Name:       constants.DefaultIOSCommands[key],
			Identifier: constants.DefaultIOSCommands[key],
		})
	}

	return OSModel{
		list:        list,
		iosCommands: commandsSlice,
	}
}
