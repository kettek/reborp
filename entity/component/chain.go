package component

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Chainable interface {
	Chain(component any) any
}

type ChainableUpdate interface {
	Chainable
	Update()
}

type ChainableDrawable interface {
	Chainable
	Draw(screen *ebiten.Image, camera *GeoMatrix)
}

type Chain struct {
	components []any
}

func (c *Chain) Update() {
	var lastValue any
	for _, comp := range c.components {
		if c, ok := comp.(ChainableUpdate); ok {
			c.Update()
		}
		if c, ok := comp.(Chainable); ok {
			lastValue = c.Chain(lastValue)
		}
	}
}

func (c *Chain) Draw(screen *ebiten.Image, camera *GeoMatrix) {
	for _, comp := range c.components {
		if c, ok := comp.(ChainableDrawable); ok {
			c.Draw(screen, camera)
		}
	}
}

func (c *Chain) Chain(last any) any {
	c.Update()
	return last
}

func (c *Chain) Component(k any) any {
	for _, comp := range c.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(k) {
			return comp
		}
	}
	return nil
}

func MakeChain(components ...any) Chain {
	return Chain{components}
}

func NewChain(components ...any) *Chain {
	c := MakeChain(components...)
	return &c
}
