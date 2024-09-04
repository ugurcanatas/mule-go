package platform

import (
	"mule-go/ios"

	"github.com/charmbracelet/bubbles/list"
)

type OSModel struct {
	list   list.Model
	isQuit bool

	choice        ios.RuntimesModel
	runtimeChoice ios.RuntimesModel
	deviceChoice  ios.RuntimesModel
	actionChoice  ios.RuntimesModel

	// IOS
	xcrunResult ios.XCRunDevices
	runtimes    []ios.RuntimesModel
	iosCommands []ios.RuntimesModel

	// TODO Current implementation to track list progress is scuffed. Replace it with lists or rings
	currentStepIndex uint
}
