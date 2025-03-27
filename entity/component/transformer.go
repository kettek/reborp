package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Transformer struct {
	adjuster any
	adjustee any
	cb       func(chain *Chain, adjuster, adjustee any)
}

func (c *Transformer) Draw(screen *ebiten.Image, camera *GeoMatrix) {
	if c, ok := c.adjustee.(ChainableDrawable); ok {
		c.Draw(screen, camera)
	}
}

func (c *Transformer) Update() {
	if c, ok := c.adjuster.(ChainableUpdate); ok {
		c.Update()
	}
	if c, ok := c.adjustee.(ChainableUpdate); ok {
		c.Update()
	}
}

func (c *Transformer) Chain(chain *Chain, last any) any {
	c.cb(chain, c.adjuster, c.adjustee)
	if c, ok := c.adjuster.(Chainable); ok {
		c.Chain(chain, last)
	}
	if c, ok := c.adjustee.(Chainable); ok {
		return c.Chain(chain, last)
	}
	return last
}

func MakeTransformer(adjuster, adjustee any, cb func(chain *Chain, adjuster, adjustee any)) Transformer {
	return Transformer{
		adjuster: adjuster,
		adjustee: adjustee,
		cb:       cb,
	}
}

func NewTransformer(adjuster, adjustee any, cb func(chain *Chain, adjuster, adjustee any)) *Transformer {
	p := MakeTransformer(adjuster, adjustee, cb)
	return &p
}
