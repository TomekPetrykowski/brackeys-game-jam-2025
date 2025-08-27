package game

import (
	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/game/scenes"
	"github.com/TomekPetrykowski/egt/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

type Config struct {
	WindowHeight     int
	WindowWidth      int
	WindowResizeable bool
	Title            string
}

type Game struct {
	sceneManager *engine.SceneManager
}

func NewGame(config *Config) *Game {
	g := Game{sceneManager: &engine.SceneManager{}}

	// Add new scenes here
	g.sceneManager.AddScene(scenes.StartSceneId, scenes.NewStartScene())
	g.sceneManager.AddScene(scenes.InventorySceneId, scenes.NewInventoryScene())
	g.sceneManager.AddScene(scenes.BattleSceneId, scenes.NewBattleScene())
	g.sceneManager.SetActiveSceneId(scenes.BattleSceneId)
	g.Set(config)

	return &g
}

func (g *Game) SetSceneManager(s *engine.SceneManager) {
	g.sceneManager = s
}

func (g *Game) Update() error {
	return g.sceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / settings.SCALE, outsideHeight / settings.SCALE
}

func (g *Game) Set(opts *Config) {
	ebiten.SetWindowSize(opts.WindowWidth, opts.WindowHeight)
	ebiten.SetWindowTitle(opts.Title)
	if opts.WindowResizeable {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	}
}
