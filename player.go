package main

import (
	"bufio"
	"fmt"
	"os"
)

// Player - player struct
type Player struct {
	name       string
	board      Board
	score      int
	enemyBoard Board
}

func NewPlayer(name string) Player {
	var board = newBoard()
	var enemyBoard = newBoard()
	var score = 0
	p := Player{name, board, score, enemyBoard}
	return p
}

func (p Player) fireRocket() (int, int) {
	// read the coordinates
	reader := bufio.NewReader(os.Stdin)
	var x, y int
	for {
		fmt.Print("\n" + p.name + " please enter rocket coordinates: ")
		first, _ := reader.ReadString('\n')
		x, y = convertStringCoordinatesToInt(first)

		validRocketCoordinates := validateRocketCoordinates(x, y)
		if !validRocketCoordinates {
			fmt.Println("Invalid rocket coordinates, please try again")
		} else {
			break
		}
	}

	// fire the rocket
	fmt.Printf("Player "+p.name+": Firing rocket at coordinates (%v, %v)", x, y)
	fmt.Println("")
	return x, y
}

func (p *Player) incrementScore() {
	p.score++
}

