package resource

import (
	"fmt"

	"github.com/jupiterrider/purego-sdl3/sdl"
	"github.com/jupiterrider/purego-sdl3/ttf"
)

type ResourceManager struct {
	fontCache map[string]*ttf.Font
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		fontCache: map[string]*ttf.Font{},
	}
}

func (r *ResourceManager) Init() error {
	if !ttf.Init() {
		return fmt.Errorf("error initializing font: %s", sdl.GetError())
	}
	return nil
}

func (r *ResourceManager) GetFont(name string, fontSize float32) (*ttf.Font, error) {
	var key = fmt.Sprintf("%s-%f", name, fontSize)

	font, ok := r.fontCache[key]

	if !ok {
		font = ttf.OpenFont(fmt.Sprintf("resources/%s", name), fontSize)

		if font == nil {
			return nil, fmt.Errorf("error loading font: %s", sdl.GetError())
		}

		r.fontCache[key] = font
	}

	return font, nil
}
