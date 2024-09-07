package platform

import (
	"mule-go/ios"

	"github.com/charmbracelet/bubbles/list"
)

type OSModel struct {
	list   list.Model
	isQuit bool

	choice        Record
	runtimeChoice Record
	deviceChoice  Record
	actionChoice  Record

	// IOS
	xcrunResult ios.XCRunDevices
	runtimes    []Record

	//

	// TODO Current implementation to track list progress is scuffed. Replace it with lists or rings
	currentStepIndex uint
}

type Record struct {
	Name       string
	Identifier string
}
