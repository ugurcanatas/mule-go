package android

import (
	"mule-go/constants"
	"mule-go/customList"
	"mule-go/sharedState"

	tea "github.com/charmbracelet/bubbletea"
)

func (m AndroidModel) Init() tea.Cmd {
	return nil
}

func (m AndroidModel) View() string {
	if sharedState.QuitProgram {
		return constants.QuitTextStyle.Render("Bye 👋🏽")
	}
	return "\n" + m.uiModel.list.View()
}

func (m AndroidModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.uiModel.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.QuitProgram()
			return m, tea.Quit

		case "enter":
			i, ok := m.uiModel.list.SelectedItem().(customList.Record)

			if !ok {
				return m, tea.Quit
			}

			if m.stateModel.currentStep == StepDevice {
				// update the list to commands
				m.stateModel.SetDeviceSelection(i.Name)
				m.stateModel.SetCurrentStep(StepCommand)

				m.createAndroidCommands()

				m.uiModel.list, cmd = m.uiModel.list.Update(msg)
				return m, cmd
			}
			if m.stateModel.currentStep == StepCommand {
				m.stateModel.SetCommandSelection(i.Name)

				// TODO, Handle based on commands
				RunAndroidEmulator(m.stateModel.deviceSelection)
				m.uiModel.list, cmd = m.uiModel.list.Update(msg)
				return m, tea.Quit
			}

		}

	}

	m.uiModel.list, cmd = m.uiModel.list.Update(msg)
	return m, cmd
}

// Create the initial os model
func (m *AndroidModel) InitialModel() AndroidModel {
	m.createAvdDevicesList()

	return AndroidModel{
		stateModel: StateModel{
			currentStep: StepDevice,
		},
		uiModel: UIModel{
			list: m.uiModel.list,
		},
	}
}
