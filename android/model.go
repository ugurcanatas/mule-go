package android

import (
	"mule-go/sharedState"

	"github.com/charmbracelet/bubbles/list"
)

type AndroidModel struct {
	uiModel    UIModel
	stateModel StateModel
}

type UIModel struct {
	list list.Model
}

type Step uint

const (
	StepDevice Step = iota
	StepCommand
)

type StateModel struct {
	currentStep     Step
	deviceSelection string
	cmdSelection    string
}

func (model *AndroidModel) QuitProgram() {
	sharedState.QuitProgram = true
}

func (model *UIModel) SetList(list list.Model) {
	model.list = list
}

func (model *StateModel) SetCurrentStep(step Step) {
	model.currentStep = step
}

func (model *StateModel) SetDeviceSelection(deviceSelection string) {
	model.deviceSelection = deviceSelection
}

func (model *StateModel) SetCommandSelection(cmdSelection string) {
	model.cmdSelection = cmdSelection
}
