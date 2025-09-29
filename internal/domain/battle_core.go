package domain

import (
	"fmt"
	"math/rand"
	"strings"
)

type Battle struct {
	p, e actor
	attP bool //if true = player turn to attack, false = enemy turn
	rng  *rand.Rand
	init bool
}

func NewBattle(player Stats, enemy Stats, pPerks, ePerks Perks, rng *rand.Rand) *Battle {
	if rng == nil {
		rng = rand.New(rand.NewSource(1))
	}
	b := &Battle{
		p:    newActorFromStats(player, pPerks),
		e:    newActorFromStats(enemy, ePerks),
		rng:  rng,
		init: true,
	}
	if player.AGI >= enemy.AGI {
		b.attP = true
	} else {
		b.attP = false
	}
	return b
}

func (b *Battle) PlayerStats() Stats {
	return b.p.toStats()
}
func (b *Battle) EnemyStats() Stats {
	return b.e.toStats()
}

func (b *Battle) Step() (events []Event, done bool) {
	if b.init {
		b.init = false
		events = append(events,
			Event{Kind: EvStart, Note: "Battle Begin"},
			Event{Kind: EvInitiative, Note: fmt.Sprintf("Initiative - %s", b.attName())},
		)
	}

	att, def := b.attRef()

	events = append(events, Event{
		Kind: EvTurnBegin, Who: att.name,
		Note: fmt.Sprintf("Turn %d - %s", att.turnCount+1, att.name),
	})

	if att.poisoned {
		att.poisonTurns++
		tick := 0
		if att.poisonTurns == 2 {
			tick = 1
		} else if att.poisonTurns >= 3 {
			tick = 2
		}
		if tick > 0 {
			redShield := 0
			if att.perks.Shield && att.str > def.str {
				redShield = 3
			}
			redStone := 0
			if att.perks.StoneSkin {
				redStone = att.sta
			}
			appliedTick := tick - redShield - redStone
			if appliedTick < 0 {
				appliedTick = 0
			}
			att.hp -= appliedTick

			events = append(events, Event{
				Kind: EvPoisonTick,
				Who:  att.name,
				Note: fmt.Sprintf("Poison tick: %d", appliedTick),
				Dmg: &DamageBreakdown{
					Weapon:             0,
					STR:                0,
					Bonus:              tick,
					WeaponAfterPartial: 0,
					Shield:             redShield,
					StoneSkin:          redStone,
					Applied:            appliedTick,
				},
			})
			if att.hp <= 0 {
				events = append(events, Event{
					Kind:   EvVictory,
					Who:    def.name,
					Target: att.name,
					Note:   fmt.Sprintf("Victory by poison: %s", def.name),
				})
				return events, true
			}
		}
	}

	totalAGI := att.agi + def.agi
	if totalAGI < 1 {
		totalAGI = 1
	}
	roll := 1 + b.rng.Intn(totalAGI)
	events = append(events, Event{
		Kind:   EvAttackRoll,
		Who:    att.name,
		Target: def.name,
		Roll:   roll,
		Note:   fmt.Sprintf("Attack roll %d vs AGI_def=%d", roll, def.agi),
	})
	if roll <= def.agi {
		events = append(events, Event{
			Kind:   EvMiss,
			Who:    att.name,
			Target: def.name,
			Note:   "Miss",
		})
		b.afterAttack()
		return events, false
	}

	weapon := att.weaponDMG
	str := att.str

	bonus, tags := b.attackBonus(att, def)

	weaponAfterPartial := weapon
	if def.perks.ImmuneSlashingWeapon && att.weaponType == DmgSlashing {
		weaponAfterPartial = 0
	}
	if def.perks.VulnBludge && att.weaponType == DmgBludge {
		weaponAfterPartial = weaponAfterPartial * 2
	}

	redShield := 0
	if def.perks.Shield && def.str > att.str {
		redShield = 3
	}
	redStone := 0
	if def.perks.StoneSkin {
		redStone = def.sta
	}

	applied := weaponAfterPartial + str + bonus - redShield - redStone
	if applied < 0 {
		applied = 0
	}
	def.hp -= applied

	noteParts := []string{fmt.Sprintf("Hit for %d", applied)}
	if len(tags) > 0 {
		noteParts = append(noteParts, "+["+strings.Join(tags, ",")+"]")
	}
	defTags := []string{}
	if redShield > 0 {
		defTags = append(defTags, fmt.Sprintf("Shield:-%d", redShield))
	}
	if redStone > 0 {
		defTags = append(defTags, fmt.Sprintf("Stone:-%d", redStone))
	}
	if len(defTags) > 0 {
		noteParts = append(noteParts, "-["+strings.Join(defTags, ",")+"]")
	}
	note := strings.Join(noteParts, " ")

	events = append(events, Event{
		Kind:   EvHit,
		Who:    att.name,
		Target: def.name,
		Note:   note,
		Dmg: &DamageBreakdown{
			Weapon:             weapon,
			STR:                str,
			Bonus:              bonus,
			WeaponAfterPartial: weaponAfterPartial,
			Shield:             redShield,
			StoneSkin:          redStone,
			Applied:            applied,
		},
	})

	if att.perks.Poison && applied > 0 {
		def.poisoned = true
	}

	if def.hp <= 0 {
		events = append(events, Event{
			Kind:   EvVictory,
			Who:    att.name,
			Target: def.name,
			Note:   fmt.Sprintf("Victory: %s", att.name),
		})
		return events, true
	}

	b.afterAttack()
	return events, false
}


func (b *Battle) attRef() (att *actor, def *actor) {
	if b.attP {
		return &b.p, &b.e
	}
	return &b.e, &b.p
}

func (b *Battle) attName() string {
	if b.attP {
		return b.p.name
	}
	return b.e.name
}

func (b *Battle) afterAttack() {
	if b.attP {
		b.p.turnCount++
	} else {
		b.e.turnCount++
	}
	b.attP = !b.attP
}

func (b *Battle) attackBonus(att, def *actor) (int, []string) {
	bonus := 0
	var tags []string

	if att.perks.ActionSurge && att.turnCount == 0 {
		bonus += att.weaponDMG
		tags = append(tags, fmt.Sprintf("AS:%d", att.weaponDMG))
	}

	if att.perks.Rage {
		if att.turnCount < 3 {
			bonus += 2
			tags = append(tags, "Rage:+2")
		} else {
			bonus -= 1
			tags = append(tags, "Rage:-1")
		}
	}

	if att.perks.SneakAttack && att.agi > def.agi {
		bonus += 1
		tags = append(tags, "Sneak:+1")
	}

	if att.perks.DragonBreath && (att.turnCount+1)%3 == 0 {
		bonus += 3
		tags = append(tags, "Breath:+3")
	}

	return bonus, tags
}
