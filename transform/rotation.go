package transform

import "github.com/hajimehoshi/ebiten/v2"

type Rotation struct {
	r    float64
	next any
}

func (r *Rotation) SetRotation(rotation float64) {
	r.r = rotation
}

func (r *Rotation) Rotation() float64 {
	return r.r
}

func (r *Rotation) TransformGeoMatrix(geom *ebiten.GeoM) {
	geom.Rotate(r.r)
}
