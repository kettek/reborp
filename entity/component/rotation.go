package component

type Rotation struct {
	r   float64
	inc float64
}

func MakeRotation(r float64, inc float64) Rotation {
	return Rotation{r, inc}
}

func NewRotation(r, inc float64) *Rotation {
	return &Rotation{r, inc}
}

func (p *Rotation) Update() {
	p.r += p.inc
}

func (p *Rotation) UpdateGeoMatrix(g GeoMatrix) GeoMatrix {
	g.Rotate(p.r)
	return g
}

func (p *Rotation) SetRotation(r float64) {
	p.r = r
}

func (p *Rotation) Rotation() float64 {
	return p.r
}

func (p *Rotation) SetInc(inc float64) {
	p.inc = inc
}

func (p *Rotation) Inc() float64 {
	return p.inc
}

func (p *Rotation) Chain(last any) any {
	if c, ok := last.(GeoMatrix); ok {
		c.Rotate(p.r)
		return c
	}
	g := MakeGeoMatrix()
	g.Rotate(p.r)
	return g
}
