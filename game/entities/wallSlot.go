package entities

import (
	"fmt"

	"github.com/TomekPetrykowski/egt/assets"
	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type WallSlot struct {
	Wall          *Wall
	IsHoveredOver bool
	IsSelected    bool
	Rect          *engine.Rect
}

func (w *WallSlot) Draw(screen *ebiten.Image) {
	var image *ebiten.Image
	opts := ebiten.DrawImageOptions{}
	scale := float64(settings.INVENTORY_SLOT_SIZE) / float64(settings.WALL_SPRITE_SIZE)
	opts.GeoM.Scale(scale, scale) //make sure the image is the same size as tileSize

	if w.IsHoveredOver || w.IsSelected {
		opts.GeoM.Scale(1.08, 1.08)
		opts.GeoM.Translate(-w.Rect.Width*0.04, -w.Rect.Height*0.04)
	}

	opts.GeoM.Translate(w.Rect.Pos.Unpack())

	if w.Wall != nil {
		image = GetImageFromFlavor(w.Wall.Flavor)
		ebitenutil.DebugPrintAt(image, fmt.Sprintf("%d", w.Wall.Power), 12, 8) // temp solution, it has to be rendered text on image surface
	} else {
		image = assets.WallEmpty
	}
	screen.DrawImage(image, &opts)
}

func (w *WallSlot) IsMouseInside(x, y float64) bool {
	if w.Rect.IsPointInside(x, y) {
		w.IsHoveredOver = true
	} else {
		w.IsHoveredOver = false
	}
	return w.IsHoveredOver
}
