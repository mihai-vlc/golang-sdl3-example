package component

import (
	"myapp/ui/context"
	"myapp/ui/resource"
	"myapp/ui/theme"

	"github.com/jupiterrider/purego-sdl3/sdl"
	"github.com/jupiterrider/purego-sdl3/ttf"
)

type Text struct {
	BaseComponent

	resourceManager *resource.ResourceManager
	texture         *sdl.Texture
	destination     *sdl.FRect
	value           string
}

func NewText(value string) *Text {
	return &Text{
		value:       value,
		destination: &sdl.FRect{},
	}
}

func (t *Text) Init(ctx *context.InitContext) error {
	t.resourceManager = ctx.Resource
	return t.BaseComponent.Init(ctx)
}

func (t *Text) Draw(ctx *context.DrawContext) {
	if t.texture == nil {
		t.createTexture(ctx.Renderer, ctx.Theme)
	}

	var w int32
	var h int32
	sdl.GetRenderOutputSize(ctx.Renderer, &w, &h)
	sdl.GetTextureSize(t.texture, &t.destination.W, &t.destination.H)
	t.destination.X = (float32(w) - t.destination.W) / 2
	t.destination.Y = (float32(h) - t.destination.H) / 2

	sdl.RenderTexture(ctx.Renderer, t.texture, nil, t.destination)
}

func (t *Text) createTexture(renderer *sdl.Renderer, theme theme.Theme) {
	font, _ := t.resourceManager.GetFont(theme.Typography.FontFamily, theme.Typography.FontSizeL)
	var surface = ttf.RenderTextBlended(font, t.value, 0, theme.Colors.TextPrimary)
	t.texture = sdl.CreateTextureFromSurface(renderer, surface)
	sdl.DestroySurface(surface)
}
