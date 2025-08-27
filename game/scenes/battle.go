package scenes

import (
	"image/color"

	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/game/entities"
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
	opts := ebiten.DrawImageOptions{}
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
