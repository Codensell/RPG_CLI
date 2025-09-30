package term

import (
	"fmt"

	gc "github.com/rthornton128/goncurses"
	"github.com/Codensell/RPG_CLI/internal/domain"
)

func OfferWeaponSwap(std *gc.Window, current, drop domain.Weapon) bool {
	if drop.Name == "" { return false }

	maxY, maxX := std.MaxYX()
	h, w := 11, 56
	if h > maxY { h = maxY }
	if w > maxX { w = maxX }
	y := (maxY - h) / 2
	x := (maxX - w) / 2

	win, err := gc.NewWindow(h, w, y, x)
	if err != nil { return false }
	defer win.Delete()

	win.Box(0, 0)
	win.MovePrint(0, 2, " Loot: Replace weapon? ")

	win.MovePrint(2, 2, "Current:")
	win.MovePrint(3, 4, fmt.Sprintf("%s  (DMG %d, %s)", current.Name, current.Base, current.Type))

	win.MovePrint(5, 2, "Drop:")
	win.MovePrint(6, 4, fmt.Sprintf("%s  (DMG %d, %s)", drop.Name, drop.Base, drop.Type))

	win.MovePrint(8, 2, "[E] Equip drop   [K] Keep current   [ESC] Cancel")
	win.Refresh()

	for {
		switch ch := win.GetChar(); ch {
		case 'e', 'E':
			return true
		case 'k', 'K', 27:
			return false
		}
	}
}
