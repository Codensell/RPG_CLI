package main

import (
	"math/rand"
	"time"

	"github.com/Codensell/RPG_CLI/internal/domain"
	"github.com/Codensell/RPG_CLI/internal/ui/term"
	gc "github.com/rthornton128/goncurses"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// ncurses: один раз за всю жизнь процесса
	std, err := gc.Init()
	if err != nil {
		panic(err)
	}
	defer gc.End()
	gc.Echo(false)
	gc.Cursor(0)
	gc.Raw(true)

	// 1) Создаём персонажа
	p := domain.NewCharacter("Player")

	// 2) Выбор класса (через существующий std)
	chosen := term.SelectClass(std, p.CharacterData())

	// 3) Применяем выбранный класс
	_ = p.ApplyClassLevel(chosen)

	// 4) Основной экран (через тот же std)
	term.DrawFrames(std, p.CharacterData())
}
