package platform

import (
	"mule-go/sharedState"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type OSModel struct {
	list list.Model

	model tea.Model
}

func (m *OSModel) QuitProgram() {
	sharedState.QuitProgram = true
}

func (u *OSModel) SetList(list list.Model) {
	u.list = list
}
