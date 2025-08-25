package main

import (
	"log"

	"github.com/TomekPetrykowski/egt/assets"
	"github.com/TomekPetrykowski/egt/game"
	"github.com/TomekPetrykowski/egt/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	assets.MustLoadAssets()
}

func main() {
	config := game.Config{
		WindowHeight: settings.WINDOW_HEIGHT,
		WindowWidth:  settings.WINDOW_WIDTH,
		Title:        "Game Title",
	}
	game := game.NewGame(&config)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
