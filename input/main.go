package input

import (
	input "github.com/quasilyte/ebitengine-input"
)

var System input.System

func Update() {
	System.Update()
}

func init() {
	System.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
}
