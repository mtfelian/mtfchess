package piece

import (
	"github.com/mtfelian/mtfchess/base"
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
// it returns false if an imaginary beam is still going meats no barrier
func stroke(to base.ICoord, moving bool, on base.IBoard, mine base.IPiece, path *[]base.ICoord, moveType int) bool {
	if dstPiece := on.Cell(to).Piece(); dstPiece != nil {
		if !moving || dstPiece.Colour() != mine.Colour() {
			*path = append(*path, to)
		}
		return true
	}
	*path = append(*path, to)
	return false
}
