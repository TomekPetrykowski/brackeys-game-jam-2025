package entities

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
