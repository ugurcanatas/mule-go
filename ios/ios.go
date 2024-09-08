package ios

import (
	"mule-go/constants"
	"mule-go/customList"
	"mule-go/sharedState"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func (m IOSModel) Init() tea.Cmd {
	return nil
}

func (m IOSModel) View() string {
	if sharedState.QuitProgram {
		return constants.QuitTextStyle.Render("Bye 👋🏽")
	}
	return "\n" + m.uiModel.list.View()
}

func (m IOSModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

			if m.stateModel.currentStep == StepRuntime {
				// update the list to commands
				m.stateModel.SetRuntimeSelection(i)
				m.stateModel.SetCurrentStep(StepDevice)

				m.createDevicesListByRuntimeIdentifier(i.Identifier)

				m.uiModel.list, cmd = m.uiModel.list.Update(msg)
				return m, cmd
			}
			if m.stateModel.currentStep == StepDevice {
				// run the commands
				m.stateModel.SetCurrentStep(StepCommand)
				m.stateModel.SetDeviceSelection(i)

				m.createIOSCommands()

				m.uiModel.list, cmd = m.uiModel.list.Update(msg)
				return m, cmd
			}
			if m.stateModel.currentStep == StepCommand {
				// run the commands
				m.stateModel.SetCommandSelection(i)

				d := IOSDevicesByRuntimeIdentifier(m.stateModel.runtimeSelection.Identifier)

				deviceIndex := slices.IndexFunc(d, func(d IOSDevice) bool {
					return d.Udid == m.stateModel.deviceSelection.Identifier
				})

				device := d[deviceIndex]

				RunAppleScript(m.stateModel.cmdSelection.Name, device.Udid, device.State)

				m.uiModel.list.Update(msg)
				return m, tea.Quit
			}

		}

	}

	m.uiModel.list, cmd = m.uiModel.list.Update(msg)
	return m, cmd
}

// Create the initial os model
func (m *IOSModel) InitialModel() IOSModel {
	m.createRuntimesList()

	return IOSModel{
		stateModel: StateModel{
			currentStep: StepRuntime,
		},
		uiModel: UIModel{
			list: m.uiModel.list,
		},
	}
}
