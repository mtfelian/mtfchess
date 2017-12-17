package base

import "fmt"

// Pieces
type Pieces []Piece

// Piece
type Piece interface {
	fmt.Stringer
	// Name returns piece name
	Name() string
	// Attacks returns a slice of cells coords attacked by piece
	Attacks(b Board) Coords
	// Destinations returns a slice of cells coords to destination cells of possible moves
	Destinations(b Board) Coords
	// Colour returns piece colour
	Colour() Colour
	// SetCoords sets the figure coords to
	SetCoords(to Coord)
	// Coord returns piece coords
	Coord() Coord
	// Copy returns a deep copy of a piece
	Copy() Piece
	// Project a piece to coords, returns a pointer to a new copy of a board, don't check legality
	// this don't change coords of a piece
	Project(to Coord, b Board) Board
}
