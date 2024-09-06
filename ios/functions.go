package ios

import (
	"encoding/json"
	"fmt"
	"log"
	"mule-go/constants"
	"os/exec"
	"slices"
)

func (u *XCRunDevices) SetXCRunResult(xcrunResult XCRunDevices) {
	*u = xcrunResult
}

func (m *FilteredIOSDeviceList) DeviceByDeviceUDID(udid string) IOSDevice {
	index := slices.IndexFunc(*m, func(d IOSDevice) bool {
		return d.Udid == udid
	})
	return (*m)[index]
}

func NewDevicesSlice() *FilteredIOSDeviceList {
	inital := []IOSDevice{}
	return (*FilteredIOSDeviceList)(&inital)
}

func NewXCRunResultSlice() *XCRunDevices {
	inital := XCRunDevices{}
	return &inital
}

func IOSDevicesByRuntimeIdentifier(runtimeUuid string) []IOSDevice {
	cmd := exec.Command("xcrun", "simctl", "list", "devices", "--json")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
		panic("could not run command")
	}
	var devicesJSON Devices
	err = json.Unmarshal(out, &devicesJSON)

	if err != nil {
		panic("Unmarshal failed")
	}

	l := devicesJSON.Devices[runtimeUuid]
	CurrentDevices = l
	return l
}

func IOSRuntimes() XCRunDevices {
	cmd := exec.Command("xcrun", "simctl", "list", "runtimes", "--json")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
		panic("could not run command")
	}
	var devices XCRunDevices
	err = json.Unmarshal(out, &devices)

	if err != nil {
		panic("Unmarshal failed")
	}

	return devices
}

func execute(path string, udid string, state string) {
	// Create the command to run the script using osascript
	cmd := exec.Command("osascript", path, udid, state)

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()

	// Handle any errors that occurred
	if err != nil {
		log.Fatalf("Error running AppleScript: %v", err)
	}

	// Print the output of the script
	fmt.Printf("Output:\n%s\n", string(output))
	log.Printf("Output:\n%s\n", string(output))
}

func RunAppleScript(actionName string, udid string, state string) {
	bootPath := "applescript/boot.applescript"
	shutdownPath := "applescript/shutdown.applescript"

	switch actionName {
	case constants.DefaultIOSCommands["Boot"]:
		execute(bootPath, udid, state)
	case constants.DefaultIOSCommands["Shutdown"]:
		execute(shutdownPath, udid, state)
	default:
		log.Fatalf("Error running AppleScript: unknown action name %s", actionName)
	}
}
