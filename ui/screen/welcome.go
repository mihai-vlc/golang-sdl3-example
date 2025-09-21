package screen

import (
	"myapp/ui/component"
)

type WelcomeScreen struct {
	component.BaseScreen
}

func NewWelcomeScreen() *WelcomeScreen {
	var s = &WelcomeScreen{}

	var msg = component.NewText("Welcome!")
	s.AddChild(msg)

	return s
}
