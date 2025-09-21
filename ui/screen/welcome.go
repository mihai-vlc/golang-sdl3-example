package screen

import (
	"myapp/ui/component"
)

type WelcomeScreen struct {
	component.BaseScreen
}

func NewWelcomeScreen() *WelcomeScreen {
	var s = &WelcomeScreen{}

	var msg = component.NewText("Hello Moto")
	s.AddChild(msg)

	return s
}
