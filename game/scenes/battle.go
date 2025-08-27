package scenes

import (
	"image/color"

	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/game/entities"
	"github.com/TomekPetrykowski/egt/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BattleScene struct {
	enemies []*entities.Enemy
	loaded  bool
}

func NewBattleScene() *BattleScene {
	return &BattleScene{
		loaded: false,
	}
}

func (s *BattleScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	background, _, _ := ebitenutil.NewImageFromFile("assets/graphics/backgrounds/dungeon.png")
	scaleX := float64(settings.WINDOW_WIDTH) / (settings.BACKGROUND_WIDTH * 2)
	scaleY := float64(settings.WINDOW_HEIGHT) / (settings.BACKGROUND_HEIGHT * 2)
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(background, &opts)
	ebitenutil.DebugPrint(screen, "Battle!")
}

func (s *BattleScene) Update() engine.SceneId {
	return BattleSceneId
}

func (s *BattleScene) FirstLoad() {
	s.loaded = true
}

func (s *BattleScene) IsLoaded() bool {
	return s.loaded
}

func (s *BattleScene) OnEnter() {}

func (s *BattleScene) OnExit() {}

var _ engine.Scene = (*BattleScene)(nil)
