package component

type Rect struct {
	width, height int
}

func (s *Rect) SetSize(width, height int) {
	s.width = width
	s.height = height
}

func (s *Rect) Size() (int, int) {
	return s.width, s.height
}

func (s *Rect) Width() int {
	return s.width
}

func (s *Rect) Height() int {
	return s.height
}

type Circle struct {
	radius int
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) Radius() int {
	return c.radius
}
