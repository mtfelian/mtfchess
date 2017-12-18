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

// likeKnight21 launches piece's arrow like knight (+/- 2, rot90, +/- 1) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func likeKnight21(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	o := []base.ICoord{
		rect.Coord{-2, -1}, rect.Coord{-2, 1}, rect.Coord{-1, -2}, rect.Coord{-1, 2},
		rect.Coord{1, -2}, rect.Coord{1, 2}, rect.Coord{2, -1}, rect.Coord{2, 1},
	}
	result := []base.ICoord{}
	for i := range o {
		to := piece.Coord().Add(o[i])
		if to.Out(board) {
			continue
		}
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		stroke(to, board, piece, &result) // here should not break even if true!
	}
	return result
}

// east launches piece's arrow to the east (x increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func east(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	pX, bW := piece.Coord().(rect.Coord).X, board.Dim().(rect.Coord).X
	result := []base.ICoord{}
	for x := 1; x <= bW-pX; x++ {
		to := piece.Coord().Add(rect.Coord{x, 0})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}

// west launches piece's arrow to the west (x decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func west(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	pX := piece.Coord().(rect.Coord).X
	result := []base.ICoord{}
	for x := -1; x >= 1-pX; x-- {
		to := piece.Coord().Add(rect.Coord{x, 0})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}

// north launches piece's arrow to the north (y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func north(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y
	result := []base.ICoord{}
	for y := 1; y <= bH-pY; y++ {
		to := piece.Coord().Add(rect.Coord{0, y})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}

// south launches piece's arrow to the south (y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func south(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	pY := piece.Coord().(rect.Coord).Y
	result := []base.ICoord{}
	for y := -1; y >= 1-pY; y-- {
		to := piece.Coord().Add(rect.Coord{0, y})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}

// northWest launches piece's arrow to the north-west (x decreasing, y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func northWest(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y
	pX := piece.Coord().(rect.Coord).X
	result := []base.ICoord{}
	for x, y := -1, 1; x >= 1-pX && y <= bH-pY; x, y = x-1, y+1 {
		to := piece.Coord().Add(rect.Coord{x, y})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}

// northEast launches piece's arrow to the north-east (x increasing, y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func northEast(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y
	pX, bW := piece.Coord().(rect.Coord).X, board.Dim().(rect.Coord).X
	result := []base.ICoord{}
	for x, y := 1, 1; x <= bW-pX && y <= bH-pY; x, y = x+1, y+1 {
		to := piece.Coord().Add(rect.Coord{x, y})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}

// southWest launches piece's arrow to the south-west (x decreasing, y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func southWest(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	pX, pY := piece.Coord().(rect.Coord).X, piece.Coord().(rect.Coord).Y
	result := []base.ICoord{}
	for x, y := -1, -1; x >= 1-pX && y >= 1-pY; x, y = x-1, y-1 {
		to := piece.Coord().Add(rect.Coord{x, y})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}

// southEast launches piece's arrow to the south-east (x increasing, y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func southEast(piece base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	pX, pY := piece.Coord().(rect.Coord).X, piece.Coord().(rect.Coord).Y
	bW := board.Dim().(rect.Coord).X
	result := []base.ICoord{}
	for x, y := 1, -1; x <= bW-pX && y >= 1-pY; x, y = x+1, y-1 {
		to := piece.Coord().Add(rect.Coord{x, y})
		if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
			continue
		}
		if stroke(to, board, piece, &result) {
			break
		}
	}
	return result
}
