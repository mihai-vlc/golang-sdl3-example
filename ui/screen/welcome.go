package screen

import (
	"myapp/ui/component"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

type WelcomeScreen struct {
	components []component.Component
	background sdl.Color
}

func NewWelcomeScreen() *WelcomeScreen {
	var s = &WelcomeScreen{}

	s.background = sdl.Color{R: 100, G: 150, B: 200, A: 255}

	return s
}

func (s *WelcomeScreen) Init() {
	for _, c := range s.components {
		c.Init()
	}
}

func (s *WelcomeScreen) Input(event sdl.Event) {
	for _, c := range s.components {
		c.Input(event)
	}
}

func (s *WelcomeScreen) Update(dt float32) {
	for _, c := range s.components {
		c.Update(dt)
	}
}

func (s *WelcomeScreen) Draw(renderer *sdl.Renderer) {
	sdl.SetRenderDrawColor(renderer, s.background.R, s.background.G, s.background.B, s.background.A)
	sdl.RenderClear(renderer)

	for _, c := range s.components {
		c.Draw(renderer)
	}

	sdl.RenderPresent(renderer)
}

func (s *WelcomeScreen) Destroy() {
	for _, c := range s.components {
		c.Destroy()
	}
}
