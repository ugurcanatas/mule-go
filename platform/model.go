package platform

import (
	"mule-go/sharedState"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type PlatformModel struct {
	list list.Model

	model tea.Model
}

func (m *PlatformModel) QuitProgram() {
	sharedState.QuitProgram = true
}

func (u *PlatformModel) SetList(list list.Model) {
	u.list = list
}
