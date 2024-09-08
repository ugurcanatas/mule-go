package android

import (
	"mule-go/constants"
	"mule-go/customList"

	"github.com/charmbracelet/bubbles/list"
)

func (m *AndroidModel) createAvdDevicesList() {
	devices := AVDEmulators()

	deviceRecords := []list.Item{}
	for d := range devices {
		deviceRecords = append(deviceRecords, customList.Record{
			Name: devices[d],
			// there is no unique identifier returned from the command
			Identifier: devices[d],
		})
	}

	l := customList.CreateNewList(deviceRecords, constants.IOSRuntimeTitle)
	m.uiModel.SetList(l)
}

func (model *AndroidModel) createAndroidCommands() {
	deviceRecords := []list.Item{}
	for d := range constants.DefaultAndroidCommands {
		deviceRecords = append(deviceRecords, customList.Record{
			Name: constants.DefaultAndroidCommands[d],
			// there is no unique identifier returned from the command
			Identifier: constants.DefaultAndroidCommands[d],
		})
	}

	// TODO move string to a constant
	// commandsListTitle := fmt.Sprintf("Select a command for device %s ", model.deviceChoice.Name)

	l := customList.CreateNewList(deviceRecords, constants.IOSRuntimeTitle)
	model.uiModel.SetList(l)

}
