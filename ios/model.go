package ios

import (
	"mule-go/customList"
	"mule-go/sharedState"

	"github.com/charmbracelet/bubbles/list"
)

// Credit goes to https://mholt.github.io/json-to-go/
type RuntimesObject struct {
	Runtimes []struct {
		BundlePath           string `json:"bundlePath"`
		Buildversion         string `json:"buildversion"`
		Platform             string `json:"platform"`
		RuntimeRoot          string `json:"runtimeRoot"`
		Identifier           string `json:"identifier"`
		Version              string `json:"version"`
		IsInternal           bool   `json:"isInternal"`
		IsAvailable          bool   `json:"isAvailable"`
		Name                 string `json:"name"`
		SupportedDeviceTypes []struct {
			BundlePath    string `json:"bundlePath"`
			Name          string `json:"name"`
			Identifier    string `json:"identifier"`
			ProductFamily string `json:"productFamily"`
		} `json:"supportedDeviceTypes"`
	} `json:"runtimes"`
}

type Devices struct {
	Devices map[string][]IOSDevice `json:"devices"`
}

// Credit goes to https://mholt.github.io/json-to-go/
type IOSDevice struct {
	DataPath             string `json:"dataPath"`
	DataPathSize         int    `json:"dataPathSize"`
	LogPath              string `json:"logPath"`
	Udid                 string `json:"udid"`
	IsAvailable          bool   `json:"isAvailable"`
	LogPathSize          int    `json:"logPathSize,omitempty"`
	DeviceTypeIdentifier string `json:"deviceTypeIdentifier"`
	State                string `json:"state"`
	Name                 string `json:"name"`
}

type IOSModel struct {
	uiModel    UIModel
	stateModel StateModel
}

type UIModel struct {
	list list.Model
}

type Step uint

const (
	StepRuntime Step = iota
	StepDevice
	StepCommand
)

type StateModel struct {
	currentStep    Step
	runtimesObject RuntimesObject
	runtimes       []customList.Record

	deviceSelection  customList.Record
	runtimeSelection customList.Record
	cmdSelection     customList.Record
}

func (model *IOSModel) QuitProgram() {
	sharedState.QuitProgram = true
}

func (model *UIModel) SetList(list list.Model) {
	model.list = list
}

func (model *StateModel) SetDeviceSelection(selection customList.Record) {
	model.deviceSelection = selection
}

func (model *StateModel) SetRuntimeSelection(selection customList.Record) {
	model.runtimeSelection = selection
}

func (model *StateModel) SetCommandSelection(selection customList.Record) {
	model.cmdSelection = selection
}

func (model *StateModel) SetCurrentStep(step Step) {
	model.currentStep = step
}

func (model *StateModel) SetRuntimes(runtimes []customList.Record) {
	model.runtimes = runtimes
}

func (model *StateModel) SetRuntimesObject(runtimesObject RuntimesObject) {
	model.runtimesObject = runtimesObject
}
