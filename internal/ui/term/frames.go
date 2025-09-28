package term

import (
	"github.com/Codensell/RPG_CLI/internal/domain"
	gc "github.com/rthornton128/goncurses"
)

func DrawFrames(std *gc.Window, s domain.Stats) {
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
		return
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

	playerW.Box(0, 0)
	playerW.MovePrint(0, 2, " Player ")
	DrawPlayerFrame(playerW, s)
	playerW.Refresh()

	logW.Box(0, 0)
	logW.MovePrint(0, 2, " Battle Logs ")
	logW.Refresh()

	enemyW.Box(0, 0)
	enemyW.MovePrint(0, 2, " Enemy ")
	enemyW.MovePrint(frameH-2, 2, "Press any key to quit")
	enemyW.Refresh()

	std.GetChar()
}
