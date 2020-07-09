package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Starting Battleship Game...\n")
	fmt.Print("Player name: ")
	player1, _ := reader.ReadString('\n')
	//fmt.Print("Player 2 name: ")
	//player2, _ := reader.ReadString('\n')

	p1 := NewPlayer(strings.TrimSpace(player1))
	//p2 := NewPlayer(strings.TrimSpace(player2))

	game := newGame(p1)

	game.run()
}
