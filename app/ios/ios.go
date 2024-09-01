package ios

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type iosmodel struct {
	value string
}

// // Hello returns a greeting for the named person.
// func Hello(name string) string {
// 	// Return a greeting that embeds the name in a message.
// 	message := fmt.Sprintf("Hi, %v. Welcome!", name)
// 	return message
// }

// func RuntimesIOS() {
// 	cmd := exec.Command("xcrun", "simctl", "list", "runtimes", "--json")

// 	// The `Output` method executes the command and
// 	// collects the output, returning its value
// 	out, err := cmd.Output()
// 	if err != nil {
// 		// if there was any error, print it here
// 		fmt.Println("could not run command: ", err)
// 	}
// 	// otherwise, print the output from running the command
// 	fmt.Println("Output: ", string(out))
// }




func (m iosmodel) Init() tea.Cmd {
	return nil
}

func (m iosmodel) View() string {
	return fmt.Sprintf("Hi. This program will exit in %d seconds.\n\nTo quit sooner press ctrl-c, or press ctrl-z to suspend...\n", m)
}

func (m iosmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}

	case string:
		return m, tea.Quit
	}
	return m, nil
}

// Setup initial list model
func InitialModel() iosmodel {
	

	return iosmodel{
		value: "sdafasd",
	}
}
