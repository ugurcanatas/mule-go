package platform

import (
	"mule-go/ios"

	"github.com/charmbracelet/bubbles/list"
)

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

func (u *OSModel) SetActionChoice(choice ios.RuntimesModel) {
	u.actionChoice = choice
}

func (u *OSModel) SetIsQuit(isQuit bool) {
	u.isQuit = isQuit
}

func (u *OSModel) AppendIOSCommands(command ios.RuntimesModel) {
	u.iosCommands = append(u.iosCommands, command)
}

func (u *OSModel) IncrementCurrentStepIndex() {
	u.currentStepIndex += 1
}

// Setters
