package piece

import (
	. "github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/rect"
)

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func InCheck(board Board, colour Colour) bool {
	baseFilter := BasePieceFilter{
		Colours: []Colour{colour},
		Names:   []string{NewKingPiece(Transparent).Name()},
		Condition: func(p Piece) bool {
			opponentPieces := RectPieceFilter{
				BasePieceFilter: BasePieceFilter{Colours: []Colour{colour.Invert()}},
			}
			return board.FindAttackedCellsBy(opponentPieces).Contains(p.Coord())
		},
	}
	switch board.(type) {
	case *RectBoard:
		return len(board.FindPieces(RectPieceFilter{BasePieceFilter: baseFilter})) > 0
	default:
		panic("invalid board type")
	}
}
