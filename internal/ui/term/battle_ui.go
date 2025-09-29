package term

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Codensell/RPG_CLI/internal/domain"
	gc "github.com/rthornton128/goncurses"
)

func RunBattleFrames(std *gc.Window, player domain.Stats, enemy domain.Stats, pPerks, ePerks domain.Perks) bool {
	const frameH = 10
	const logH = frameH * 2

	std.Clear()
	std.Refresh()

	maxY, maxX := std.MaxYX()
	totalH := frameH + logH + frameH
	if totalH > maxY {
		std.MovePrint(1, 2, "Terminal is too low")
		std.MovePrint(2, 2, "Please, increase the terminal height")
		std.Refresh()
		std.GetChar()
		return false
	}

	playerY := 0
	logY := frameH
	enemyY := frameH + logH

	playerW, _ := gc.NewWindow(frameH, maxX, playerY, 0)
	defer playerW.Delete()
	logW, _ := gc.NewWindow(logH, maxX, logY, 0)
	defer logW.Delete()
	enemyW, _ := gc.NewWindow(frameH, maxX, enemyY, 0)
	defer enemyW.Delete()

	playerW.Erase(); logW.Erase(); enemyW.Erase()
	playerW.Box(0, 0); playerW.MovePrint(0, 2, " Player "); playerW.Refresh()
	logW.Box(0, 0);    logW.MovePrint(0, 2, " Battle Logs "); logW.Refresh()
	enemyW.Box(0, 0);  enemyW.MovePrint(0, 2, " Enemy "); enemyW.Refresh()

	DrawPlayerFrame(playerW, player)
	DrawEnemyFrame(enemyW, enemy)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := domain.NewBattle(player, enemy, pPerks, ePerks, rng)

	var lines []string
	renderLog := func(w *gc.Window, add ...string) {
		lines = append(lines, add...)
		contentH := logH - 2
		start := 0
		if len(lines) > contentH {
			start = len(lines) - contentH
		}

		w.Erase()
		for i, s := range lines[start:] {
			w.MovePrint(1+i, 2, s)
		}
		w.Box(0, 0)
		w.MovePrint(0, 2, " Battle Logs ")
		w.Refresh()
	}

	evs, done := b.Step()
	for _, ev := range evs {
		renderLog(logW, formatEvent(ev))
	}
	DrawPlayerFrame(playerW, b.PlayerStats())
	DrawEnemyFrame(enemyW, b.EnemyStats())

	if done {
		renderLog(logW, "Battle End", "Press any key...")
		logW.Refresh()
		std.GetChar()
		playerWon := b.EnemyStats().HP <= 0
		return playerWon
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		evs, done = b.Step()
		for _, ev := range evs {
			renderLog(logW, formatEvent(ev))
		}
		DrawPlayerFrame(playerW, b.PlayerStats())
		DrawEnemyFrame(enemyW, b.EnemyStats())

		if done {
			renderLog(logW, "Battle End", "Press any key...")
			logW.Refresh()
			std.GetChar()
			playerWon := b.EnemyStats().HP <= 0
			return playerWon
		}
	}
	return b.EnemyStats().HP <= 0
}

func formatEvent(ev domain.Event) string {
	switch ev.Kind {
	case domain.EvStart:
		return "Battle Begin"
	case domain.EvInitiative:
		return ev.Note
	case domain.EvTurnBegin:
		return ev.Note
	case domain.EvAttackRoll:
		return fmt.Sprintf("%s attacks %s: roll=%d", ev.Who, ev.Target, ev.Roll)
	case domain.EvMiss:
		return fmt.Sprintf("%s missed %s", ev.Who, ev.Target)
	case domain.EvHit:
		if ev.Dmg != nil {
			return fmt.Sprintf("%s hits %s for %d (W: %d - %d, STR: %d, Bonus: %d, Shield:%d, Stone: %d)",
				ev.Who, ev.Target, ev.Dmg.Applied,
				ev.Dmg.Weapon, ev.Dmg.WeaponAfterPartial,
				ev.Dmg.STR, ev.Dmg.Bonus, ev.Dmg.Shield, ev.Dmg.StoneSkin,
			)
		}
		return fmt.Sprintf("%s hits %s", ev.Who, ev.Target)
	case domain.EvPoisonTick:
		return fmt.Sprintf("%s suffers poison: %s", ev.Who, ev.Note)
	case domain.EvVictory:
		return fmt.Sprintf("Victory: %s defeats %s", ev.Who, ev.Target)
	default:
		return ev.Note
	}
}
