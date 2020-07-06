package main

import "bytes"

// Items on the board
const (
	WATER = iota
	MISS
	HIT
	SHIP
)

// Board - game board
type Board [][]int

// Player - player struct
type Player struct{} // TODO - interface with events?

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

// set grid to ShipPresent
func (b Board) addShip(s *Ship) {
	if s.First.X == s.Last.X {
		for i := s.First.Y; i <= s.Last.Y; i++ {
			b[s.First.X][i] = ShipPresent
		}
	} else {
		for i := s.First.X; i <= s.Last.X; i++ {
			b[s.First.Y][i] = ShipPresent
		}
	}
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

func (b Board) printBoard() string {
	var buf bytes.Buffer

	for _, r := range b {
		for _, c := range r {
			switch c {
			case WATER:
				buf.WriteRune('W')
			case SHIP:
				buf.WriteRune('S')
			}
			buf.WriteRune(' ')
		}
		buf.WriteRune('\n')
	}

	return buf.String()
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

//TODO: read ships and print board by Georgiana
//TODO: shot the canon  by Marian
func (o *Observer) run() {
	o.p1.setupShips()
	o.p2.setupShips()

	// run the game
	var cannonBall Grid
	var response int
	var sunk *Ship

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
	}
}
