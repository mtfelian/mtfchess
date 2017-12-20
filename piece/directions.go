package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
	"reflect"
)

// inOneStep returns legal moves for pieces which move in one step, like knight and king
func inOneStep(piece base.IPiece, board base.IBoard, excludeCheckExpose bool, o []base.ICoord) []base.ICoord {
	result := []base.ICoord{}
	for i := range o {
		to := piece.Coord().Add(o[i])
		if to.OutOf(board) {
			continue
		}
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		stroke(to, board, piece, &result) // here should not break even if true!
	}
	return result
}

// stroke returns true if mine imaginary beam strokes some piece on coords on board, memorizing it's path
// it returns false if an imaginary beam is still going meats no barrier
func stroke(to base.ICoord, on base.IBoard, mine base.IPiece, path *[]base.ICoord) bool {
	if dstPiece := on.Cell(to).Piece(); dstPiece != nil && !reflect.ValueOf(dstPiece).IsNil() {
		if dstPiece.Colour() != mine.Colour() {
			*path = append(*path, to)
		}
		return true
	}
	*path = append(*path, to)
	return false
}

// king launches piece's (king) beams on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func king(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return append(leaper(1, 0, piece, b, excludeCheckExpose, 0), leaper(1, 1, piece, b, excludeCheckExpose, 0)...)
	default:
		panic("invalid coord type")
	}
}

// knight launches piece's (knight) beams on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func knight(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return leaper(1, 2, piece, b, excludeCheckExpose, 0)
	default:
		panic("invalid coord type")
	}
}

// rook launches piece's (rook) beams on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func rook(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return reader(1, 0, piece, b, excludeCheckExpose, 0, 0)
	default:
		panic("invalid coord type")
	}
}

// bishopRect launches piece's (bishop) beams on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func bishop(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return reader(1, 1, piece, b, excludeCheckExpose, 0, 0)
	default:
		panic("invalid coord type")
	}
}

// queen launches piece's (queen) beams on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func queen(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	return append(rook(piece, board, excludeCheckExpose), bishop(piece, board, excludeCheckExpose)...)
}
