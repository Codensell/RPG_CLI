package term

import (
	"fmt"

	gc "github.com/rthornton128/goncurses"
)

func centeredModal(std *gc.Window, h, w int, title string, lines []string) {
	maxY, maxX := std.MaxYX()
	if h > maxY { h = maxY }
	if w > maxX { w = maxX }
	y := (maxY - h) / 2
	x := (maxX - w) / 2

	win, err := gc.NewWindow(h, w, y, x)
	if err != nil { return }
	defer win.Delete()

	win.Box(0, 0)
	if title != "" {
		win.MovePrint(0, 2, " "+title+" ")
	}
	for i, s := range lines {
		if 1+i >= h-1 { break }
		win.MovePrint(1+i, 2, s)
	}
	win.MovePrint(h-2, 2, "Press any key...")
	win.Refresh()
	win.GetChar()
}

func ShowWinProgress(std *gc.Window, wins int) {
	lines := []string{
		fmt.Sprintf("Victory %d/5", wins),
		"HP fully restored.",
	}
	centeredModal(std, 7, 40, " Victory ", lines)
}

func ShowDefeat(std *gc.Window) {
	lines := []string{
		"You were defeated.",
		"New character will be created.",
	}
	centeredModal(std, 7, 44, " Defeat ", lines)
}

func ShowGameWon(std *gc.Window) {
	lines := []string{
		"You won 5 battles in a row!",
		"Game completed.",
	}
	centeredModal(std, 7, 44, " Congratulations ", lines)
}
