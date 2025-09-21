package context

import (
	"myapp/ui/resource"
	"myapp/ui/theme"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

type DrawContext struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Dt       float32
	Theme    theme.Theme
}

type InitContext struct {
	Resource *resource.ResourceManager
}
