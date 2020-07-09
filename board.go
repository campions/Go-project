package main

import "bytes"

// Board - game board
type Board [][]int

// set grid to ShipPresent
func (b Board) addShip(s *Ship) (added bool) {
	if validateShipCoordinates(s) {
		if s.First.X == s.Last.X {
			//check if it`s any ship on top
			if s.First.Y != 0 && b[s.First.Y-1][s.First.X] == ShipPresent {
				return false
			}
			//check if it's any ship on bottom
			if s.Last.Y != 9 && b[s.Last.Y+1][s.Last.X] == ShipPresent {
				return false
			}

			for i := s.First.Y; i <= s.Last.Y; i++ {
				if s.First.X != 0 && b[i][s.First.X-1] == ShipPresent {
					return false
				}
				if s.First.X != 9 && b[i][s.First.X+1] == ShipPresent {
					return false
				}
				if b[i][s.First.X] == ShipPresent {
					return false
				}
				b[i][s.First.X] = ShipPresent

			}
			return true
		} else {
			//LEFT
			if s.First.X != 0 && b[s.First.Y][s.First.X-1] == ShipPresent {
				return false
			}
			//RIGHT
			if s.Last.X != 9 && b[s.Last.Y][s.Last.X+1] == ShipPresent {
				return false
			}

			for i := s.First.X; i <= s.Last.X; i++ {
				if s.First.Y != 0 && b[s.First.Y-1][i] == ShipPresent {
					return false
				}
				if s.First.Y != 9 && b[s.First.Y+1][i] == ShipPresent {
					return false
				}
				if b[s.First.Y][i] == ShipPresent {
					return false
				}
				b[s.First.Y][i] = ShipPresent
			}
			return true

		}
	} else {
		return false
	}
}

func (b Board) printBoard() string {
	var buf bytes.Buffer

	for _, r := range b {
		for _, c := range r {
			switch c {
			case WATER:
				buf.WriteRune('~')
			case SHIP:
				buf.WriteRune('■')
			}
			buf.WriteRune(' ')
		}
		buf.WriteRune('\n')
	}

	return buf.String()
}
