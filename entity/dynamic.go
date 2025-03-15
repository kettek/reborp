package entity

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/reborp/acts"
	"github.com/kettek/reborp/entity/component"
	"github.com/kettek/reborp/entity/factory"
)

// Dynamic is an entity that represents a character.
type Dynamic struct {
	components []any
}

func (c *Dynamic) Update() []acts.Action {
	for _, comp := range c.components {
		if c, ok := comp.(component.ChainableUpdate); ok {
			c.Update()
		}
	}
	return nil
}

func (c *Dynamic) Draw(screen *ebiten.Image, camera *component.GeoMatrix) {
	for _, comp := range c.components {
		if c, ok := comp.(component.ChainableDrawable); ok {
			c.Draw(screen, camera)
		}
	}
}

func (c *Dynamic) Component(k any) any {
	for _, comp := range c.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(k) {
			return comp
		}
	}
	return nil
}

func (c *Dynamic) Components() []any {
	return c.components
}

func (c *Dynamic) SetComponent(comp any) {
	for i, comp2 := range c.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(comp2) {
			c.components[i] = comp
			return
		}
	}
	c.components = append(c.components, comp)
}

func init() {
	factory.RegisterEntityFunc("Dynamic", func(components ...any) factory.Entity {
		character := &Dynamic{}
		for _, comp := range components {
			character.SetComponent(comp)
		}
		return character
	})
}
