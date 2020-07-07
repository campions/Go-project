package main

// Player - player struct
type Player struct {
	name  string
	board Board
}

func NewPlayer(name string) Player {
	var board = newBoard()
	p := Player{name, board}
	return p
}
