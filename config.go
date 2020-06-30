package main

const (
	// GridSize size
	GridSize = 10
	// ShipHit ship hit
	ShipHit = iota
	// ShipPresent - grid has ship
	ShipPresent = iota
	// ShipCount - number of ships
	ShipCount = 5
	//Max square to cover by ships
	MaxSquareCount = 17
)

var (
	expectSize = []uint{5, 4, 3, 3, 2}
)
