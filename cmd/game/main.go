package main

import (
	"math/rand"
	"time"

	"github.com/Codensell/RPG_CLI/internal/domain"
	"github.com/Codensell/RPG_CLI/internal/ui/term"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p := domain.NewCharacter("Player")
	_ = p.ApplyClassLevel(domain.ClassWarrior)
	term.DrawFrames(p.CharacterData())
}
