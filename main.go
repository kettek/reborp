package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/reborp/entity/component"
	"github.com/kettek/reborp/entity/factory"
)

type game struct {
	world *World
}

func (g *game) Update() error {
	g.world.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &game{
		world: NewWorld(),
	}

	g.world.AddEntity(factory.CreateEntity("Character", component.MakePosition(100, 100), component.MakeRotation(0, 0.05), component.MakeSprite("character.png")))

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

	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
