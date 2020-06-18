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
)

var (
	expectSize = []uint{5, 4, 3, 3, 2}
)
