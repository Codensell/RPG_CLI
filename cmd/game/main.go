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
	
	std, err := gc.Init()
	if err != nil {
		panic(err)
	}
	defer gc.End()
	gc.Echo(false)
	gc.Cursor(0)
	gc.Raw(true)

	p := domain.NewCharacter("Player")

	chosen := term.SelectClass(std, p.CharacterData())

	_ = p.ApplyClassLevel(chosen)

	e := domain.NewRandomEnemy()
	term.DrawFrames(std, p.CharacterData(), e.CharacterData())
}
