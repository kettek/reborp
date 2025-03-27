package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/reborp/entity/component"
	"github.com/kettek/reborp/entity/factory"
	"github.com/kettek/reborp/input"
	einput "github.com/quasilyte/ebitengine-input"
)

type game struct {
	world *World
}

func (g *game) Update() error {
	input.Update()
	g.world.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

const (
	ActionMoveLeft einput.Action = iota
	ActionMoveRight
	ActionGrow
	ActionShrink
)

func main() {
	g := &game{
		world: NewWorld(),
	}

	ent := g.world.AddEntity(factory.CreateEntity("Character", component.MakePosition(0, 0), component.MakeRotation(0, 0.05), component.MakeSprite("character.png")))
	ent.Component(component.Position{}).(*component.Position).SetPosition(100, 100)

	g.world.AddEntity(factory.CreateEntity("Character", component.MakePosition(100, 100), component.MakeRotation(0.5, 0.01), component.MakeSprite("character.png")))

	g.world.AddEntity(factory.CreateEntity("Dynamic",
		component.NewChain(
			component.NewGeoMatrix(),
			component.NewScale(4, 4),
			component.NewRotation(0.5, 0.1),
			component.NewPosition(300, 300),
			component.NewSprite("character.png"),
			component.NewGeoMatrix(),
			component.NewScale(1, 1),
			component.NewRotation(0.5, 0.1),
			component.NewPosition(100, 100),
			component.NewSprite("character.png"),
		),
	))

	g.world.AddEntity(factory.CreateEntity("Dynamic",
		component.NewChain(
			component.NewGeoMatrix(),
			component.NewTransformer(
				component.NewInput(0, einput.Keymap{ActionGrow: {einput.KeyUp}, ActionShrink: {einput.KeyDown}}),
				component.NewScale(4, 4),
				func(adjuster, adjustee any) {
					inp := adjuster.(*component.Input)
					scale := adjustee.(*component.Scale)
					if inp.ActionIsPressed(ActionGrow) {
						scale.SetScale(scale.X()+0.1, scale.Y()+0.1)
					} else if inp.ActionIsPressed(ActionShrink) {
						scale.SetScale(scale.X()-0.1, scale.Y()-0.1)
					}
				},
			),
			component.NewRotation(0.1, 0.01),
			component.NewTransformer(
				component.NewMouse(),
				component.NewPosition(300, 300),
				func(adjuster any, adjustee any) {
					mouse := adjuster.(*component.Mouse)
					pos := adjustee.(*component.Position)
					pos.SetPosition(mouse.Position())
				},
			),
			component.NewSpriteStack("autoCannon.png", "top", "attack"),
		),
	))

	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
