package main

// Board - game board
type Board [][]int

// Player - player struct
type Player struct{}

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

	} else if g.Y == s.First.Y {

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

	return ""
}

func validateShips(ships []*Ship) (valid bool) {
	if len(ships) != ShipCount {
		return
	}

	board := newBoard()

	for i, ship := range ships {
		if ship.Size != expectSize[i] {
			return
		}

		if ship.First.X > ship.Last.X || ship.First.Y > ship.Last.Y {
			return
		}

		board.addShip(ship)
	}

	return valid
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

func (o *Observer) run() {
	o.p1.setupShips()
	o.p2.setupShips()
}
