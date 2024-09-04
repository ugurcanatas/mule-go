package ios

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

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

type RuntimesModel struct {
	Name       string
	Identifier string
}

type DevicesMap map[string][]RuntimesModel

func GetIOSRuntimes() XCRunDevices {
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
