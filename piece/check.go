package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/rect"
)

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func InCheck(board base.IBoard, colour base.Colour) bool {
	baseFilter := base.PieceFilter{
		Colours: []base.Colour{colour},
		Names:   []string{NewKingPiece(base.Transparent).Name()},
		Condition: func(p base.IPiece) bool {
			opponentPieces := RectPieceFilter{
				PieceFilter: base.PieceFilter{Colours: []base.Colour{colour.Invert()}},
			}
			return board.FindAttackedCellsBy(opponentPieces).Contains(p.Coord())
		},
	}
	switch board.(type) {
	case *RectBoard:
		return len(board.FindPieces(RectPieceFilter{PieceFilter: baseFilter})) > 0
	default:
		panic("invalid board type")
	}
}
