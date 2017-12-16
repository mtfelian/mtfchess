package board

import (
	"fmt"
)

// Pair is a coordinate pair
type Pair struct {
	X, Y int
}

// Offsets is a slice of pair offsets
type Offsets []Pair

// Pairs is a slice of pairs
type Pairs []Pair

// Piece
type Piece interface {
	fmt.Stringer
	// Name returns piece name
	Name() string
	// Offsets returns a slice of offsets to possible moves
	Offsets(b Board) Offsets
	// Colour returns piece colour
	Colour() Colour
	// SetCoords sets the figure coords to (x,y)
	SetCoords(x, y int)
	X() int
	Y() int
	// Copy returns a deep copy of a piece
	Copy() Piece
	// Project a piece to coords (x,y), returns a pointer to a new copy of a board, don't check legality
	// this don't change coords of a piece
	Project(x, y int, b Board) Board
}
