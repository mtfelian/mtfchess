package mtfchess

import "fmt"

// Pair is a coordinate pair
type Pair struct {
	X, Y int
}

// Offsets is a slice of pair offsets
type Offsets []Pair

// Piece
type Piece interface {
	fmt.Stringer
	// Board returns pointer to a Board
	Board() *Board
	// Name returns piece name
	Name() string
	// Offsets returns a slice of offsets to possible moves
	Offsets() Offsets
	// CanJump should return true if piece don't know barriers like a chess knight
	CanJump() bool
	// Colour returns piece colour
	Colour() Colour
	// SetCoords sets the figure coords
	SetCoords(x, y int)
}
