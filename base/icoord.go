package base

import "fmt"

// Coord
type Coord interface {
	fmt.Stringer
	// Add returns a sum of coords
	Add(c Coord) Coord
	// Out returns true if coords is out of board
	Out(b IBoard) bool
	// Equals returns true if coords are equal
	Equals(to Coord) bool
	// Copy returns a copy of coord
	Copy() Coord
}
