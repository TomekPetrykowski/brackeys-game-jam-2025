package entities

import (
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
	opts := ebiten.DrawImageOptions{}
	scale := float64(settings.INVENTORY_SLOT_SIZE) / float64(settings.WALL_SPRITE_SIZE)
	opts.GeoM.Scale(scale, scale) //make sure the image is the same size as tileSize
	if w.IsHoveredOver || w.IsSelected {
		opts.GeoM.Scale(1.1, 1.1)
		opts.GeoM.Translate(-w.Rect.Width*0.05, -w.Rect.Height*0.05)
	}
	opts.GeoM.Translate(w.Rect.Pos.Unpack())
	if w.Wall != nil {
		screen.DrawImage(GetImageFromFlavor(w.Wall.Flavor), &opts)
		//dorysuj numerek
	} else {
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/empty.png")
		screen.DrawImage(image, &opts)

	}
}

func (w *WallSlot) IsMouseInside(x, y float64) bool {
	if w.Rect.IsPointInside(x, y) {
		w.IsHoveredOver = true
	} else {
		w.IsHoveredOver = false
	}
	return w.IsHoveredOver
}
