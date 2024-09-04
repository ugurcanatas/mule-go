package ios

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"slices"
)

// Credit goes to https://mholt.github.io/json-to-go/
type XCRunDevices struct {
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

type FilteredIOSDeviceList struct {
	currentDevices []IOSDevice
}

func (m *FilteredIOSDeviceList) DeviceByDeviceUDID(udid string) IOSDevice {
	index := slices.IndexFunc(m.currentDevices, func(d IOSDevice) bool {
		return d.Udid == udid
	})
	return m.currentDevices[index]
}

func NewDevicesSlice() *FilteredIOSDeviceList {
	inital := []IOSDevice{}
	return &FilteredIOSDeviceList{
		currentDevices: inital,
	}
}

var Ddevices FilteredIOSDeviceList = *NewDevicesSlice()

type RuntimesModel struct {
	Name       string
	Identifier string
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
	Ddevices.currentDevices = l
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
	bootPath := "app/ios/boot.applescript"
	shutdownPath := "app/ios/shutdown.applescript"

	switch actionName {
	case "Boot":
		execute(bootPath, udid, state)
	case "Shutdown":
		execute(shutdownPath, udid, state)
	default:
		log.Fatalf("Error running AppleScript: unknown action name %s", actionName)
	}
}
