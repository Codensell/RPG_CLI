package domain

import "fmt"

type Class string

const (
	ClassWarrior   Class = "Warrior"
	ClassBarbarian Class = "Barbarian"
	ClassRogue     Class = "Rogue"
)

// ApplyClassLevel adds exactly one level of the given class.
// Rules:
// - total level must be <= 3
// - HP gain per level: Rogue +4, Warrior +5, Barbarian +6
// - if this is the first overall level (pre-total == 0), set starter weapon
// - HP is restored to MaxHP after the change

func (c *Character) ApplyClassLevel(class Class) error {
	totalBefore := c.LvlRogue + c.LvlWarrior + c.LvlBarbarian
	if totalBefore >= 3 {
		return fmt.Errorf("total lvl limit reached (3)")
	}

	switch class {
	case ClassRogue:
		prev := c.LvlRogue
		c.LvlRogue++
		c.MaxHP += 4
		if totalBefore == 0 {
			c.Weapon = Weapon{Name: "Dager", Base: 2, Type: DamageType("piercing")}
		}
		switch prev {
		case 0:
			c.HasSneakAttack = true
		case 1:
			c.AGI += 1
		case 2:
			c.HasPoison = true
		}
	case ClassWarrior:
		prev := c.LvlWarrior
		c.LvlWarrior++
		c.MaxHP += 5
		if totalBefore == 0 {
			c.Weapon = Weapon{Name: "Sword", Base: 3, Type: DamageType("slashing")}
		}
		switch prev {
		case 0:
			c.HasSurge = true
		case 1:
			c.HasShield = true
		case 2:
			c.STR += 1
		}
	case ClassBarbarian:
		prev := c.LvlBarbarian
		c.LvlBarbarian++
		c.MaxHP += 6
		if totalBefore == 0 {
			c.Weapon = Weapon{Name: "Club", Base: 3, Type: DamageType("bludgeoning")}
		}
		switch prev {
		case 0:
			c.HasRage = true
		case 1:
			c.HasStoneSkin = true
		case 2:
			c.STA += 1
		}
	default:
		return fmt.Errorf("unknown class: %s", class)
	}

	c.HP = c.MaxHP
	return nil

}
func (c *Character) TotalLevel() int {
	return c.LvlRogue + c.LvlWarrior + c.LvlBarbarian
}
