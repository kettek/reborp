package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Mouse struct {
	deltaX, deltaY int
	lastX, lastY   int
}

func (i *Mouse) Delta() (int, int) {
	return i.deltaX, i.deltaY
}

func (i *Mouse) Position() (int, int) {
	return i.lastX, i.lastY
}

func (i *Mouse) Update() {
	x, y := ebiten.CursorPosition()
	i.deltaX = x - i.lastX
	i.deltaY = y - i.lastY
	i.lastX = x
	i.lastY = y
}

func (i *Mouse) Chain(chain *Chain, last any) any {
	if c, ok := chain.Component(&Scale{}).(*Scale); ok {
		x, _ := c.Scale()
		i.deltaX = int(float64(i.deltaX) * x)
		i.deltaY = int(float64(i.deltaY) * x)
	}

	return last
}

func MakeMouse() Mouse {
	return Mouse{}
}

func NewMouse() *Mouse {
	p := MakeMouse()
	return &p
}

func init() {
	Register("Mouse", &Mouse{})
}
