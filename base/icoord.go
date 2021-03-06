package base

import "fmt"

// ICoord
type ICoord interface {
	fmt.Stringer
	// Add returns a sum of coords
	Add(c ICoord) ICoord
	// OutOf returns true if coords is out of board
	OutOf(b IBoard) bool
	// Equals returns true if coords are equal
	Equals(to ICoord) bool
	// Copy returns a copy of coord
	Copy() ICoord
}
