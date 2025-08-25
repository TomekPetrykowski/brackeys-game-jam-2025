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

func ExampleInventory(numOfWalls int) *entities.Inventory {
	walls := make([]*entities.Wall, numOfWalls)
	walls[1] = &entities.Wall{}
	walls[3] = &entities.Wall{Flavor: entities.Bland}
	walls[7] = &entities.Wall{Flavor: entities.Salty} //TODO:load from player data
	inv_dice := make([]*entities.Dice, 3)
	diewalls1 := make([]*entities.Wall, 6)
	for i, _ := range diewalls1 {
		diewalls1[i] = &entities.Wall{Power: i + 1}
	}
	diewalls2 := make([]*entities.Wall, 6)
	for i, _ := range diewalls2 {
		diewalls2[i] = &entities.Wall{Power: i + 1, Flavor: 1}
	}
	diewalls3 := make([]*entities.Wall, 4)
	for i, _ := range diewalls3 {
		diewalls3[i] = &entities.Wall{Power: i + 1, Flavor: 2}
	}
	die1 := &entities.Dice{Walls: &diewalls1}
	die2 := &entities.Dice{Walls: &diewalls2}
	die3 := &entities.Dice{Walls: &diewalls3}
	inv_dice[0] = die1
	inv_dice[1] = die2
	inv_dice[2] = die3
	return &entities.Inventory{Dice: &inv_dice, Walls: &walls}

}

func NewInventoryScene() *InventoryScene {
	gridSizeX := 4
	tileSize := settings.INVENTORY_SLOT_SIZE
	gridSizeY := 8
	gridSpace := settings.INVENTORY_GAP

	diceWalls := 0
	inventory := ExampleInventory(gridSizeX * gridSizeY)
	for _, dice := range *inventory.Dice {
		diceWalls += len(*dice.Walls)
	}
	grid := make([]*entities.WallSlot, gridSizeX*gridSizeY+diceWalls)
	current := 0
	for i, dice := range *inventory.Dice {
		for j, wall := range *dice.Walls {
			grid[current] = &entities.WallSlot{Wall: wall, Rect: engine.NewRect(float64((tileSize+gridSpace)*(j)+gridSpace), float64((tileSize+gridSpace)*(i)+gridSpace), float64(tileSize), float64(tileSize))}
			current++
		}

	}
	for i, wall := range *inventory.Walls {
		grid[i+diceWalls] = &entities.WallSlot{Wall: wall, Rect: engine.NewRect(float64((tileSize+gridSpace)*(i%gridSizeX+7)+gridSpace), float64((tileSize+gridSpace)*(i/gridSizeX)+gridSpace), float64(tileSize), float64(tileSize))}
	}
	return &InventoryScene{
		loaded:       false,
		tileSize:     tileSize,
		gridSizeX:    gridSizeX,
		gridSizeY:    gridSizeY,
		gridSpace:    5,
		inventory:    inventory,
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
