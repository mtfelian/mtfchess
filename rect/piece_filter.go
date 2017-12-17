package rect

import (
	"github.com/mtfelian/mtfchess/base"
)

// RectPieceFilter is a piece filter for rectangular board
type RectPieceFilter struct {
	base.PieceFilter
	X []int
	Y []int
}
