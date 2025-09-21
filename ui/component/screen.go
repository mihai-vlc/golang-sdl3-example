package component

import (
	"myapp/ui/context"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

type BaseScreen struct {
	BaseComponent
}

func (s *BaseScreen) Draw(ctx *context.DrawContext) {
	var bg = ctx.Theme.Colors.Background
	sdl.SetRenderDrawColor(ctx.Renderer, bg.R, bg.G, bg.B, bg.A)
	sdl.RenderClear(ctx.Renderer)

	s.BaseComponent.Draw(ctx)
}
