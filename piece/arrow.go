package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
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

// stroke returns true if mine imaginary arrow stroke some piece on coords on board, memorizing it's path
// it returns false if an imaginary arrow still flying
func stroke(to base.ICoord, on base.IBoard, mine base.IPiece, path *[]base.ICoord) bool {
	if dstPiece := on.Cell(to).Piece(); dstPiece != nil {
		if dstPiece.Colour() != mine.Colour() {
			*path = append(*path, to)
		}
		return true
	}
	*path = append(*path, to)
	return false
}

// king launches piece's (king) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func king(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return likeKing(piece, b, excludeCheckExpose)
	default:
		panic("invalid coord type")
	}
}

// knight launches piece's (knight) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func knight(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	switch b := board.(type) {
	case *rect.Board:
		return likeKnight21(piece, b, excludeCheckExpose)
	default:
		panic("invalid coord type")
	}
}

// rook launches piece's (rook) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func rook(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	switch b := board.(type) {
	case *rect.Board:
		d = append(d, east(piece, b, excludeCheckExpose, b.Dim().(rect.Coord).X-1)...)
		d = append(d, west(piece, b, excludeCheckExpose, b.Dim().(rect.Coord).X-1)...)
		d = append(d, north(piece, b, excludeCheckExpose, b.Dim().(rect.Coord).Y-1)...)
		d = append(d, south(piece, b, excludeCheckExpose, b.Dim().(rect.Coord).Y-1)...)
	default:
		panic("invalid coord type")
	}
	return d
}

// bishopRect launches piece's (bishop) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func bishop(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	switch b := board.(type) {
	case *rect.Board:
		max, h := b.Dim().(rect.Coord).X, b.Dim().(rect.Coord).Y
		if h > max {
			max = h
		}
		d = append(d, northWest(piece, b, excludeCheckExpose, max-1)...)
		d = append(d, northEast(piece, b, excludeCheckExpose, max-1)...)
		d = append(d, southWest(piece, b, excludeCheckExpose, max-1)...)
		d = append(d, southEast(piece, b, excludeCheckExpose, max-1)...)
	default:
		panic("invalid coord type")
	}
	return d
}

// queen launches piece's (queen) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func queen(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	return append(rook(piece, board, excludeCheckExpose), bishop(piece, board, excludeCheckExpose)...)
}
