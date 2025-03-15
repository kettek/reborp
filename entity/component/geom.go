package component

import "github.com/hajimehoshi/ebiten/v2"

type GeoMatrix struct {
	geom ebiten.GeoM
}

func (g *GeoMatrix) Reset() {
	g.geom.Reset()
}

func (g *GeoMatrix) Translate(x, y float64) {
	g.geom.Translate(x, y)
}

func (g *GeoMatrix) Scale(x, y float64) {
	g.geom.Scale(x, y)
}

func (g *GeoMatrix) Rotate(theta float64) {
	g.geom.Rotate(theta)
}

func (g *GeoMatrix) GeoM() ebiten.GeoM {
	return g.geom
}

func (g *GeoMatrix) Chain(last any) any {
	g.Reset()
	return *g
}

// MakeGeoMatrix creates a GeoMatrix with the given values. scale x, scale y, rotate, translate x, translate y
func MakeGeoMatrix(v ...float64) GeoMatrix {
	g := GeoMatrix{}
	for i, v := range v {
		switch i {
		case 0:
			g.Scale(v, 1)
		case 1:
			g.Scale(1, v)
		case 2:
			g.Rotate(v)
		case 3:
			g.Translate(v, 0)
		case 4:
			g.Translate(0, v)
		}
	}
	return g
}

func NewGeoMatrix(v ...float64) *GeoMatrix {
	g := MakeGeoMatrix(v...)
	return &g
}
