package component

import (
	"myapp/ui/context"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

type Component interface {
	Init(ctx *context.InitContext) error
	Input(event sdl.Event)
	Draw(ctx *context.DrawContext)
	Destroy()
	AddChild(child Component)
	SetParent(parent Component)
}

type BaseComponent struct {
	parent   Component
	children []Component
}

func (c *BaseComponent) Init(ctx *context.InitContext) error {
	for _, child := range c.children {
		if err := child.Init(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c *BaseComponent) Input(event sdl.Event) {
	for _, child := range c.children {
		child.Input(event)
	}
}

func (c *BaseComponent) Draw(ctx *context.DrawContext) {
	for _, child := range c.children {
		child.Draw(ctx)
	}
}

func (c *BaseComponent) Destroy() {
	for _, child := range c.children {
		child.Destroy()
	}
}

func (c *BaseComponent) AddChild(child Component) {
	c.children = append(c.children, child)
	child.SetParent(c)
}

func (c *BaseComponent) SetParent(parent Component) {
	c.parent = parent
}
