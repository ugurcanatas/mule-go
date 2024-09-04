package platform

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"

	"mule-go/constants"
	"mule-go/customList"
	"mule-go/ios"
)

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
	devices := ios.IOSDevicesByRuntimeIdentifier(identifier)

	deviceListItems := []list.Item{}

	for x := range devices {
		deviceListItems = append(deviceListItems, customList.Item{
			Name:       devices[x].Name,
			Identifier: devices[x].Udid,
		})
	}

	runtimeIndex := slices.IndexFunc(model.runtimes, func(c ios.RuntimesModel) bool { return c.Identifier == identifier })

	var deviceListTitle string = constants.IOSDevicesTitle

	if runtimeIndex > -1 {
		deviceListTitle = strings.Join([]string{constants.IOSDevicesTitle, "(" + model.runtimes[runtimeIndex].Name + ")"}, " - ")
	}

	l := customList.CreateNewList(deviceListItems, deviceListTitle)
	model.SetList(l)
}

func (model *OSModel) createRuntimesList() {

	// get runtimes with xcrun command and set inside the struct
	model.SetXCRunResult(ios.IOSRuntimes())

	runtimesByNameAndIdentifier := []ios.RuntimesModel{}

	// View slices
	runtimesListItems := []list.Item{}

	for index := range model.xcrunResult.Runtimes {
		runtimes := model.xcrunResult.Runtimes[index]
		runtimesByNameAndIdentifier = append(runtimesByNameAndIdentifier, ios.RuntimesModel{
			Name:       runtimes.Name,
			Identifier: runtimes.Identifier,
		})
		i := customList.Item{
			Name:       runtimes.Name,
			Identifier: runtimes.Identifier,
		}
		runtimesListItems = append(runtimesListItems, i)
	}

	model.SetRuntimes(runtimesByNameAndIdentifier)

	l := customList.CreateNewList(runtimesListItems, constants.IOSRuntimeTitle)
	model.SetList(l)
}
