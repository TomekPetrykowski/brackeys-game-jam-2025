package assets

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed graphics/walls/*
	wallsFS embed.FS

	// walss sprites
	WallBitter *ebiten.Image
	WallBland  *ebiten.Image
	WallEmpty  *ebiten.Image
	WallUmami  *ebiten.Image
	WallSalty  *ebiten.Image
	WallSour   *ebiten.Image
	WallSpicy  *ebiten.Image
	WallSweet  *ebiten.Image
)

func MustLoadAssets() {
	WallBitter = mustNewEbitenImage(mustLoadWallImage("bitter.png"))
	WallBland = mustNewEbitenImage(mustLoadWallImage("bland.png"))
	WallEmpty = mustNewEbitenImage(mustLoadWallImage("empty.png"))
	WallUmami = mustNewEbitenImage(mustLoadWallImage("umami.png"))
	WallSalty = mustNewEbitenImage(mustLoadWallImage("salty.png"))
	WallSour = mustNewEbitenImage(mustLoadWallImage("sour.png"))
	WallSpicy = mustNewEbitenImage(mustLoadWallImage("spicy.png"))
	WallSweet = mustNewEbitenImage(mustLoadWallImage("sweet.png"))
}

func mustLoadWallImage(filename string) []byte {
	imageData, err := fs.ReadFile(wallsFS, "graphics/walls/"+filename)
	if err != nil {
		panic(fmt.Sprintf("Could not find wall image file: %s", filename))
	}

	return imageData
}

func mustNewEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
