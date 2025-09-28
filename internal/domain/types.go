package domain

type DamageType string

const(
	DmgSlashing DamageType = "Slashing"
	DmgPiercing DamageType = "Piercing"
	DmgBludge DamageType = "Bludge"
)

type Weapon struct{
	Name string
	Base int
	Type DamageType
}