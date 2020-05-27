package main

// Ship struct
type Ship struct {
	X    int
	Y    int
	Size int
	Hits int
}

// check if Ship is sunk
func (s *Ship) isSunk() bool {
	return s.Hits >= s.Size
}

// Player struct
type Player struct {
	ID       string
	Shots    []int
	Ships    []Ship
	ShipGrid []int
}

// init player
func (p *Player) init(id string) {

}

func (p *Player) shoot(index int) (hit bool) {
	return
}

func (p *Player) createShips() {

}

// Game struct
type Game struct {
	ID     int
	Status int
}

func (g *Game) init(id int) {

}

func (g *Game) getState() {

}
