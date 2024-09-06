package android

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"
)

func AVDEmulators() []string {
	cmd := exec.Command("emulator", "-list-avds")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
		panic("could not run command")
	}

	devices := strings.Split(string(out), "\n")

	// emulator -list-avds command result has a new line at the end, hence we need to remove it
	devices = slices.DeleteFunc(devices, func(e string) bool {
		return e == ""
	})

	return devices
}
