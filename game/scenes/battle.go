package scenes

import (
	"fmt"
	"image/color"

	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/game/entities"
	"github.com/TomekPetrykowski/egt/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type BattleScene struct {
	player     *entities.Player
	enemies    []*entities.Enemy
	playerDice []*entities.DiceContainer
	loaded     bool
	rolledWall *entities.Wall
}

func ExampleDice(numOfWalls, flavor int) *entities.Dice {
	diewalls := make([]*entities.Wall, numOfWalls)
	for i := 0; i < numOfWalls; i++ {
		diewalls[i] = &entities.Wall{Power: i + 1, Flavor: flavor, Cost: i}
	}
	return &entities.Dice{Walls: &diewalls}
}

func NewBattleScene() *BattleScene {

	die1 := ExampleDice(4, entities.Sour)
	die2 := ExampleDice(6, entities.Sour)
	enemies := make([]*entities.Enemy, 2)
	enemySprite1, _, _ := ebitenutil.NewImageFromFile("assets/graphics/enemies/lemon.png")
	enemySprite2, _, _ := ebitenutil.NewImageFromFile("assets/graphics/enemies/lemondemon.png")
	enemy1 := &entities.Enemy{Dice: die1, Health: 20, MaxHealth: 20, Sprite: enemySprite1, Rect: engine.NewRect(100, 100, 100, 100)}
	enemy2 := &entities.Enemy{Dice: die2, Health: 30, MaxHealth: 30, Sprite: enemySprite2, Rect: engine.NewRect(200, 100, 100, 100)}
	enemies[0] = enemy1
	enemies[1] = enemy2
	inventory := &entities.Inventory{}
	incentoryDice := make([]*entities.Dice, 7)
	for i := range 7 {
		incentoryDice[i] = ExampleDice(6, i)
	}
	inventory.Dice = &incentoryDice
	player := &entities.Player{Inventory: inventory, MaxHealth: 20, Health: 10}
	playerDice := make([]*entities.DiceContainer, len(*player.Inventory.Dice))
	diceStartX := (settings.WINDOW_WIDTH/settings.SCALE - ((settings.INVENTORY_SLOT_SIZE + settings.INVENTORY_GAP) * len(*player.Inventory.Dice))) / 2
	diceStartY := settings.WINDOW_HEIGHT/settings.SCALE - (settings.INVENTORY_SLOT_SIZE + settings.INVENTORY_GAP)
	for i, dice := range *player.Inventory.Dice {
		diceSlot := &entities.DiceContainer{Rect: engine.NewRect(float64(diceStartX+(settings.INVENTORY_SLOT_SIZE+settings.INVENTORY_GAP)*i), float64(diceStartY), settings.INVENTORY_SLOT_SIZE, settings.INVENTORY_SLOT_SIZE)}
		diceSlot.SetDice(dice)
		playerDice[i] = diceSlot
	}

	return &BattleScene{
		loaded:     false,
		enemies:    enemies,
		player:     player,
		playerDice: playerDice,
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
	for _, enemy := range s.enemies {
		enemy.Draw(screen)
	}

	for _, dice := range s.playerDice {
		dice.Draw(screen)
	}

	if s.rolledWall != nil {
		ebitenutil.DebugPrint(screen, "Click an enemy to attack")
	}
	// opts.GeoM.Reset()
	// opts.GeoM.Translate(0, ((settings.WINDOW_HEIGHT / settings.SCALE) - 20))
	// text.Draw(screen,"HP")

}

func (s *BattleScene) Update() engine.SceneId {
	xInt, yInt := ebiten.CursorPosition()
	x := float64(xInt)
	y := float64(yInt)
	currentDice := -1
	currentEnemy := -1

	for i, dice := range s.playerDice {
		if dice.IsMouseInside(x, y) {
			currentDice = i
		}
	}
	for i, enemy := range s.enemies {
		if enemy.IsMouseInside(x, y) {
			currentEnemy = i
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		if currentDice != -1 {
			wall := s.playerDice[currentDice].Roll()
			targets := make([]entities.BattleActor, 0)
			switch wall.Flavor {
			case entities.Salty:
				for _, enemy := range s.enemies {
					targets = append(targets, enemy)
				}
				s.player.Action(wall, targets)
			case entities.Sweet, entities.Bland, entities.Spicy, entities.Umami:
				targets = append(targets, s.player)
				s.player.Action(wall, targets)
			case entities.Sour, entities.Bitter:
				s.rolledWall = wall
				print("choose an enemy")
			}
			fmt.Println(s.player)
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		if currentEnemy != -1 && s.rolledWall != nil {
			target := make([]entities.BattleActor, 1)
			target[0] = s.enemies[currentEnemy]
			s.player.Action(s.rolledWall, target)
			s.rolledWall = nil
			print(s.enemies[currentEnemy].Health)
		}
	}
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
