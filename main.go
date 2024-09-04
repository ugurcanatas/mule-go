package main

import (
	"fmt"
	"mule-go/platform"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tea.LogToFile("debug.log", "debug")

	p := tea.NewProgram(platform.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
