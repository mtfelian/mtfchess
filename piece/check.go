package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func InCheck(board base.IBoard, colour Colour) bool {
	return len(board.FindPieces(base.PieceFilter{
		Colours: []Colour{colour},
		Names:   []string{NewKing(Transparent).Name()},
		Condition: func(p base.IPiece) bool {
			opponentPieces := base.PieceFilter{Colours: []Colour{colour.Invert()}}
			return board.FindAttackedCellsBy(opponentPieces).Contains(p.Coord())
		},
	})) > 0
}
