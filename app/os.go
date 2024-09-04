package app

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"mule-go/app/ios"
	"mule-go/constants"
	"mule-go/customList"
)

const STEP_IOS_1 = "RUNTIME"
const STEP_IOS_2 = "DEVICE"

type OSModel struct {
	list           list.Model
	choice         ios.RuntimesModel
	runtimeChoice  ios.RuntimesModel
	deviceChoice   ios.RuntimesModel
	isQuit         bool
	currentIOSStep string

	// IOS
	xcrunResult   ios.XCRunDevices
	runtimes      []ios.RuntimesModel
	devicesMapped ios.DevicesMap
	iosCommands   []ios.RuntimesModel

	currentStepIndex uint
}

// Setters

func (u *OSModel) SetRuntimes(runtimes []ios.RuntimesModel) {
	u.runtimes = runtimes
}

func (u *OSModel) SetXCRunResult(xcrunResult ios.XCRunDevices) {
	u.xcrunResult = xcrunResult
}

func (u *OSModel) SetList(list list.Model) {
	u.list = list
}

func (u *OSModel) SetChoice(choice ios.RuntimesModel) {
	u.choice = choice
}

func (u *OSModel) SetRuntimeChoice(choice ios.RuntimesModel) {
	u.runtimeChoice = choice
}

func (u *OSModel) SetDeviceChoice(choice ios.RuntimesModel) {
	u.deviceChoice = choice
}

func (u *OSModel) SetIsQuit(isQuit bool) {
	u.isQuit = isQuit
}

func (u *OSModel) SetDevicesMapped(devicesMapped ios.DevicesMap) {
	u.devicesMapped = devicesMapped
}

func (u *OSModel) SetCurrentIOSStep(step string) {
	u.currentIOSStep = step
}

func (u *OSModel) AppendIOSCommands(command ios.RuntimesModel) {
	u.iosCommands = append(u.iosCommands, command)
}

func (u *OSModel) IncrementCurrentStepIndex() {
	u.currentStepIndex += 1
}

// Setters

func (u *OSModel) FilterDevicesByRuntimeIdentifier(runtimeIdentifier string) []ios.RuntimesModel {
	return u.devicesMapped[runtimeIdentifier]
}

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
					m.SetChoice(ios.RuntimesModel{
						Name:       i.Name,
						Identifier: i.Identifier,
					})
				}
				if m.currentStepIndex == 1 {
					m.SetRuntimeChoice(ios.RuntimesModel{
						Name:       i.Name,
						Identifier: i.Identifier,
					})
				}
				if m.currentStepIndex == 2 {
					m.SetDeviceChoice(ios.RuntimesModel{
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
					// RUN THE COMMAND
				}
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (model *OSModel) createCommandsList() {
	commandsListItems := []list.Item{}

	for index := range model.iosCommands {
		commandsListItems = append(commandsListItems, customList.Item{
			Name:       model.iosCommands[index].Name,
			Identifier: model.iosCommands[index].Name,
		})
	}

	commandsListTitle := fmt.Sprintf("Select a command for device %s ", model.deviceChoice.Name)

	l := customList.CreateNewList(commandsListItems, commandsListTitle)
	model.SetList(l)
}

func (model *OSModel) createDevicesListByRuntimeIdentifier(identifier string) {
	devices := model.FilterDevicesByRuntimeIdentifier(identifier)

	deviceListItems := []list.Item{}

	for x := range devices {
		deviceListItems = append(deviceListItems, customList.Item{
			Name:       devices[x].Name,
			Identifier: devices[x].Identifier,
		})
	}

	runtimeIndex := slices.IndexFunc(model.runtimes, func(c ios.RuntimesModel) bool { return c.Identifier == identifier })

	var deviceListTitle string = constants.IOS_DEVICES_TITLE

	if runtimeIndex > -1 {
		deviceListTitle = strings.Join([]string{constants.IOS_DEVICES_TITLE, "(" + model.runtimes[runtimeIndex].Name + ")"}, " - ")
	}

	l := customList.CreateNewList(deviceListItems, deviceListTitle)
	model.SetList(l)
}

func (model *OSModel) createRuntimesList() {
	tea.LogToFile("debug.log", "debug")

	// get runtimes with xcrun command and set inside the struct
	model.SetXCRunResult(ios.GetIOSRuntimes())

	runtimesByNameAndIdentifier := []ios.RuntimesModel{}
	devicesMapByNameAndIdentifier := make(ios.DevicesMap)

	// View slices
	runtimesListItems := []list.Item{}

	for index := range model.xcrunResult.Runtimes {
		runtimes := model.xcrunResult.Runtimes[index]
		runtimesByNameAndIdentifier = append(runtimesByNameAndIdentifier, ios.RuntimesModel{
			Name:       runtimes.Name,
			Identifier: runtimes.Identifier,
		})
		itemA := customList.Item{
			Name:       runtimes.Name,
			Identifier: runtimes.Identifier,
		}
		runtimesListItems = append(runtimesListItems, itemA)
		for index2 := range runtimes.SupportedDeviceTypes {
			deviceName := runtimes.SupportedDeviceTypes[index2].Name
			deviceIdentifier := runtimes.SupportedDeviceTypes[index2].Identifier

			// Append RuntimesModel to the map for a specific key
			devicesMapByNameAndIdentifier[runtimes.Identifier] = append(devicesMapByNameAndIdentifier[runtimes.Identifier], ios.RuntimesModel{Name: deviceName, Identifier: deviceIdentifier})
		}
	}

	model.SetRuntimes(runtimesByNameAndIdentifier)
	model.SetDevicesMapped(devicesMapByNameAndIdentifier)

	l := customList.CreateNewList(runtimesListItems, constants.IOS_RUNTIME_TITLE)
	model.SetList(l)
}

// Create the initial os model
func InitialModel() OSModel {
	items := []list.Item{
		customList.Item{Name: constants.IOS, Identifier: constants.IOS},
		customList.Item{Name: constants.ANDROID, Identifier: constants.ANDROID},
	}

	list := customList.CreateNewList(items, constants.OS_TITLE)

	// set default ios commands
	commandsSet := []string{"Boot", "Erase", "Send Link", "Shutdown"}
	commandsSlice := []ios.RuntimesModel{}
	for i := 0; i < len(commandsSet); i++ {
		commandsSlice = append(commandsSlice, ios.RuntimesModel{
			Name:       commandsSet[i],
			Identifier: commandsSet[i],
		})
	}

	return OSModel{
		list:           list,
		currentIOSStep: STEP_IOS_1,
		iosCommands:    commandsSlice,
	}
}