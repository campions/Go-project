package main

import (
	"bufio"
	"fmt"
	"os"
)

// Player - player struct
type Player struct {
	name  string
	board Board
	score int
}

func NewPlayer(name string) Player {
	var board = newBoard()
	var score = 0
	p := Player{name, board, score}
	return p
}

func (p Player) fireRocket() {
	reader := bufio.NewReader(os.Stdin)
	var x, y int
	for {
		fmt.Print("Enter rocket coordinates: ")
		first, _ := reader.ReadString('\n')
		x, y = convertStringCoordinatesToInt(first)

		validRocketCoordinates := validateRocketCoordinates(x, y)
		if !validRocketCoordinates {
			fmt.Println("Invalid rocket coordinates, please try again")
		} else {
			break
		}
	}

	fmt.Printf("Firing rocket at coordinates (%v, %v)", x, y)
	fmt.Println("")

}
