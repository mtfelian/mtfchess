package base

import . "github.com/mtfelian/mtfchess/colour"

// PieceFilter is a base piece filter
type PieceFilter struct {
	Names     []string
	Colours   []Colour
	Condition func(IPiece) bool
}
