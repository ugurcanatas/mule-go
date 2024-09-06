package platform

import (
	"github.com/charmbracelet/bubbles/list"
)

// Setters

func (u *OSModel) SetRuntimes(runtimes []Record) {
	u.runtimes = runtimes
}

func (u *OSModel) SetList(list list.Model) {
	u.list = list
}

func (u *OSModel) SetChoice(choice Record) {
	u.choice = choice
}

func (u *OSModel) SetRuntimeChoice(choice Record) {
	u.runtimeChoice = choice
}

func (u *OSModel) SetDeviceChoice(choice Record) {
	u.deviceChoice = choice
}

func (u *OSModel) SetActionChoice(choice Record) {
	u.actionChoice = choice
}

func (u *OSModel) SetIsQuit(isQuit bool) {
	u.isQuit = isQuit
}

func (u *OSModel) AppendIOSCommands(command Record) {
	u.iosCommands = append(u.iosCommands, command)
}

func (u *OSModel) IncrementCurrentStepIndex() {
	u.currentStepIndex += 1
}

// Setters
