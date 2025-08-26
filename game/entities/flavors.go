package entities

import (
	"github.com/TomekPetrykowski/egt/assets"
	"github.com/hajimehoshi/ebiten/v2"
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

func GetImageFromFlavor(id int) *ebiten.Image {
	switch id {
	case Sweet:
		return assets.WallSweet
	case Bitter:
		return assets.WallBitter
	case Sour:
		return assets.WallSour
	case Bland:
		return assets.WallBland
	case Umami:
		return assets.WallUmami
	case Salty:
		return assets.WallSalty
	case Spicy:
		return assets.WallSpicy
	default:
		return assets.WallEmpty
	}
}
