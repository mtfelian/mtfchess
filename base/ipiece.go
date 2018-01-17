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
	// Capital returns piece's capital letter
	Capital() rune
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
	// Set piece to p1
	Set(p1 IPiece)

	// Promote returns a promoted piece
	Promote() IPiece
	// SetPromote sets a piece promote to
	SetPromote(to IPiece)
	// Promotion returns a piece in which piece will be promoted
	Promotion() IPiece

	// WasMoved returns true if a piece was moved from it's starting position
	WasMoved() bool
	// MarkMoved marks piece as moved
	MarkMoved()

	// Equals returns true if two pieces are equal
	Equals(to IPiece) bool
}
