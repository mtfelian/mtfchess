package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
)

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
		d = append(d, east(piece, b, excludeCheckExpose)...)
		d = append(d, west(piece, b, excludeCheckExpose)...)
		d = append(d, north(piece, b, excludeCheckExpose)...)
		d = append(d, south(piece, b, excludeCheckExpose)...)
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
		d = append(d, northWest(piece, b, excludeCheckExpose)...)
		d = append(d, northEast(piece, b, excludeCheckExpose)...)
		d = append(d, southWest(piece, b, excludeCheckExpose)...)
		d = append(d, southEast(piece, b, excludeCheckExpose)...)
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
