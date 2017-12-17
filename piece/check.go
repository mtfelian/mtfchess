package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/mtfelian/mtfchess/rect"
)

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func InCheck(board base.IBoard, colour Colour) bool {
	baseFilter := base.PieceFilter{
		Colours: []Colour{colour},
		Names:   []string{NewKing(Transparent).Name()},
		Condition: func(p base.IPiece) bool {
			opponentPieces := PieceFilter{
				PieceFilter: base.PieceFilter{Colours: []Colour{colour.Invert()}},
			}
			return board.FindAttackedCellsBy(opponentPieces).Contains(p.Coord())
		},
	}
	switch board.(type) {
	case *Board:
		return len(board.FindPieces(PieceFilter{PieceFilter: baseFilter})) > 0
	default:
		panic("invalid board type")
	}
}
