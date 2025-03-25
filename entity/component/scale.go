package component

type Scale struct {
	x, y float64
}

func MakeScale(x, y float64) Scale {
	return Scale{x, y}
}

func NewScale(x, y float64) *Scale {
	p := MakeScale(x, y)
	return &p
}

func (p *Scale) SetScale(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Scale) Scale() (float64, float64) {
	return p.x, p.y
}

func (p *Scale) X() float64 {
	return p.x
}

func (p *Scale) Y() float64 {
	return p.y
}

func (p *Scale) Chain(chain *Chain, last any) any {
	if c, ok := last.(GeoMatrix); ok {
		c.Scale(p.x, p.y)
		return c
	}
	return MakeGeoMatrix(p.x, p.y)
}
