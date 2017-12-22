package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/utils"
)

// inOneStep returns legal moves for pieces which move in one step, like knight and king
func inOneStep(piece base.IPiece, board base.IBoard, moving bool, o []base.ICoord, moveType int) []base.ICoord {
	result := []base.ICoord{}
	for i := range o {
		to := piece.Coord().Add(o[i])
		if to.OutOf(board) {
			continue
		}
		if moving && InCheck(board.Project(piece, to), piece.Colour()) {
			continue
		}
		stroke(to, moving, board, piece, &result, moveType) // here should not break even if true!
	}
	return result
}

// stroke returns true if mine imaginary beam strokes some piece on coords on board, memorizing it's path
// it returns false if an imaginary beam is still going meating no barrier
// to is a destination cell coords
// moving - set it to true if the func should return possible legal moves, set it to false to return attacked cells
// on is a board on which piece is moving
// mine is a moving piece
// path is a pointer to a slice of coords to add
// moveType is a type of move: only capturing, only non-capturing, or any
func stroke(to base.ICoord, moving bool, on base.IBoard, mine base.IPiece, path *[]base.ICoord, moveType int) bool {
	// destination cell contains another piece
	if dstPiece := on.Cell(to).Piece(); dstPiece != nil && SliceContains(moveType, []int{moveAny, moveCapture}) {
		// if we are only calculating attacking cells, or if can capture
		if !moving || dstPiece.Colour() != mine.Colour() {
			*path = append(*path, to)
		}
		return true
	}
	// dstPiece == nil, empty cell
	if (moving && moveType != moveCapture) || moveType == moveAny || (!moving && moveType == moveCapture) {
		*path = append(*path, to)
	}
	return false
}
