package entities

import "github.com/hajimehoshi/ebiten/v2"

type Inventory struct {
	Dice  *[]*Dice
	Walls *[]*Wall
}

func (i Inventory) Draw(screen *ebiten.Image) {

}
