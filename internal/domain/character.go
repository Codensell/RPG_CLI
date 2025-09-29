package domain

import "math/rand"

type Character struct {
	Name  string
	HP    int
	MaxHP int
	STR   int
	AGI   int
	STA   int

	LvlRogue     int
	LvlWarrior   int
	LvlBarbarian int

	Weapon Weapon

	HasSneakAttack bool
	HasPoison      bool

	HasSurge   bool
	HasShield  bool

	HasRage      bool
	HasStoneSkin bool
}

func (c *Character) CharacterData() Stats {
	return Stats{
		Name:       c.Name,
		HP:         c.HP,
		MaxHP:      c.MaxHP,
		STR:        c.STR,
		AGI:        c.AGI,
		STA:        c.STA,
		Weapon:     c.Weapon.Name,
		WeaponDMG:  c.Weapon.Base,
		WeaponTYPE: string(c.Weapon.Type),
	}
}

func NewCharacter(name string) Character{
	c := Character{
		Name: name,
		STR: 1 + rand.Intn(3),
		AGI: 1 + rand.Intn(3),
		STA: 1 + rand.Intn(3),

		MaxHP: 1,
		HP: 1,
	}
	return c
}
