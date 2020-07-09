package main

import "fmt"

func printPlayerBoard(player Player) {
	fmt.Printf("%v's board: \n", player.name)
	fmt.Println(player.board.printBoard())
}

func validateRocketCoordinates(x int, y int) bool {
	if x < 0 || x > 9 || y < 0 || y > 9 {
		return false
	}
	return true
}

func checkIfHitOrMiss(board Board, x int, y int){

}