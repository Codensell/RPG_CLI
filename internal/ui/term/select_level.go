package term

import (
	"github.com/Codensell/RPG_CLI/internal/domain"
	gc "github.com/rthornton128/goncurses"
)

func SelectLevel(std *gc.Window, totalLevels int) (domain.Class, bool) {
	std.Clear()
	std.Refresh()
	maxY, maxX := std.MaxYX()
	h, w := 9, 50
	y := (maxY - h) / 2
	x := (maxX - w) / 2

	win, err := gc.NewWindow(h, w, y, x)
	if err != nil {
		return domain.ClassWarrior, false
	}
	defer win.Delete()
	win.Box(0, 0)
	win.MovePrint(0, 2, " Level Up ")

	if totalLevels >= 3 {
		win.MovePrint(2, 2, "Max total level reached (3). Press any key.")
		win.Refresh()
		win.GetChar()
		return domain.ClassWarrior, false
	}

	win.MovePrint(2, 2, "Choose a class to gain a level (total < 3):")
	win.MovePrint(4, 2, "[1] Warrior   (+HP5/level, Action Surge -> Shield -> STR+1)")
	win.MovePrint(5, 2, "[2] Barbarian (+HP6/level, Rage -> Stone Skin -> STA+1)")
	win.MovePrint(6, 2, "[3] Rogue     (+HP4/level, Sneak -> (AGI+1) -> Poison)")
	win.MovePrint(7, 2, "[S] Skip")
	win.Refresh()

	for {
		ch := win.GetChar()
		switch ch {
		case '1':
			return domain.ClassWarrior, true
		case '2':
			return domain.ClassBarbarian, true
		case '3':
			return domain.ClassRogue, true
		case 's', 'S', 27:
			return domain.ClassWarrior, false
		}
	}
}
