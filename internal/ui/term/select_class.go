package term

import (
	"github.com/Codensell/RPG_CLI/internal/domain"
	gc "github.com/rthornton128/goncurses"
)

func SelectClass(std *gc.Window, s domain.Stats) domain.Class {
	std.Clear()
	std.Refresh()

	maxY, maxX := std.MaxYX()
	h, w := 12, 48
	if maxY < h || maxX < w {
		std.MovePrint(1, 2, "Terminal is too small for class selection")
		std.Refresh()
		std.GetChar()
		return domain.ClassWarrior
	}

	y := (maxY - h) / 2
	x := (maxX - w) / 2
	win, _ := gc.NewWindow(h, w, y, x)
	defer win.Delete()

	win.Box(0, 0)
	win.MovePrint(0, 2, " Class selection ")

	win.MovePrintf(2, 2, "Name:   %s", s.Name)
	win.MovePrintf(3, 2, "HP:     %d/%d", s.HP, s.MaxHP)
	win.MovePrintf(4, 2, "STR:    %d", s.STR)
	win.MovePrintf(5, 2, "AGI:    %d", s.AGI)
	win.MovePrintf(6, 2, "STA:    %d", s.STA)

	win.MovePrint(8, 2,  "[1] Warrior   (+5 HP/level, Sword)")
	win.MovePrint(9, 2,  "[2] Barbarian (+6 HP/level, Club)")
	win.MovePrint(10, 2, "[3] Rogue     (+4 HP/level, Dagger)")
	win.Refresh()

	for {
		ch := win.GetChar()
		switch ch {
		case '1':
			return domain.ClassWarrior
		case '2':
			return domain.ClassBarbarian
		case '3':
			return domain.ClassRogue
		}
	}
}
