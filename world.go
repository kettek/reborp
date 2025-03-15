package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/reborp/acts"
	"github.com/kettek/reborp/entity"
	"github.com/kettek/reborp/entity/component"
	"github.com/kettek/reborp/entity/factory"
)

type Entity interface {
	Update() []acts.Action
}

type DrawableEntity interface {
	Entity
	Draw(screen *ebiten.Image, matrix *component.GeoMatrix)
}

type World struct {
	entities         []Entity
	drawableEntities []DrawableEntity
	camera           *entity.Camera
}

func NewWorld() *World {
	world := &World{}
	camera := factory.CreateEntity("Camera", component.MakePosition(100, 100)).(*entity.Camera)
	world.camera = camera
	world.AddEntity(camera)
	return world
}

func (w *World) Update() {
	actions := []acts.Action{}
	for _, e := range w.entities {
		actions = append(actions, e.Update()...)
	}
	for _, a := range actions {
		fmt.Println(a)
	}
}

func (w *World) Draw(screen *ebiten.Image) {
	w.camera.SetSize(screen.Bounds().Dx(), screen.Bounds().Dy())

	for _, e := range w.drawableEntities {
		e.Draw(screen, &w.camera.GeoMatrix)
	}
}

func (w *World) AddEntity(e Entity) {
	w.entities = append(w.entities, e)
	if de, ok := e.(DrawableEntity); ok {
		w.drawableEntities = append(w.drawableEntities, de)
	}
}
