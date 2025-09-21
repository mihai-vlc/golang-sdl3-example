package ui

import (
	"fmt"
	"myapp/ui/screen"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

type App struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	running  bool

	activeScreen screen.Screen
}

func NewApp() *App {
	return &App{
		activeScreen: screen.NewWelcomeScreen(),
	}
}

func (a *App) Init() error {
	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		return fmt.Errorf("error setting VSync hint: %s", sdl.GetError())
	}

	if !sdl.Init(sdl.InitVideo) {
		return fmt.Errorf("error initializing video: %s", sdl.GetError())
	}

	if !sdl.CreateWindowAndRenderer("Hello, World!", 1280, 720, sdl.WindowResizable, &a.window, &a.renderer) {
		return fmt.Errorf("error creating window: %s", sdl.GetError())
	}

	a.activeScreen.Init()

	return nil
}

func (a *App) Run() {
	a.running = true

	var lastTick = sdl.GetTicks()
	for a.running {
		var event sdl.Event
		sdl.WaitEventTimeout(&event, 1000)
		a.input(event)

		var dt = float32(sdl.GetTicks()-lastTick) / 1000.0
		a.update(dt)
		a.draw()
	}
}

func (a *App) Destroy() {
	sdl.DestroyWindow(a.window)
	sdl.DestroyRenderer(a.renderer)
	sdl.Quit()
}

func (a *App) input(event sdl.Event) {
	switch event.Type() {
	case sdl.EventQuit:
		a.running = false
	}

	a.activeScreen.Input(event)
}

func (a *App) update(dt float32) {
	a.activeScreen.Update(dt)
}

func (a *App) draw() {
	a.activeScreen.Draw(a.renderer)
}
