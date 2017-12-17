package mtfchess

import (
	. "github.com/mtfelian/mtfchess/board"
)

// RectPieceFilter is a piece filter for rectangular board
type RectPieceFilter struct {
	BasePieceFilter
	X []int
	Y []int
}
