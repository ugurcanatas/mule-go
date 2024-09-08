package android

import (
	"log"
	"os/exec"
	"slices"
	"strings"
)

func Execute(path string, device string) {
	// Create the command to run the script using osascript
	cmd := exec.Command("osascript", path, device)

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()

	// Handle any errors that occurred
	if err != nil {
		log.Fatalf("Error running AppleScript: %v", err)
	}

	// Print the output of the script
	log.Printf("Output:\n%s\n", string(output))
}

func AVDEmulators() []string {
	cmd := exec.Command("emulator", "-list-avds")

	out, err := cmd.Output()
	if err != nil {
		panic("could not run command")
	}

	devices := strings.Split(string(out), "\n")

	// emulator -list-avds command result has a new line at the end, hence we need to remove it
	devices = slices.DeleteFunc(devices, func(e string) bool {
		return e == ""
	})

	return devices
}
