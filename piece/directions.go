package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
)

// inOneStep returns legal moves for pieces which move in one step, like knight and king
func inOneStep(piece base.IPiece, board base.IBoard, moving bool, o []base.ICoord) []base.ICoord {
	result := []base.ICoord{}
	for i := range o {
		to := piece.Coord().Add(o[i])
		if to.OutOf(board) {
			continue
		}
		if moving && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		stroke(to, moving, board, piece, &result) // here should not break even if true!
	}
	return result
}

// stroke returns true if mine imaginary beam strokes some piece on coords on board, memorizing it's path
// it returns false if an imaginary beam is still going meats no barrier
func stroke(to base.ICoord, moving bool, on base.IBoard, mine base.IPiece, path *[]base.ICoord) bool {
	if dstPiece := on.Cell(to).Piece(); dstPiece != nil {
		if !moving || dstPiece.Colour() != mine.Colour() {
			*path = append(*path, to)
		}
		return true
	}
	*path = append(*path, to)
	return false
}

// king launches piece's (king) beams on a board.
// Set moving to true to exclude check exposing path and defending own pieces path.
// Returns a slice of destination coords.
func king(piece base.IPiece, board base.IBoard, moving bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return append(leaper(1, 0, piece, b, moving, 0), leaper(1, 1, piece, b, moving, 0)...)
	default:
		panic("invalid coord type")
	}
}

// knight launches piece's (knight) beams on a board.
// Set moving to true to exclude check exposing path and defending own pieces path.
// Returns a slice of destination coords.
func knight(piece base.IPiece, board base.IBoard, moving bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return leaper(1, 2, piece, b, moving, 0)
	default:
		panic("invalid coord type")
	}
}

// rook launches piece's (rook) beams on a board.
// Set moving to true to exclude check exposing path and defending own pieces path.
// Returns a slice of destination coords.
func rook(piece base.IPiece, board base.IBoard, moving bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return reader(1, 0, piece, b, moving, 0, 0)
	default:
		panic("invalid coord type")
	}
}

// bishopRect launches piece's (bishop) beams on a board.
// Set moving to true to exclude check exposing path and defending own pieces path.
// Returns a slice of destination coords.
func bishop(piece base.IPiece, board base.IBoard, moving bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return reader(1, 1, piece, b, moving, 0, 0)
	default:
		panic("invalid coord type")
	}
}

// queen launches piece's (queen) beams on a board.
// Set moving to true to exclude check exposing path and defending own pieces path.
// Returns a slice of destination coords.
func queen(piece base.IPiece, board base.IBoard, moving bool) []base.ICoord {
	return append(rook(piece, board, moving), bishop(piece, board, moving)...)
}
