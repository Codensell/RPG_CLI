package term

import gc "github.com/rthornton128/goncurses"

func DrawFrames() {
	std, err := gc.Init()
	if err != nil {
		panic(err)
	}
	defer gc.End()

	gc.Echo(false)
	gc.Cursor(0)
	gc.Raw(true)

	std.Clear()
	std.Refresh()

	maxY, maxX := std.MaxYX()
	if maxY < 12 {
		maxY = 12
	}
	playerH := 5
	enemyH := 5
	logH := maxY - playerH - enemyH
	if logH < 3 {
		logH = 3
	}

	playerW, _ := gc.NewWindow(playerH, maxX, 0, 0)
	logW, _ := gc.NewWindow(logH, maxX, playerH, 0)
	enemyW, _ := gc.NewWindow(enemyH, maxX, playerH+logH, 0)

	playerW.Erase()
	logW.Erase()
	enemyW.Erase()

	playerW.Box(0, 0)
	playerW.MovePrint(0, 2, " Player ")
	playerW.MovePrint(2, 2, "[test] upper frame seen?")
	playerW.Refresh()

	logW.Box(0, 0)
	logW.MovePrint(0, 2, " Battle Logs ")
	logW.MovePrint(1, 2, "[test] mid frame seen?")
	logW.Refresh()

	enemyW.Box(0, 0)
	enemyW.MovePrint(0, 2, " Enemy ")
	enemyW.MovePrint(enemyH-2, 2, "Press any button to quit")
	enemyW.Refresh()

	std.GetChar()
}