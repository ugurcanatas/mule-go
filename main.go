package main

import (
	"fmt"
	"mule-go/app"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// type viewModel struct {
// 	currentModel tea.Model
// }

// func (m viewModel) Init() tea.Cmd {
// 	return nil
// }

// func (v viewModel) View() string {
//     str := "This is the first screen. Press any key to switch to the second screen."
//     return str
// }

// func (v viewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
//     case tea.KeyMsg:
//         switch msg.Type {
//         case tea.KeyCtrlC:
//             return v, tea.Quit
//         default:
//             // any other key switches the screen
//             return screenTwo().Update(msg)
//         }
//     default:
//         return v, nil
//     }
// }

func main() {
	p := tea.NewProgram(app.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
