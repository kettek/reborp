package entity

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/reborp/acts"
	"github.com/kettek/reborp/entity/component"
	"github.com/kettek/reborp/entity/factory"
)

// Character is an entity that represents a character.
type Character struct {
	component.Name
	component.Position
	component.Rotation
	component.Circle
	component.Sprite
}

func (c *Character) Update() []acts.Action {
	c.Rotation.Update()
	return nil
}

func (c *Character) Draw(screen *ebiten.Image, camera *component.GeoMatrix) {
	if c.Image() == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	// Draw from center of image.
	op.GeoM.Translate(-float64(c.Image().Bounds().Dx())/2, -float64(c.Image().Bounds().Dy())/2)
	op.GeoM.Scale(10, 10)
	op.GeoM.Rotate(c.Rotation.Rotation())
	op.GeoM.Translate(float64(c.X()), float64(c.Y()))
	op.GeoM.Concat(camera.GeoM())
	screen.DrawImage(c.Image(), op)
}

func (c *Character) Component(k any) any {
	switch k.(type) {
	case component.Name:
		return &c.Name
	case component.Position:
		return &c.Position
	case component.Rotation:
		return &c.Rotation
	case component.Circle:
		return &c.Circle
	case component.Sprite:
		return &c.Sprite
	}
	return nil
}

func (c *Character) Components() []any {
	return []any{c.Name, c.Position, c.Rotation, c.Circle, c.Sprite}
}

func (c *Character) SetComponent(comp any) {
	switch comp := comp.(type) {
	case component.Name:
		c.Name = comp
	case component.Position:
		c.Position = comp
	case component.Rotation:
		c.Rotation = comp
	case component.Circle:
		c.Circle = comp
	case component.Sprite:
		c.Sprite = comp
	default:
		fmt.Printf("Unknown component %+T %+v\n", comp, comp)
	}
}

type Action interface {
	Execute()
}

func init() {
	defImage := ebiten.NewImage(32, 32)
	vector.DrawFilledCircle(defImage, 16, 16, 16, color.NRGBA{255, 0, 255, 255}, true)

	factory.RegisterEntityFunc("Character", func(components ...any) factory.Entity {
		character := &Character{}
		character.SetImage(defImage)

		for _, comp := range components {
			character.SetComponent(comp)
		}

		return character
	})
}
