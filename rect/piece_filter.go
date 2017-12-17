package rect

import (
	"github.com/mtfelian/mtfchess/base"
)

// PieceFilter is a piece filter for rectangular board
type PieceFilter struct {
	base.PieceFilter
	X []int
	Y []int
}
