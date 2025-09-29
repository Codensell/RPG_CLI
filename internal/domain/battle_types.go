package domain

type EventKind int

const(
	EvStart EventKind = iota
	EvInitiative
	EvTurnBegin
	EvPoisonTick
	EvAttackRoll
	EvMiss
	EvHit
	EvVictory
)

type DamageBreakdown struct{
	Weapon int
	STR int
	Bonus int
	WeaponAfterPartial int
	Shield int
	StoneSkin int
	Applied int
}

type Event struct{
	Kind EventKind
	Who string
	Target string
	Note string
	Roll int
	Dmg *DamageBreakdown
}

type Perks struct{
	ActionSurge bool
	Rage bool
	SneakAttack bool
	DragonBreath bool
	Poison bool
	Shield bool
	StoneSkin bool
	ImmuneSlashingWeapon bool
	VulnBludge bool
}