package utils

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type PlaceholderImage struct {
	Width, Height int
	Color         color.Color
}

var DefaultPlaceholder = PlaceholderImage{50, 50, color.RGBA{255, 0, 0, 255}}

func CreatePlaceholderImage(placeholder *PlaceholderImage) *ebiten.Image {
	img := ebiten.NewImage(placeholder.Width, placeholder.Height)
	img.Fill(placeholder.Color)

	return img
}

func LoadImage(path string, errorMsg string, placeholder *PlaceholderImage) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		fmt.Println(errorMsg)
		return CreatePlaceholderImage(placeholder)
	}

	return img
}

// Wrapper for casting subimage to ebiten.Image
func SubImage(spriteImg *ebiten.Image, img image.Rectangle) *ebiten.Image {
	return spriteImg.SubImage(img).(*ebiten.Image)
}
