package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	scenes        map[SceneId]Scene
	activeSceneId SceneId
}

func (s *SceneManager) AddScene(id SceneId, scene Scene) {
	if s.scenes == nil {
		s.scenes = make(map[SceneId]Scene)
	}

	s.scenes[id] = scene
}

func (s *SceneManager) SetActiveSceneId(id SceneId) {
	s.activeSceneId = id
}

func (s *SceneManager) GetActiveSceneId() SceneId {
	return s.activeSceneId
}

func (s *SceneManager) Update() error {
	nextSceneId := s.scenes[s.activeSceneId].Update()

	if nextSceneId != s.activeSceneId {
		nextScene := s.scenes[nextSceneId]
		// if not loaded then load in
		if !nextScene.IsLoaded() {
			nextScene.FirstLoad()
		}

		nextScene.OnEnter()
		s.scenes[s.activeSceneId].OnExit()
	}

	s.activeSceneId = nextSceneId

	return nil
}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	s.scenes[s.activeSceneId].Draw(screen)
}
