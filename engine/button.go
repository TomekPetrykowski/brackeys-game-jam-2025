package engine

import (
	"image/color"

	"github.com/TomekPetrykowski/egt/engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
	IsHoveredOver bool
	IsClicked     bool
	Sprite        *ebiten.Image
	ClickSprite   *ebiten.Image
	HoverSprite   *ebiten.Image //possible improvement:add option to pass a funtion that modifies Sprite insted of another image
	// FontSize      int
	Text    string //possible improvement: add a modular system for creating a button, so it can have an image instead of text etc or
	Rect    *Rect
	OnClick *func()
	// OnRelease *func()
}

func (b *Button) Update() {
	x, y := ebiten.CursorPosition()
	if b.Rect.IsPointInside(float64(x), float64(y)) { //possible improvement :remove seperate type conversion for every button, its only important if there is A LOT of buttons
		b.IsHoveredOver = true //^its not really an issue
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			b.IsClicked = true
			if b.OnClick != nil {
				(*b.OnClick)()
			}
		}
		if b.IsClicked && inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			b.IsClicked = false
		}
	} else {
		b.IsClicked = false
		b.IsHoveredOver = false
	} //add smart button pressing //if you dont understand what i mean then ignore this
}

func (b *Button) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.Rect.Pos.Unpack())
	if b.IsClicked && b.ClickSprite != nil {
		screen.DrawImage(b.ClickSprite, opts)
	} else if b.IsHoveredOver && b.HoverSprite != nil {
		screen.DrawImage(b.HoverSprite, opts)
	} else {
		if b.Sprite == nil {
			sprite := utils.CreatePlaceholderImage(&utils.PlaceholderImage{Width: int(b.Rect.Width), Height: int(b.Rect.Height), Color: color.RGBA{100, 100, 100, 255}})
			screen.DrawImage(sprite, opts)
		} else {
			screen.DrawImage(b.Sprite, opts)
		}
	}
	if b.Text != "" {
		length := len(b.Text)
		x := int(b.Rect.Width)/2 - ((length * 6) / 2)
		y := int(b.Rect.Height)/2 - 8 //this works for now for centering but should probably be changed later
		ebitenutil.DebugPrintAt(screen, b.Text, int(b.Rect.Pos.X)+x, int(b.Rect.Pos.Y)+y)
	}

}

func NewButton(tect *Rect, text string) *Button {
	button := Button{}
	return &button
}
