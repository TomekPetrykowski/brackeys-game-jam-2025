package engine

import "github.com/hajimehoshi/ebiten/v2"

type SceneId uint

type Scenes map[SceneId]Scene

type Scene interface {
	Update() SceneId
	Draw(screen *ebiten.Image)
	FirstLoad()
	OnEnter()
	OnExit()
	IsLoaded() bool
}
