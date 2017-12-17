package rect

import (
	. "github.com/mtfelian/mtfchess/base"
)

// RectPieceFilter is a piece filter for rectangular board
type RectPieceFilter struct {
	BasePieceFilter
	X []int
	Y []int
}
