package main

import (
	"math/rand"
	"time"

	"github.com/Codensell/RPG_CLI/internal/domain"
	"github.com/Codensell/RPG_CLI/internal/loot"
	"github.com/Codensell/RPG_CLI/internal/ui/term"
	gc "github.com/rthornton128/goncurses"
)

func perksForPlayer(wLv, bLv, rLv int) domain.Perks {
	var p domain.Perks
	if wLv >= 1 {
		p.ActionSurge = true
	}
	if wLv >= 2 {
		p.Shield = true
	}
	if bLv >= 1 {
		p.Rage = true
	}
	if bLv >= 2 {
		p.StoneSkin = true
	}
	if rLv >= 1 {
		p.SneakAttack = true
	}
	if rLv >= 3 {
		p.Poison = true
	}
	return p
}

func perksForEnemy(name string) domain.Perks {
	var e domain.Perks
	switch name {
	case "Skeleton":
		e.VulnBludge = true
	case "Slime":
		e.ImmuneSlashingWeapon = true
	case "Ghost":
		e.SneakAttack = true
	case "Golem":
		e.StoneSkin = true
	case "Dragon":
		e.DragonBreath = true
	}
	return e
}

func main() {
	rand.Seed(time.Now().UnixNano())

	std, err := gc.Init()
	if err != nil {
		panic(err)
	}
	defer gc.End()
	gc.Echo(false)
	gc.Cursor(0)
	gc.Raw(true)
	std.Keypad(true)

	p := domain.NewCharacter("Player")

	chosen := term.SelectClass(std, p.CharacterData())
	_ = p.ApplyClassLevel(chosen)

	wLv, bLv, rLv := 0, 0, 0
	switch chosen {
	case domain.ClassWarrior:
		wLv++
	case domain.ClassBarbarian:
		bLv++
	case domain.ClassRogue:
		rLv++
	}

	wins := 0

	for wins < 5 {
		e := domain.NewRandomEnemy()
		eStats := e.CharacterData()

		pPerks := perksForPlayer(wLv, bLv, rLv)
		ePerks := perksForEnemy(eStats.Name)

		playerWon := term.RunBattleFrames(std, p.CharacterData(), eStats, pPerks, ePerks)

		if !playerWon {
			wins = 0
			term.ShowDefeat(std)
			p = domain.NewCharacter("Player")
			chosen = term.SelectClass(std, p.CharacterData())
			_ = p.ApplyClassLevel(chosen)
			wLv, bLv, rLv = 0, 0, 0
			switch chosen {
			case domain.ClassWarrior:
				wLv++
			case domain.ClassBarbarian:
				bLv++
			case domain.ClassRogue:
				rLv++
			}
			continue
		}

		wins++

		term.ShowWinProgress(std, wins)

		totalLv := wLv + bLv + rLv
		if totalLv < 3 {
			if up, ok := term.SelectLevel(std, totalLv); ok {
				_ = p.ApplyClassLevel(up)
				switch up {
				case domain.ClassWarrior:
					wLv++
				case domain.ClassBarbarian:
					bLv++
				case domain.ClassRogue:
					rLv++
				}
			}
		}
		drop := loot.Default.Drop(eStats.Name)
		if drop.Name != "" {
			if term.OfferWeaponSwap(std, p.Weapon, drop) {
				p.Weapon = drop
			}
		}
	}

	term.ShowGameWon(std)
}
