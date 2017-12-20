package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func InCheck(board base.IBoard, colour Colour) bool {
	king := board.King(colour)
	return king != nil && board.FindAttackedCellsBy(base.PieceFilter{Colours: []Colour{colour.Invert()}}).Contains(king.Coord())
}
