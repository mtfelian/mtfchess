package base

import (
	"fmt"

	. "github.com/mtfelian/mtfchess/colour"
)

// Pieces
type Pieces []IPiece

// IPiece
type IPiece interface {
	fmt.Stringer
	// Name returns piece name
	Name() string
	// Attacks returns a slice of cells coords attacked by piece
	Attacks(b IBoard) ICoords
	// Destinations returns a slice of cells coords to destination cells of possible moves
	Destinations(b IBoard) ICoords
	// Colour returns piece colour
	Colour() Colour
	// SetCoords sets the figure coords to
	SetCoords(board IBoard, to ICoord)
	// Coord returns piece coords
	Coord() ICoord
	// Copy returns a deep copy of a piece
	Copy() IPiece
}
