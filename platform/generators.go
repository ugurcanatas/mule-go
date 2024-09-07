package platform

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"

	"mule-go/android"
	"mule-go/constants"
	"mule-go/customList"
	"mule-go/ios"
)

func (model *OSModel) createIOSCommands() list.Model {
	commandsListItems := []list.Item{}

	for index := range constants.DefaultIOSCommands {
		commandsListItems = append(commandsListItems, customList.Item{
			Name:       constants.DefaultIOSCommands[index],
			Identifier: constants.DefaultIOSCommands[index],
		})
	}

	// TODO move string to a constant
	commandsListTitle := fmt.Sprintf("Select a command for device %s ", model.deviceChoice.Name)

	return customList.CreateNewList(commandsListItems, commandsListTitle)
}

func (model *OSModel) createAndroidCommands() list.Model {
	deviceRecords := []list.Item{}
	for d := range constants.DefaultAndroidCommands {
		deviceRecords = append(deviceRecords, customList.Item{
			Name: constants.DefaultAndroidCommands[d],
			// there is no unique identifier returned from the command
			Identifier: constants.DefaultAndroidCommands[d],
		})
	}

	// TODO move string to a constant
	// commandsListTitle := fmt.Sprintf("Select a command for device %s ", model.deviceChoice.Name)

	return customList.CreateNewList(deviceRecords, constants.IOSRuntimeTitle)

}

func (model *OSModel) createCommandsList() {
	commandsList := list.Model{}

	if model.choice.Name == constants.IOS {
		commandsList = model.createIOSCommands()
	} else if model.choice.Name == constants.ANDROID {
		commandsList = model.createAndroidCommands()
	}

	model.SetList(commandsList)
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

	runtimeIndex := slices.IndexFunc(model.runtimes, func(c Record) bool { return c.Identifier == identifier })

	var deviceListTitle string = constants.IOSDevicesTitle

	if runtimeIndex > -1 {
		deviceListTitle = strings.Join([]string{constants.IOSDevicesTitle, "(" + model.runtimes[runtimeIndex].Name + ")"}, " - ")
	}

	l := customList.CreateNewList(deviceListItems, deviceListTitle)
	model.SetList(l)
}

func (model *OSModel) createRuntimesList() {

	// get runtimes with xcrun command and set inside the struct
	// model.SetXCRunResult(ios.IOSRuntimes())

	xcRunResult := ios.IOSRuntimes()
	model.SetXCRunResult(xcRunResult)

	runtimesByNameAndIdentifier := []Record{}

	// View slices
	runtimesListItems := []list.Item{}

	for index := range xcRunResult.Runtimes {
		runtimes := xcRunResult.Runtimes[index]
		runtimesByNameAndIdentifier = append(runtimesByNameAndIdentifier, Record{
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

func (model *OSModel) createAvdDevicesList() {
	devices := android.AVDEmulators()

	deviceRecords := []list.Item{}
	for d := range devices {
		deviceRecords = append(deviceRecords, customList.Item{
			Name: devices[d],
			// there is no unique identifier returned from the command
			Identifier: devices[d],
		})
	}

	l := customList.CreateNewList(deviceRecords, constants.IOSRuntimeTitle)
	model.SetList(l)
}
