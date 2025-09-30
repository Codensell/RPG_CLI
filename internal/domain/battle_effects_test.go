package domain

import (
	"math/rand"
	"testing"
)

func rng() *rand.Rand { return rand.New(rand.NewSource(1)) }

func stats(name string, hp, str, agi, sta, wdm int, wtype DamageType) Stats {
	return Stats{
		Name:       name,
		HP:         hp,
		MaxHP:      hp,
		STR:        str,
		AGI:        agi,
		STA:        sta,
		WeaponDMG:  wdm,
		WeaponTYPE: string(wtype),
	}
}

// ждём первый попавшийся EvHit от заданного атакующего
func hitBy(t *testing.T, b *Battle, who string) Event {
	t.Helper()
	for i := 0; i < 200; i++ {
		evs, _ := b.Step()
		for _, ev := range evs {
			if ev.Kind == EvHit && ev.Who == who {
				return ev
			}
		}
	}
	t.Fatalf("не дождались удара от %s", who)
	return Event{}
}

// ждём первый EvPoisonTick для нужного актёра
func poisonTickFor(t *testing.T, b *Battle, who string) Event {
	t.Helper()
	for i := 0; i < 400; i++ {
		evs, _ := b.Step()
		for _, ev := range evs {
			if ev.Kind == EvPoisonTick && ev.Who == who {
				return ev
			}
		}
	}
	t.Fatalf("не дождались PoisonTick для %s", who)
	return Event{}
}

func Test_Initiative_TieBreak_PlayerFirst(t *testing.T) {
	p := stats("Player", 100, 0, 3, 0, 1, DmgSlashing)
	e := stats("Enemy", 100, 0, 3, 0, 1, DmgSlashing)
	b := NewBattle(p, e, Perks{}, Perks{}, rng())

	evs, _ := b.Step()
	found := false
	for _, ev := range evs {
		if ev.Kind == EvTurnBegin && ev.Who == "Player" {
			found = true
		}
	}
	if !found {
		t.Fatalf("при равном AGI первым должен ходить Player")
	}
}

func Test_ActionSurge_Ignores_Slime_Immunity(t *testing.T) {
	p := stats("Player", 100, 0, 5, 0, 3, DmgSlashing)
	e := stats("Enemy", 100, 0, 0, 0, 0, DmgSlashing) // AGI_def=0 => попадание всегда
	pPerks := Perks{ActionSurge: true}
	ePerks := Perks{ImmuneSlashingWeapon: true}

	b := NewBattle(p, e, pPerks, ePerks, rng())
	ev := hitBy(t, b, "Player")

	if ev.Dmg.WeaponAfterPartial != 0 {
		t.Fatalf("оружейная часть против Slime должна быть 0, got %d", ev.Dmg.WeaponAfterPartial)
	}
	if ev.Dmg.Bonus != 3 {
		t.Fatalf("ActionSurge должен дать +3 (бонус = weaponDMG), got %d", ev.Dmg.Bonus)
	}
	if ev.Dmg.Applied != 3 {
		t.Fatalf("итоговый урон должен быть 3 (0 + STR0 + AS3), got %d", ev.Dmg.Applied)
	}
}

func Test_Skeleton_Vulnerability_Doubles_Weapon_Only(t *testing.T) {
	p := stats("Player", 100, 2, 5, 0, 3, DmgBludge)
	e := stats("Skeleton", 100, 0, 0, 0, 0, DmgSlashing)
	ePerks := Perks{VulnBludge: true}

	b := NewBattle(p, e, Perks{}, ePerks, rng())
	ev := hitBy(t, b, "Player")

	if ev.Dmg.WeaponAfterPartial != 6 {
		t.Fatalf("оружейная часть должна удвоиться до 6, got %d", ev.Dmg.WeaponAfterPartial)
	}
	if ev.Dmg.Applied != 8 {
		t.Fatalf("итог = 6 (оруж) + 2 (STR) = 8, got %d", ev.Dmg.Applied)
	}
}

func Test_Rage_Bonus_On_Turns(t *testing.T) {
	// считаем только СВОИ удары игрока; враг может промахиваться — нам не важно
	p := stats("Player", 200, 0, 5, 0, 2, DmgSlashing)
	e := stats("Enemy", 200, 0, 0, 0, 0, DmgSlashing)
	pPerks := Perks{Rage: true}

	b := NewBattle(p, e, pPerks, Perks{}, rng())

	ev1 := hitBy(t, b, "Player")
	if ev1.Dmg.Bonus != 2 || ev1.Dmg.Applied != 4 {
		t.Fatalf("turn1: bonus +2, applied 2+0+2=4; got bonus %d applied %d", ev1.Dmg.Bonus, ev1.Dmg.Applied)
	}
	ev2 := hitBy(t, b, "Player")
	if ev2.Dmg.Bonus != 2 {
		t.Fatalf("turn2: bonus +2, got %d", ev2.Dmg.Bonus)
	}
	ev3 := hitBy(t, b, "Player")
	if ev3.Dmg.Bonus != 2 {
		t.Fatalf("turn3: bonus +2, got %d", ev3.Dmg.Bonus)
	}
	ev4 := hitBy(t, b, "Player")
	if ev4.Dmg.Bonus != -1 || ev4.Dmg.Applied != 1 {
		t.Fatalf("turn4: bonus -1, applied 2-1=1; got bonus %d applied %d", ev4.Dmg.Bonus, ev4.Dmg.Applied)
	}
}

func Test_DragonBreath_Every_Third_Turn(t *testing.T) {
	// делаем защитника с AGI=0, чтобы на 3-м ХОДУ был гарантированный хит
	p := stats("Player", 300, 0, 0, 0, 0, DmgSlashing) // AGI_def=0
	e := stats("Dragon", 300, 0, 5, 0, 2, DmgSlashing) // ходит первым
	ePerks := Perks{DragonBreath: true}

	b := NewBattle(p, e, Perks{}, ePerks, rng())

	_ = hitBy(t, b, "Dragon") // 1-й ход дракона
	_ = hitBy(t, b, "Dragon") // 2-й ход дракона
	ev := hitBy(t, b, "Dragon") // 3-й ход дракона — должен быть бонус

	if ev.Dmg.Bonus != 3 {
		t.Fatalf("на каждом 3-м ходе дракона bonus должен быть +3, got %d", ev.Dmg.Bonus)
	}
}

func Test_Shield_and_StoneSkin_Global_Defense(t *testing.T) {
	// Защитник: STR=5 (>2) => Shield сработает; STA=4 => каменная кожа 4
	p := stats("Player", 100, 2, 5, 0, 3, DmgSlashing)
	e := stats("Enemy", 100, 5, 0, 4, 0, DmgSlashing) // STR_def=5 !!
	ePerks := Perks{Shield: true, StoneSkin: true}

	b := NewBattle(p, e, Perks{}, ePerks, rng())
	ev := hitBy(t, b, "Player")

	if ev.Dmg.Shield != 3 || ev.Dmg.StoneSkin != 4 {
		t.Fatalf("должны примениться Shield=3 и StoneSkin=4, got shield=%d stone=%d", ev.Dmg.Shield, ev.Dmg.StoneSkin)
	}
	expect := (3 /*weapon*/ + 2 /*STR*/ + 0 /*bonus*/) - 3 - 4
	if expect < 0 { expect = 0 }
	if ev.Dmg.Applied != expect {
		t.Fatalf("неверный итоговый урон: expect %d, got %d", expect, ev.Dmg.Applied)
	}
}

func Test_MinDamage_ZeroFloor(t *testing.T) {
	p := stats("Player", 100, 1, 5, 0, 1, DmgSlashing)  // суммарно 2
	e := stats("Tank",   100, 5, 0, 5, 0, DmgSlashing)  // Shield=3 (5>1) + Stone=5
	ePerks := Perks{Shield: true, StoneSkin: true}

	b := NewBattle(p, e, Perks{}, ePerks, rng())
	ev := hitBy(t, b, "Player")
	if ev.Dmg.Applied != 0 {
		t.Fatalf("после редукций урон не должен уходить ниже 0, got %d", ev.Dmg.Applied)
	}
}

func Test_Poison_Ticks_On_Victim_Turn(t *testing.T) {
	p := stats("Player", 500, 0, 5, 0, 1, DmgSlashing)
	e := stats("Enemy",  500, 0, 0, 0, 0, DmgSlashing)
	pPerks := Perks{Poison: true}

	b := NewBattle(p, e, pPerks, Perks{}, rng())

	_ = hitBy(t, b, "Player") // наложили яд

	// первый ход жертвы — тика нет
	evs, _ := b.Step()
	for _, ev := range evs {
		if ev.Kind == EvPoisonTick {
			t.Fatalf("на первом ходу жертвы тика быть не должно")
		}
	}

	// второй ход жертвы — +1
	_ = hitBy(t, b, "Player")
	tick1 := poisonTickFor(t, b, "Enemy")
	if tick1.Dmg.Applied != 1 {
		t.Fatalf("на 2-м ходу жертвы яд должен тикать на +1, got %d", tick1.Dmg.Applied)
	}

	// третий ход жертвы — +2
	_ = hitBy(t, b, "Player")
	tick2 := poisonTickFor(t, b, "Enemy")
	if tick2.Dmg.Applied != 2 {
		t.Fatalf("на 3-м и далее ходах жертвы яд должен тикать на +2, got %d", tick2.Dmg.Applied)
	}
}
