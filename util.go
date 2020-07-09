package main

import "fmt"

func printPlayerBoard(player Player) {
	fmt.Printf("Player %v board: \n", player.name)
	fmt.Println(player.board.printBoard())
}
