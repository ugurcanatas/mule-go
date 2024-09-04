package ios

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

// TODO Change name and implement an interface for more generic usage
type RuntimesModel struct {
	Name       string
	Identifier string
}

// To hold the value of currently selected runtime devices slice
type FilteredIOSDeviceList []IOSDevice

var CurrentDevices FilteredIOSDeviceList = *NewDevicesSlice()
