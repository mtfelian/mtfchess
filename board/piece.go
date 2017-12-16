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

// Pairs converts offsets of piece to paies
func (offsets Offsets) Pairs(piece Piece) Pairs {
	pairs := make(Pairs, len(offsets))
	for i, o := range offsets {
		pairs[i] = Pair{X: o.X + piece.X(), Y: o.Y + piece.Y()}
	}
	return pairs
}

// Pairs is a slice of pairs
type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pairs) Less(i, j int) bool { return p[i].Y < p[j].Y || (p[i].Y == p[j].Y && p[i].X < p[j].X) }

// Pieces
type Pieces []Piece

// Piece
type Piece interface {
	fmt.Stringer
	// Name returns piece name
	Name() string
	// Attacks returns a slice of cells coords attacked by piece
	Attacks(b Board) Pairs
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

// PieceFilter
type PieceFilter struct {
	Names     []string
	X         []int
	Y         []int
	Colours   []Colour
	Condition func(Piece) bool
}
