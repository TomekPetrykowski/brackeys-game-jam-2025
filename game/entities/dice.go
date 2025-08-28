package entities

import "math/rand"

type Dice struct {
	Walls *[]*Wall
	Cost  int
}

func (d *Dice) CalculateCost() int {
	cost := 0
	for _, wall := range *d.Walls {
		cost += wall.Cost
	}
	d.Cost = cost
	return cost
}

func (d *Dice) Roll() *Wall {
	index := rand.Int() % len(*d.Walls)
	return (*d.Walls)[index]
}
