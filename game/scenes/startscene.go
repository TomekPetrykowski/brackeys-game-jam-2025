package scenes

import (
	"image/color"

	"github.com/TomekPetrykowski/egt/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type StartScene struct {
	loaded bool
}

func NewStartScene() *StartScene {
	return &StartScene{
		loaded: false,
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DebugPrint(screen, "This game is awesome!")
}

func (s *StartScene) Update() engine.SceneId {
	return StartSceneId
}

func (s *StartScene) FirstLoad() {
	s.loaded = true
}

func (s *StartScene) IsLoaded() bool {
	return s.loaded
}

func (s *StartScene) OnEnter() {}

func (s *StartScene) OnExit() {}

var _ engine.Scene = (*StartScene)(nil)
