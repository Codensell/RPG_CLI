package domain

import "math/rand"

type EnemyKind string

const (
	EnemyGoblin   EnemyKind = "Goblin"
	EnemySkeleton EnemyKind = "Skeleton"
	EnemySlime    EnemyKind = "Slime"
	EnemyGhost    EnemyKind = "Ghost"
	EnemyGolem    EnemyKind = "Golem"
	EnemyDragon   EnemyKind = "Dragon"
)

// боевая логика будет читать их и применять эффекты / battle logic reads and applies these 

type Enemy struct {
	Kind   EnemyKind
	Name   string
	HP     int
	MaxHP  int
	STR    int
	AGI    int
	STA    int
	Weapon Weapon

	VulnBludge           bool
	ImmuneSlashingWeapon bool
	HasSneakAttack       bool
	HasStoneSkin         bool
	HasDragonBreath      bool

	TurnCounter int

	// Лут/loot
	Drop Weapon
}

// Ghost/Golem HP не указаны поэтому сделал свои / Ghost/Golem have no HP as per task, made myself

func NewEnemy(k EnemyKind) Enemy {
	switch k {
	case EnemyGoblin:
		return Enemy{
			Kind:  k, Name: "Goblin",
			MaxHP: 5, HP: 5, STR: 1, AGI: 1, STA: 1,
			Weapon: Weapon{Name: "Dagger", Base: 2, Type: DmgPiercing},
			Drop:   Weapon{Name: "Dagger", Base: 2, Type: DmgPiercing},
		}
	case EnemySkeleton:
		return Enemy{
			Kind:  k, Name: "Skeleton",
			MaxHP: 10, HP: 10, STR: 1, AGI: 1, STA: 1,
			Weapon:     Weapon{Name: "Club", Base: 3, Type: DmgBludge},
			VulnBludge: true,
			Drop:       Weapon{Name: "Club", Base: 3, Type: DmgBludge},
		}
	case EnemySlime:
		return Enemy{
			Kind:  k, Name: "Slime",
			MaxHP: 8, HP: 8, STR: 1, AGI: 1, STA: 1,
			Weapon:               Weapon{Name: "Spear", Base: 3, Type: DmgPiercing},
			ImmuneSlashingWeapon: true,
			Drop:                 Weapon{Name: "Spear", Base: 3, Type: DmgPiercing},
		}
	case EnemyGhost:
		return Enemy{
			Kind:  k, Name: "Ghost",
			MaxHP: 8, HP: 8, STR: 1, AGI: 2, STA: 1,
			Weapon:         Weapon{Name: "Sword", Base: 3, Type: DmgSlashing},
			HasSneakAttack: true,
			Drop:           Weapon{Name: "Sword", Base: 3, Type: DmgSlashing},
		}
	case EnemyGolem:
		return Enemy{
			Kind:  k, Name: "Golem",
			MaxHP: 14, HP: 14, STR: 2, AGI: 1, STA: 3,
			Weapon:       Weapon{Name: "Axe", Base: 4, Type: DmgSlashing},
			HasStoneSkin: true,
			Drop:         Weapon{Name: "Axe", Base: 4, Type: DmgSlashing},
		}
	case EnemyDragon:
		return Enemy{
			Kind:  k, Name: "Dragon",
			MaxHP: 20, HP: 20, STR: 3, AGI: 2, STA: 3,
			Weapon:          Weapon{Name: "Claws", Base: 4, Type: DmgSlashing},
			HasDragonBreath: true,
			Drop:            Weapon{Name: "Legendary Sword", Base: 10, Type: DmgSlashing},
		}
	default:
		return NewEnemy(EnemyGoblin)
	}
}

func NewRandomEnemy() Enemy {
	kinds := []EnemyKind{EnemyGoblin, EnemySkeleton, EnemySlime, EnemyGhost, EnemyGolem, EnemyDragon}
	return NewEnemy(kinds[rand.Intn(len(kinds))])
}

func (e Enemy) CharacterData() Stats {
	return Stats{
		Name:       e.Name,
		HP:         e.HP,
		MaxHP:      e.MaxHP,
		STR:        e.STR,
		AGI:        e.AGI,
		STA:        e.STA,
		Weapon:     e.Weapon.Name,
		WeaponDMG:  e.Weapon.Base,
		WeaponTYPE: string(e.Weapon.Type),
	}
}
