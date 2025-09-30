package loot

import "github.com/Codensell/RPG_CLI/internal/domain"

type LootTable interface {
	Drop(enemyName string) domain.Weapon
}

type GDDLoot struct{}

func (GDDLoot) Drop(name string) domain.Weapon {
	switch name {
	case "Goblin":
		return domain.Weapon{Name: "Dagger", Base: 2, Type: domain.DmgPiercing}
	case "Skeleton":
		return domain.Weapon{Name: "Club", Base: 3, Type: domain.DmgBludge}
	case "Slime":
		return domain.Weapon{Name: "Spear", Base: 3, Type: domain.DmgPiercing}
	case "Ghost":
		return domain.Weapon{Name: "Sword", Base: 3, Type: domain.DmgSlashing}
	case "Golem":
		return domain.Weapon{Name: "Axe", Base: 4, Type: domain.DmgSlashing}
	case "Dragon":
		return domain.Weapon{Name: "Legendary Sword", Base: 10, Type: domain.DmgSlashing}
	default:
		return domain.Weapon{}
	}
}

var Default LootTable = GDDLoot{}
