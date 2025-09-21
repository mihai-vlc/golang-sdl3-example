package component

import "github.com/jupiterrider/purego-sdl3/sdl"

type Component interface {
	Init()
	Input(event sdl.Event)
	Update(dt float32)
	Draw(renderer *sdl.Renderer)
	Destroy()
}
