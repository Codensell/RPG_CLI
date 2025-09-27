package main

import (
	"github.com/Codensell/RPG_CLI/internal/domain"
	"github.com/Codensell/RPG_CLI/internal/ui/term"
)

func main() {
	stats := domain.Stats{
		Name:      "Player",
		HP:        10,
		MaxHP:     10,
		STR:       2,
		AGI:       3,
		STA:       1,
		Weapon:    "Sword",
		WeaponDMG: 3,
		WeaponTYPE: "Slashing",
	}
	term.DrawFrames(stats)
}
