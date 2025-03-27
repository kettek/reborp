package component

type Position struct {
	x, y int
}

func MakePosition(x, y int) Position {
	return Position{x, y}
}

func NewPosition(x, y int) *Position {
	p := MakePosition(x, y)
	return &p
}

func (p *Position) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

func (p *Position) Position() (int, int) {
	return p.x, p.y
}

func (p *Position) X() int {
	return p.x
}

func (p *Position) Y() int {
	return p.y
}

func (p *Position) Chain(chain *Chain, last any) any {
	if c, ok := last.(GeoMatrix); ok {
		c.Translate(float64(p.x), float64(p.y))
		return c
	}
	return MakeGeoMatrix(float64(p.x), float64(p.y))
}

func init() {
	Register("Position", &Position{})
}
