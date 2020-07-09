package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


// Grid struct
type Grid struct {
	X, Y int
}

// Ship struct
type Ship struct {
	First Grid
	Last  Grid
	Size  uint
	Hits  []int
	Sunk  bool
}

// Game - game struct
type Game struct {
	Ships  []*Ship
	Player Player
	State  GameState
}

// GameState - state struct
type GameState struct {
	PlayerBoard Board
	EnemyBoard  Board
}

// Observer - game observer
type Observer struct {
	p1 *Game
	p2 *Game
}

// board setup
func newBoard() (b Board) {
	b = make([][]int, GridSize)
	for i := range b {
		b[i] = make([]int, GridSize)
	}
	return
}

// check point is a ship
func (s *Ship) contains(g Grid) bool {
	if g.X == s.First.X {
		if g.Y >= s.First.Y && g.Y <= s.Last.Y {
			return true
		}
	} else if g.Y == s.First.Y {
		if g.X >= s.First.X && g.X <= s.Last.X {
			return true
		}
	}

	return false
}

// hit a given point
func (s *Ship) hit(g Grid) (hit, snk bool) {
	// hit inside ship ?
	if !s.contains(g) {
		return false, false
	}

	// register hit
	if g.Y == s.First.Y {
		s.Hits[g.X-s.First.X] = ShipHit
	} else {
		s.Hits[g.Y-s.First.Y] = ShipHit
	}

	// check ship status
	if uint(len(s.Hits)) >= s.Size {
		s.Sunk = true
		return true, true
	}

	// hit, still alive
	return true, false
}

func validateShips(ships []*Ship) bool {
	if len(ships) != ShipCount {
		return false
	}

	board := newBoard()

	for i, ship := range ships {
		if ship.Size != expectSize[i] {
			return false
		}

		if ship.First.X > ship.Last.X || ship.First.Y > ship.Last.Y {
			return false
		}

		if ship.First.X < 0 || ship.First.Y < 0 {
			return false
		}
		if ship.Last.X >= 10 || ship.Last.Y >= 10 {
			return false
		}

		board.addShip(ship)
	}

	var squareCount uint
	for _, column := range board {
		for _, row := range column {
			if row == SHIP {
				squareCount++
			}
		}
	}
	if squareCount != MaxSquareCount {
		return false
	}

	return true
}

func validateShipCoordinates(ship *Ship) bool {

	if ship.First.X > ship.Last.X || ship.First.Y > ship.Last.Y {
		return false
	}

	if ship.First.X < 0 || ship.First.Y < 0 {
		return false
	}
	if ship.Last.X >= 10 || ship.Last.Y >= 10 {
		return false
	}
	if ship.First.X != ship.Last.X && ship.First.Y != ship.Last.Y {
		return false
	}
	if ship.First.X == ship.Last.X && (uint(ship.Last.Y - ship.First.Y) != ship.Size-1) {
		return false
	}
	if ship.First.Y == ship.Last.Y && (uint(ship.Last.X - ship.First.X) != ship.Size-1) {
		return false
	}

	return true
}

// setup game ships
func (g *Game) setupShips() {
	var ships []*Ship
	for _, size := range expectSize {
		ships = append(ships, &Ship{
			Size: size,
			Hits: make([]int, size),
		})
	}

	for _, ship := range ships {
		g.State.PlayerBoard.addShip(ship)
	}
	g.Ships = ships
}

func newGame(p1, p2 Player) *Observer {
	return &Observer{
		p1: &Game{
			Player: p1,
			State: GameState{
				PlayerBoard: newBoard(),
				EnemyBoard:  newBoard(),
			},
		},
		p2: &Game{
			Player: p2,
			State: GameState{
				PlayerBoard: newBoard(),
				EnemyBoard:  newBoard(),
			},
		},
	}
}
func getCell(first string) (int, int) {
	first = strings.TrimSpace(first)
	param := strings.Fields(first)
	if len(param) != 2 {
		return -1, -1
	}
	x, err := strconv.Atoi(param[0])
	y, err1 := strconv.Atoi(param[1])
	if err != nil || err1 != nil {
		return -1, -1
	}
	return x, y
}

func readShips(p Player) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Let`s add the ships for ", p.name)
	no := 0
	for j := 4; j >= 1; j-- {
		//we start with ship with 4 cells
		for k := 5 - j; k >= 1; k-- {
			no++
			if j != 1 {
				fmt.Print("Ship number: ",no, " with ", j, " cells, please add first cell: ")
				first, _ := reader.ReadString('\n')
				a, b := getCell(first)

				if a < 0 || b < 0 {
					log.Fatal("invalid ship")
				}

				fmt.Print("Ship number: ",no, " with ", j, " cells, please add last cell: ")
				last, _ := reader.ReadString('\n')
				c, d := getCell(last)
				if c < 0 || d < 0 {
					log.Fatal("invalid ship")
				}
				isAdded := p.board.addShip(&Ship{Grid{a, b}, Grid{c, d}, uint(j), nil, false})
				if !isAdded {
					log.Fatal("invalid ship")
				}
			} else {
				fmt.Print("Ship number: ",no, " with ", j, " cell, please add the cell: ")
				first, _ := reader.ReadString('\n')
				x, y := getCell(first)
				if x < 0 || y < 0 {
					log.Fatal("invalid ship")
				}
				isAdded := p.board.addShip(&Ship{Grid{x, y}, Grid{x, y}, uint(j), nil, false})
				if !isAdded {
					log.Fatal("invalid ship")
				}
			}
		}
	}
}

//TODO: read ships and print board by Georgiana
//TODO: shot the canon  by Marian
func (o *Observer) run() {
	//o.p1.setupShips()
	//o.p2.setupShips()
	readShips(o.p1.Player)
	//readShips(o.p2.Player)

	printPlayerBoard(o.p1.Player)
	//printPlayerBoard(o.p2.Player)

	/*
	// run the game
	var cannonBall Grid
	var response int
	var sunk *Ship
	// read user position - read coordinate
	// shoot the cannot
	// handle the cannot
	// handle the response
	for {

		cannonBall = o.p1.ShotTheCannon()
		response, sunk = o.p2.HandleTheCannonHit(cannonBall)
		o.p1.HandleTheResponse(cannonBall, response)
		if sunk != nil {
			o.p1.CheckTheShippedSunk(sunk)
		}

		o.OnChange(o.p1.State, o.p2.State)

		if o.p2.Lost() {
			o.p1.Win()
			//print player 1 is the winner
		}

		cannonBall = o.p2.ShotTheCannon()
		response, sunk = o.p1.HandleTheCannonHit(cannonBall)
		o.p2.HandleTheResponse(cannonBall, response)
		if sunk != nil {
			o.p2.CheckTheShippedSunk(sunk)
		}

		o.OnChange(o.p1.State, o.p2.State)

		if o.p1.Lost() {
			o.p2.Win()
			//print player 2 is the winner
		}
	}*/
}
