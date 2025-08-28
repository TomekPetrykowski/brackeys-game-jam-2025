package entities

import "math/rand"

type Dice struct {
	Walls *[]*Wall
}

func (d *Dice) GetCost() int {
	cost := 0
	for _, wall := range *d.Walls {
		cost += wall.Cost
	}
	return cost
}

func (d *Dice) Roll() *Wall {
	index := rand.Int() % len(*d.Walls)
	return (*d.Walls)[index]
}
