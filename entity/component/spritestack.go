package component

import (
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/ebistack"
)

type SpriteStack struct {
	stack ebistack.Sprite
	opts  ebiten.DrawImageOptions
}

func (s *SpriteStack) Update() {
	s.stack.Update()
}

func (s *SpriteStack) Draw(screen *ebiten.Image, camera *GeoMatrix) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Concat(s.opts.GeoM)
	s.stack.Draw(screen, &opts)
}

func MakeSpriteStack(sheet, stack, anim string) SpriteStack {
	if sheet, ok := sheets[sheet]; ok {
		sprite := ebistack.MakeSprite(sheet, stack, anim)
		return SpriteStack{sprite, ebiten.DrawImageOptions{}}
	}
	b, err := os.ReadFile(filepath.Join("..", "data", sheet))
	if err != nil {
		panic(err)
	}
	sh, err := ebistack.NewSheetFromStaxie(b)
	if err != nil {
		panic(err)
	}
	sheets[sheet] = &sh
	sprite := ebistack.MakeSprite(sheets[sheet], stack, anim)
	return SpriteStack{sprite, ebiten.DrawImageOptions{}}
}

func NewSpriteStack(sheet, stack, anim string) *SpriteStack {
	spr := MakeSpriteStack(sheet, stack, anim)
	return &spr
}

func (s *SpriteStack) Chain(chain *Chain, last any) any {
	if comp, ok := last.(GeoMatrix); ok {
		s.opts.GeoM.Reset()

		x, _ := chain.Component(&Scale{}).(*Scale).Scale()
		s.stack.SliceDistance = x

		s.opts.GeoM.Concat(comp.GeoM())
		return last
	}
	return s.opts
}

var sheets map[string]*ebistack.Sheet

func init() {
	sheets = make(map[string]*ebistack.Sheet)
}
