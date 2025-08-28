package entities

import (
	"fmt"

	"github.com/TomekPetrykowski/egt/assets"
	"github.com/TomekPetrykowski/egt/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DiceContainer struct {
	Dice          *Dice
	Clickable     bool
	IsHoveredOver bool
	LastWall      *Wall
	Rect          *engine.Rect
	Cost          int
}

func (d *DiceContainer) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(d.Rect.Pos.Unpack())
	image := (*ebiten.Image)(nil)
	if d.LastWall != nil {
		image = GetImageFromFlavor(d.LastWall.Flavor)
		screen.DrawImage(image, &opts)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", d.LastWall.Power), int(d.Rect.Pos.X)+12, int(d.Rect.Pos.Y)+8)
	} else {
		image = assets.WallEmpty
		screen.DrawImage(image, &opts)
	}
	// if d.IsHoveredOver {
	// 	opts.GeoM.Translate(0, settings.INVENTORY_SLOT_SIZE*2)
	// 	screen.DrawImage(image, &opts)
	// }
}

func (d *DiceContainer) SetDice(dice *Dice) {
	d.Dice = dice
	d.Cost = dice.CalculateCost()
	d.LastWall = (*dice.Walls)[0]
}

func (d *DiceContainer) IsMouseInside(x, y float64) bool {
	if d.Rect.IsPointInside(x, y) {
		d.IsHoveredOver = true
	} else {
		d.IsHoveredOver = false
	}
	return d.IsHoveredOver
}

func (d *DiceContainer) Roll() *Wall {
	d.LastWall = d.Dice.Roll()
	return d.LastWall

}
