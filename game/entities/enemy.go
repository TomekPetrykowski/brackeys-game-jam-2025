package entities

import (
	"github.com/TomekPetrykowski/egt/engine"
	"github.com/TomekPetrykowski/egt/engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	Health         int
	MaxHealth      int
	Dice           *Dice
	IsHoveredOver  bool
	Shield         int
	AttackModifier int
	Rect           *engine.Rect
	Offset         *utils.Vec
	Sprite         *ebiten.Image
}

func (e *Enemy) Buff(power int) {
	e.AttackModifier += power
}

func (e *Enemy) Debuff(power int) {
	e.AttackModifier -= power
}

func (e *Enemy) Defend(power int) {
	e.Shield += power
}

func (e *Enemy) Heal(power int) {
	if e.Health+power >= e.MaxHealth {
		e.Health = e.MaxHealth
	} else {
		e.Health += power
	}
}

func (e *Enemy) Hit(power int) {
	if e.Shield > 0 {
		if e.Shield >= power {
			e.Shield -= power
		} else {
			power -= e.Shield
			e.Health -= power
		}
	} else {
		e.Health -= power
	}
	if e.Health <= 0 {
		//die
	}

}

type BattleActor interface {
	Hit(power int)
	Heal(power int)
	Buff(power int)
	Debuff(power int)
	Defend(power int)
}

func (e *Enemy) Update() {

}

func (e *Enemy) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.Rect.Pos.Unpack())
	// opts.GeoM.Translate(e.Offset.Unpack())
	if e.Health <= 0 {
		opts.ColorScale.SetB(0)
		opts.ColorScale.SetG(0)
	}
	screen.DrawImage(e.Sprite, &opts)
}

func (e *Enemy) IsMouseInside(x, y float64) bool {
	if e.Rect.IsPointInside(x, y) {
		e.IsHoveredOver = true
	} else {
		e.IsHoveredOver = false
	}
	return e.IsHoveredOver
}

func (e *Enemy) Action(player *Player) {
	wall := e.Dice.Roll()
	switch wall.Flavor {
	case Sour:
		player.Hit(wall.Power + e.AttackModifier)
	case Salty:
		player.Hit(wall.Power + e.AttackModifier)
	case Bland:
		e.Defend(wall.Power)
	}
}

var _ BattleActor = (*Enemy)(nil)
