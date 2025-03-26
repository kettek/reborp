package component

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Chainable interface {
	Chain(chain *Chain, component any) any
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
		if ch, ok := comp.(Chainable); ok {
			lastValue = ch.Chain(c, lastValue)
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

func (c *Chain) Chain(chain *Chain, last any) any {
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

func (c *Chain) ComponentBefore(k any) any {
	for i, comp := range c.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(k) {
			if i == 0 {
				return nil
			}
			return c.components[i-1]
		}
	}
	return nil
}

func (c *Chain) ComponentAfter(k any) any {
	for i, comp := range c.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(k) {
			if i == len(c.components)-1 {
				return nil
			}
			return c.components[i+1]
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
