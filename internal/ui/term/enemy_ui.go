package term

import (
	gc "github.com/rthornton128/goncurses"
	"github.com/Codensell/RPG_CLI/internal/domain"
)

func DrawEnemyFrame(w *gc.Window, s domain.Stats) {
	w.Erase()
	w.Box(0, 0)
	w.MovePrint(0, 2, " Enemy ")
	w.MovePrint(1, 2, "Name: "+s.Name)
	w.MovePrint(2, 2, "HP:  ")
	w.MovePrint(2, 7, s.HP, "/", s.MaxHP)
	w.MovePrint(3, 2, "STR/AGI/STA: ")
	w.MovePrint(3, 18, s.STR, "/", s.AGI, "/", s.STA)
	w.MovePrint(4, 2, "Weapon: ")
	w.MovePrint(4, 11, s.Weapon, "(", s.WeaponDMG, ", ", s.WeaponTYPE, ")")
	w.Refresh()
}
