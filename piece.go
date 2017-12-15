package mtfchess

import "fmt"

// pair is a coordinates pair
type pair struct {
	x int
	y int
}

// offsets is a slice of pair offsets
type offsets []pair

// Piece
type Piece interface {
	fmt.Stringer
	// Board returns pointer to a Board
	Board() *Board
	// Name returns piece name
	Name() string
	// Shift returns a slice of pairs with x, y offsets
	Shift(xy pair) offsets
	// CanJump should return true if piece don't know barriers like a chess knight
	CanJump() bool
}
