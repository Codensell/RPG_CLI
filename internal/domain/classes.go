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
		return fmt.Errorf("Total lvl limit reached (3)")
	}

	switch class{
	case ClassRogue:
		c.LvlRogue++
		c.MaxHP += 4
		if totalBefore == 0{
			c.Weapon = Weapon{Name: "Dager", Base: 2, Type: DamageType("piercing")}
		}
	case ClassWarrior:
		c.LvlWarrior++
		c.MaxHP += 5
		if totalBefore == 0 {
			c.Weapon = Weapon{Name: "Sword", Base: 3, Type: DamageType("slashing")}
		}
	case ClassBarbarian:
		c.LvlBarbarian++
		c.MaxHP += 6
		if totalBefore == 0 {
			c.Weapon = Weapon{Name: "Club", Base: 3, Type: DamageType("bludgeoning")}
		}
	default:
		return fmt.Errorf("unknown class: %s", class)
	}

	c.HP = c.MaxHP
	return nil
	
}
