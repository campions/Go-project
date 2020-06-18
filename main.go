package main

import (
	"fmt"
)

func main() {
	fmt.Print("Starting Battleship Game...\n")
	p1 := Player{}
	p2 := Player{}
	game := newGame(p1, p2)
	game.run()
}
