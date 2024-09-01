package main

import (
	"fmt"
	"mule-go/app"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(app.InitialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
	// app.RuntimesIOS()

	// fmt.Println("Enter some text:")
	// var input string

	// // Read a single line of input
	// _, err := fmt.Scan(&input)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("You entered:", input)
}
