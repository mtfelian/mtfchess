package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
)

// rook launches piece's (rook) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func rook(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	d = append(d, east(piece, board, excludeCheckExpose)...)
	d = append(d, west(piece, board, excludeCheckExpose)...)
	d = append(d, north(piece, board, excludeCheckExpose)...)
	d = append(d, south(piece, board, excludeCheckExpose)...)
	return d
}

// bishop launches piece's (bishop) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func bishop(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	d = append(d, nw(piece, board, excludeCheckExpose)...)
	d = append(d, ne(piece, board, excludeCheckExpose)...)
	d = append(d, sw(piece, board, excludeCheckExpose)...)
	d = append(d, se(piece, board, excludeCheckExpose)...)
	return d
}

// queen launches piece's (queen) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func queen(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	d = append(d, rook(piece, board, excludeCheckExpose)...)
	d = append(d, bishop(piece, board, excludeCheckExpose)...)
	return d
}

// stroke returns true if mine imaginary arrow stroke some piece to coords on board, memorizing it's path
// it returns false if an imaginary arrow still flying
func stroke(to base.ICoord, board base.IBoard, mine base.IPiece, path *[]base.ICoord) bool {
	if dstPiece := board.Cell(to).Piece(); dstPiece != nil {
		if dstPiece.Colour() != mine.Colour() {
			*path = append(*path, to)
		}
		return true
	}
	*path = append(*path, to)
	return false
}

// east launches piece's arrow to the east (x increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func east(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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
func west(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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
func north(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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
func south(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// nw launches piece's arrow to the north-west (x decreasing, y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func nw(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// ne launches piece's arrow to the north-east (x increasing, y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func ne(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// sw launches piece's arrow to the south-west (x decreasing, y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func sw(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// se launches piece's arrow to the south-east (x increasing, y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func se(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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
