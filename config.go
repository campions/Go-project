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
	ships []int = []int{3, 2, 1}
)
