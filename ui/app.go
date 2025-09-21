package ui

import (
	"fmt"
	"myapp/ui/component"
	"myapp/ui/context"
	"myapp/ui/resource"
	"myapp/ui/screen"
	"myapp/ui/theme"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

type App struct {
	running bool

	activeScreen component.Component
	drawContext  *context.DrawContext
	initContext  *context.InitContext
}

func NewApp() *App {
	return &App{
		activeScreen: screen.NewWelcomeScreen(),
		drawContext: &context.DrawContext{
			Theme: theme.DefaultTheme,
		},
		initContext: &context.InitContext{
			Resource: resource.NewResourceManager(),
		},
	}
}

func (a *App) Init() error {
	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		return fmt.Errorf("error setting VSync hint: %s", sdl.GetError())
	}

	if !sdl.Init(sdl.InitVideo) {
		return fmt.Errorf("error initializing video: %s", sdl.GetError())
	}

	if !sdl.CreateWindowAndRenderer("My App", 1280, 720, sdl.WindowResizable, &a.drawContext.Window, &a.drawContext.Renderer) {
		return fmt.Errorf("error creating window: %s", sdl.GetError())
	}

	if err := a.initContext.Resource.Init(); err != nil {
		return err
	}

	a.activeScreen.Init(a.initContext)
	return nil
}

func (a *App) Run() {
	a.running = true

	var lastTick = sdl.GetTicks()
	for a.running {
		var event sdl.Event
		sdl.WaitEventTimeout(&event, 1000)
		a.input(event)

		a.drawContext.Dt = float32(sdl.GetTicks()-lastTick) / 1000.0
		a.draw()
	}
}

func (a *App) Destroy() {
	sdl.DestroyWindow(a.drawContext.Window)
	sdl.DestroyRenderer(a.drawContext.Renderer)
	sdl.Quit()
}

func (a *App) input(event sdl.Event) {
	switch event.Type() {
	case sdl.EventQuit:
		a.running = false
	}

	a.activeScreen.Input(event)
}

func (a *App) draw() {
	a.activeScreen.Draw(a.drawContext)
	sdl.RenderPresent(a.drawContext.Renderer)
}
