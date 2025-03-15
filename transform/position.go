package transform

import "github.com/hajimehoshi/ebiten/v2"

type Node interface {
	Next() Node
}

type Position struct {
	x, y int
	next any
}

func (p *Position) TransformGeoMatrix(geom *ebiten.GeoM) {
	geom.Translate(float64(p.x), float64(p.y))
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

/*
nodes := MakeNodeChain(
	PositionTransformer(&position),
	RotationTransformer(&rotation),
	ScaleTransformer(&scale),
)

NodeChain().
*/
