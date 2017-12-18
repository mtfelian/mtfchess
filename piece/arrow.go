package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
)

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

// east launches piece's arrow to the east (x increasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
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

// west launches piece's arrow to the west (x decreasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
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

// north launches piece's arrow to the north (y increasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
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

// south launches piece's arrow to the south (y decreasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
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

// nw launches piece's arrow to the north-west (x decreasing, y increasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
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

// ne launches piece's arrow to the north-east (x increasing, y increasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
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

// sw launches piece's arrow to the south-west (x decreasing, y decreasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
func sw(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	pX, pY := piece.Coord().(rect.Coord).X, board.Dim().(rect.Coord).Y
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

// se launches piece's arrow to the south-east (x increasing, y decreasing) on a board
// set excludeCheckExpose to true to exclude check exposing path
func se(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	pX, pY := piece.Coord().(rect.Coord).X, board.Dim().(rect.Coord).Y
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
