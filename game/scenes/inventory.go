package scenes

import (
	"image/color"

	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/game/entities"
	"github.com/TomekPetrykowski/egt/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InventoryScene struct {
	loaded       bool
	inventory    *entities.Inventory
	tileSize     int
	gridSizeX    int
	gridSizeY    int
	gridSpace    int
	grid         []*entities.WallSlot
	selectedSlot int
	currentSlot  int
}

func NewInventoryScene() *InventoryScene {
	gridSizeX := 11
	tileSize := settings.INVENTORY_SLOT_SIZE
	gridSizeY := 4
	gridSpace := settings.INVENTORY_GAP
	walls := make([]*entities.Wall, gridSizeX*gridSizeY)
	walls[1] = &entities.Wall{}
	walls[3] = &entities.Wall{Flavor: entities.Bland}
	walls[7] = &entities.Wall{Flavor: entities.Salty} //TODO:load from player data
	grid := make([]*entities.WallSlot, gridSizeX*gridSizeY)
	// for i,dice := {

	// }
	for i, wall := range walls {
		grid[i] = &entities.WallSlot{Wall: wall, Rect: engine.NewRect(float64((tileSize+gridSpace)*(i%gridSizeX)+gridSpace), float64((tileSize+gridSpace)*(i/gridSizeX)+gridSpace), float64(tileSize), float64(tileSize))}
	}
	return &InventoryScene{
		loaded:    false,
		tileSize:  tileSize,
		gridSizeX: gridSizeX,
		gridSizeY: gridSizeY,
		gridSpace: 5,
		inventory: &entities.Inventory{
			Walls: &walls,
			Die:   &[]*entities.Dice{{}},
		},
		grid:         grid,
		selectedSlot: -1,
	}
}

func (s *InventoryScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DebugPrint(screen, "Inventory!")
	for _, slot := range s.grid {
		slot.Draw(screen)
	}
}

func (s *InventoryScene) Update() engine.SceneId {
	xInt, yInt := ebiten.CursorPosition()
	x := float64(xInt)
	y := float64(yInt)
	s.currentSlot = -1
	for i, slot := range s.grid {
		if slot.IsMouseInside(x, y) {
			s.currentSlot = i
			break
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {

		if s.currentSlot != -1 {
			if s.selectedSlot != -1 {
				selectedWall := s.grid[s.selectedSlot].Wall
				s.grid[s.selectedSlot].Wall = s.grid[s.currentSlot].Wall
				s.grid[s.currentSlot].Wall = selectedWall
				s.grid[s.selectedSlot].IsSelected = false
				s.selectedSlot = -1
			} else {
				if s.grid[s.currentSlot].Wall != nil {

					s.selectedSlot = s.currentSlot
					s.grid[s.selectedSlot].IsSelected = true
				}
			}
		} else {
			if s.selectedSlot != -1 {
				s.grid[s.selectedSlot].IsSelected = false
				s.selectedSlot = -1
			}
		}
	}
	return InventorySceneId
}

func (s *InventoryScene) FirstLoad() {
	s.loaded = true

}

func (s *InventoryScene) IsLoaded() bool {
	return s.loaded
}

func (s *InventoryScene) OnEnter() {

}

func (s *InventoryScene) OnExit() {}

var _ engine.Scene = (*InventoryScene)(nil)
