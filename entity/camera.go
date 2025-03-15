package entity

import (
	"github.com/kettek/reborp/acts"
	"github.com/kettek/reborp/entity/component"
	"github.com/kettek/reborp/entity/factory"
)

type Camera struct {
	component.Position
	component.Rect
	component.GeoMatrix
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) Update() []acts.Action {
	c.Reset()
	c.Translate(-float64(c.X()), -float64(c.Y()))
	c.Translate(float64(c.Rect.Width())/2, float64(c.Rect.Height())/2)
	return nil
}

func (c *Camera) Component(k any) any {
	switch k.(type) {
	case component.Position:
		return c.Position
	case component.Rect:
		return c.Rect
	case component.GeoMatrix:
		return c.GeoMatrix
	}
	return nil
}

func (c *Camera) Components() []any {
	return []any{c.Position, c.Rect, c.GeoMatrix}
}

func (c *Camera) SetComponent(comp any) {
	switch comp := comp.(type) {
	case component.Position:
		c.Position = comp
	case component.Rect:
		c.Rect = comp
	case component.GeoMatrix:
		c.GeoMatrix = comp
	}
}

func init() {
	factory.RegisterEntityFunc("Camera", func(components ...any) factory.Entity {
		camera := NewCamera()

		for _, comp := range components {
			camera.SetComponent(comp)
		}

		return camera
	})
}
