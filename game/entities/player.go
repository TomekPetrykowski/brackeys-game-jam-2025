package entities

type Player struct {
	Health         int
	MaxHealth      int
	Inventory      *Inventory
	IsHoveredOver  bool
	Shield         int
	AttackModifier int
	Mana           int
	MaxMana        int
}

// Buff implements BattleActor.
func (p *Player) Buff(power int) {
	p.AttackModifier += power
}

// Debuff implements BattleActor.
func (p *Player) Debuff(power int) {
	p.AttackModifier -= power
}

// Defend implements BattleActor.
func (p *Player) Defend(power int) {
	p.Shield += power
}

// Heal implements BattleActor.
func (p *Player) Heal(power int) {
	if p.Health+power >= p.MaxHealth {
		p.Health = p.MaxHealth
	} else {
		p.Health += power
	}
}

// Hit implements BattleActor.
func (p *Player) Hit(power int) {
	if p.Shield > 0 {
		if p.Shield >= power {
			p.Shield -= power
		} else {
			power -= p.Shield
			p.Health -= power
		}
	} else {
		p.Health -= power
	}
	if p.Health <= 0 {
		//die
	}
}

func (p *Player) Roll(diceId int) *Wall {
	return (*p.Inventory.Dice)[diceId].Roll()
}

func (p *Player) AddMana(amount int) {
	p.Mana += amount
	if p.Mana > p.MaxMana {
		p.Mana = p.MaxMana
	}
}

func (p *Player) Action(wall *Wall, targets []BattleActor) {
	switch wall.Flavor {
	case Sour:
		for _, target := range targets {
			target.Hit(wall.Power + p.AttackModifier)
		}
	case Salty:
		for _, target := range targets {
			target.Hit(wall.Power + p.AttackModifier)
		}
	case Bland:
		for _, target := range targets {
			target.Defend(wall.Power)
		}
	case Spicy:
		for _, target := range targets {
			target.Defend(wall.Power)
		}
	case Sweet:
		for _, target := range targets {
			target.Buff(wall.Power)
		}
	case Umami:
		for _, target := range targets {
			target.Heal(wall.Power)
		}
	case Bitter:
		for _, target := range targets {
			target.Debuff(wall.Power)
		}
	}
}

var _ BattleActor = (*Player)(nil)
