package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Sweet = iota
	Sour
	Salty
	Bland
	Bitter
	Spicy
	Umami
)

func GetImageFromFlavor(id int) *ebiten.Image { //TODO: this is badly made :c
	switch id {
	case Sweet:
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/sweet.png")
		return image
	case Bitter:
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/bitter.png")
		return image
	case Sour:
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/sour.png")
		return image
	case Bland:
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/bland.png")
		return image
	case Umami:
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/umami.png")
		return image
	case Salty:
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/salty.png")
		return image
	case Spicy:
		image, _, _ := ebitenutil.NewImageFromFile("assets/graphics/walls/spicy.png")
		return image
	default:
		return nil
	}
}
