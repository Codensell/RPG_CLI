package domain

import "fmt"

type Class string

const (
	ClassWarrior   Class = "Warrior"
	ClassBarbarian Class = "Barbarian"
	ClassRogue     Class = "Rogue"
)

func (c *Character) ApplyClassLevel(class Class) error {
	totalBefore := c.LvlRogue + c.LvlWarrior + c.LvlBarbarian
	if totalBefore >= 3 {
		return fmt.Errorf("total lvl limit reached (3)")
	}

	hpPerLevel := func(cl Class) int {
		switch cl {
		case ClassRogue:
			return 4
		case ClassWarrior:
			return 5
		case ClassBarbarian:
			return 6
		default:
			return 0
		}
	}

	switch class {
	case ClassRogue:
		c.LvlRogue++
		if totalBefore == 0 && c.Weapon.Base == 0 {
			c.Weapon = Weapon{Name: "Dagger", Base: 2, Type: DmgPiercing}
		}
		switch c.LvlRogue {
		case 1:
			c.HasSneakAttack = true
		case 2:
			c.AGI += 1
		case 3:
			c.HasPoison = true
		}

	case ClassWarrior:
		c.LvlWarrior++
		if totalBefore == 0 && c.Weapon.Base == 0 {
			c.Weapon = Weapon{Name: "Sword", Base: 3, Type: DmgSlashing}
		}
		switch c.LvlWarrior {
		case 1:
			c.HasSurge = true
		case 2:
			c.HasShield = true
		case 3:
			c.STR += 1
		}

	case ClassBarbarian:
		c.LvlBarbarian++
		if totalBefore == 0 && c.Weapon.Base == 0 {
			c.Weapon = Weapon{Name: "Club", Base: 3, Type: DmgBludge}
		}
		switch c.LvlBarbarian {
		case 1:
			c.HasRage = true
		case 2:
			c.HasStoneSkin = true
		case 3:
			c.STA += 1
		}

	default:
		return fmt.Errorf("unknown class: %s", class)
	}

	c.MaxHP += hpPerLevel(class) + c.STA
	c.HP = c.MaxHP

	return nil
}

