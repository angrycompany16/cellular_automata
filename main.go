package main

import (
	"cellular_automata/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetFullscreen(true)

	g.Init()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
