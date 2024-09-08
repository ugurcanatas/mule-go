package ios

import (
	"fmt"
	"mule-go/constants"
	"mule-go/customList"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

func (model *IOSModel) createRuntimesList() {
	xcRunResult := IOSRuntimes()
	model.stateModel.SetRuntimesObject(xcRunResult)

	runtimesByNameAndIdentifier := []customList.Record{}

	// View slices
	runtimesListItems := []list.Item{}

	for index := range xcRunResult.Runtimes {
		runtimes := xcRunResult.Runtimes[index]
		runtimesByNameAndIdentifier = append(runtimesByNameAndIdentifier, customList.Record{
			Name:       runtimes.Name,
			Identifier: runtimes.Identifier,
		})
		i := customList.Record{
			Name:       runtimes.Name,
			Identifier: runtimes.Identifier,
		}
		runtimesListItems = append(runtimesListItems, i)
	}

	model.stateModel.SetRuntimes(runtimesByNameAndIdentifier)

	l := customList.CreateNewList(runtimesListItems, constants.IOSRuntimeTitle)
	model.uiModel.SetList(l)
}

func (model *IOSModel) createIOSCommands() {
	commandsListItems := []list.Item{}

	for index := range constants.DefaultIOSCommands {
		commandsListItems = append(commandsListItems, customList.Record{
			Name:       constants.DefaultIOSCommands[index],
			Identifier: constants.DefaultIOSCommands[index],
		})
	}

	// TODO move string to a constant
	commandsListTitle := fmt.Sprintf("Select a command for device %s ", model.stateModel.deviceSelection.Name)

	l := customList.CreateNewList(commandsListItems, commandsListTitle)
	model.uiModel.SetList(l)
}

func (model *IOSModel) createDevicesListByRuntimeIdentifier(identifier string) {
	devices := IOSDevicesByRuntimeIdentifier(identifier)

	deviceListItems := []list.Item{}

	for x := range devices {
		deviceListItems = append(deviceListItems, customList.Record{
			Name:       devices[x].Name,
			Identifier: devices[x].Udid,
		})
	}

	runtimeIndex := slices.IndexFunc(model.stateModel.runtimes, func(c customList.Record) bool { return c.Identifier == identifier })

	var deviceListTitle string = constants.IOSDevicesTitle

	if runtimeIndex > -1 {
		deviceListTitle = strings.Join([]string{constants.IOSDevicesTitle, "(" + model.stateModel.runtimes[runtimeIndex].Name + ")"}, " - ")
	}

	l := customList.CreateNewList(deviceListItems, deviceListTitle)
	model.uiModel.SetList(l)
}
