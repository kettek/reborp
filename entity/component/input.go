package component

import (
	input "github.com/kettek/reborp/input"
	einput "github.com/quasilyte/ebitengine-input"
)

type Input struct {
	handler *einput.Handler
}

func (i *Input) Update() {
}

func (i *Input) Chain(chain *Chain, last any) any {
	return last
}

func (i *Input) ActionIsPressed(act einput.Action) bool {
	return i.handler.ActionIsPressed(act)
}

func MakeInput(pid uint8, keymap einput.Keymap) Input {
	return Input{
		handler: input.System.NewHandler(pid, keymap),
	}
}

func NewInput(pid uint8, keymap einput.Keymap) *Input {
	p := MakeInput(pid, keymap)
	return &p
}
