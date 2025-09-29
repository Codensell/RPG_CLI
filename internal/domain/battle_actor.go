package domain

type actor struct{
	name string
	hp, maxhp int
	str, agi int
	sta int
	weaponDMG int
	weaponType DamageType
	perks Perks
	turnCount int
	poisoned bool
	poisonTurns int
}

func newActorFromStats(s Stats, p Perks) actor{
	return actor{
		name: s.Name,
		hp: s.HP,
		maxhp: s.MaxHP,
		str: s.STR,
		agi: s.AGI,
		sta: s.STA,
		weaponDMG: s.WeaponDMG,
		weaponType: DamageType(s.WeaponTYPE),
		perks: p,
	}
}

func (a *actor) toStats() Stats{
	return Stats{
		Name: a.name,
		HP: a.hp,
		MaxHP: a.maxhp,
		STR:        a.str,
		AGI:        a.agi,
		STA:        a.sta,
		WeaponDMG:  a.weaponDMG,
		WeaponTYPE: string(a.weaponType),
	}
}