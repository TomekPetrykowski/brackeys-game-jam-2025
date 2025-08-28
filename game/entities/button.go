package entities

import (
	"image/color"

	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Button struct {
	Rect  *engine.Rect
	Text  string
	Color color.Color
} //this struct is a needs an update later

func (b *Button) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.Rect.Pos.Unpack())
	image := utils.CreatePlaceholderImage(&utils.PlaceholderImage{Width: int(b.Rect.Width), Height: int(b.Rect.Height), Color: b.Color})
	ebitenutil.DebugPrintAt(image, b.Text, 0, 0)
	screen.DrawImage(image, &opts)
}
