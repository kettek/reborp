package component

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png" //
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite is a component that holds an image.
type Sprite struct {
	image *ebiten.Image
	opts  ebiten.DrawImageOptions
}

// SetImage sets the image.
func (s *Sprite) SetImage(image *ebiten.Image) {
	s.image = image
}

// Image returns the image.
func (s *Sprite) Image() *ebiten.Image {
	return s.image
}

func (s *Sprite) Position(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	s.opts.GeoM.Reset()
	s.opts.GeoM.Translate(-float64(s.image.Bounds().Dx())/2, -float64(s.image.Bounds().Dy())/2)
	s.opts.GeoM.Concat(opts.GeoM)
}

func (s *Sprite) Draw(screen *ebiten.Image, camera *GeoMatrix) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(-float64(s.image.Bounds().Dx())/2, -float64(s.image.Bounds().Dy())/2)
	opts.GeoM.Concat(s.opts.GeoM)
	//s.opts.GeoM.Concat(camera.GeoM())
	screen.DrawImage(s.image, &opts)
}

func MakeSprite(str string) Sprite {
	b, err := os.ReadFile(filepath.Join("..", "data", str))
	if err != nil {
		panic(err)
	}
	image, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return Sprite{ebiten.NewImageFromImage(image), ebiten.DrawImageOptions{}}
}

func NewSprite(str string) *Sprite {
	spr := MakeSprite(str)
	return &spr
}

func (s *Sprite) Chain(chain *Chain, last any) any {
	if comp, ok := last.(GeoMatrix); ok {
		s.opts.GeoM.Reset()
		s.opts.GeoM.Concat(comp.GeoM())
		fmt.Println(s.opts.GeoM)
		//return s
		return last
	}
	return s.opts
}
