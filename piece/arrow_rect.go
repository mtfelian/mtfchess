package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
)

// likeKnight21 launches piece's arrow like knight (+/- 2, rot90, +/- 1) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func likeKnight21(me base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	offsets := []base.ICoord{
		rect.Coord{-2, -1}, rect.Coord{-2, 1}, rect.Coord{-1, -2}, rect.Coord{-1, 2},
		rect.Coord{1, -2}, rect.Coord{1, 2}, rect.Coord{2, -1}, rect.Coord{2, 1},
	}
	return inOneStep(me, board, excludeCheckExpose, offsets)
}

// likeKing launches piece's arrow like king (+/-1 around) on a board.
// It is equivalent to sequently append() all the directions with 'max' set to 1, but this works much faster.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func likeKing(me base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	offsets := []base.ICoord{
		rect.Coord{-1, -1}, rect.Coord{-1, 0}, rect.Coord{-1, 1}, rect.Coord{0, -1},
		rect.Coord{0, 1}, rect.Coord{1, -1}, rect.Coord{1, 0}, rect.Coord{1, 1},
	}
	return inOneStep(me, board, excludeCheckExpose, offsets)
}

// east launches piece's arrow to the east (x increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func east(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	pX, bW := piece.Coord().(rect.Coord).X, board.Dim().(rect.Coord).X
	result := []base.ICoord{}
	for x, m := 1, 0; x <= bW-pX && m < max; x, m = x+1, m+1 {
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
func west(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	pX := piece.Coord().(rect.Coord).X
	result := []base.ICoord{}
	for x, m := -1, 0; x >= 1-pX && m < max; x, m = x-1, m+1 {
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
func north(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y
	result := []base.ICoord{}
	for y, m := 1, 0; y <= bH-pY && m < max; y, m = y+1, m+1 {
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
func south(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	pY := piece.Coord().(rect.Coord).Y
	result := []base.ICoord{}
	for y, m := -1, 0; y >= 1-pY && m < max; y, m = y-1, m+1 {
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
func northWest(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y
	pX := piece.Coord().(rect.Coord).X
	result := []base.ICoord{}
	for x, y, m := -1, 1, 0; x >= 1-pX && y <= bH-pY && m < max; x, y, m = x-1, y+1, m+1 {
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
func northEast(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y
	pX, bW := piece.Coord().(rect.Coord).X, board.Dim().(rect.Coord).X
	result := []base.ICoord{}
	for x, y, m := 1, 1, 0; x <= bW-pX && y <= bH-pY && m < max; x, y, m = x+1, y+1, m+1 {
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
func southWest(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	pX, pY := piece.Coord().(rect.Coord).X, piece.Coord().(rect.Coord).Y
	result := []base.ICoord{}
	for x, y, m := -1, -1, 0; x >= 1-pX && y >= 1-pY && m < max; x, y, m = x-1, y-1, m+1 {
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
func southEast(piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int) []base.ICoord {
	pX, pY := piece.Coord().(rect.Coord).X, piece.Coord().(rect.Coord).Y
	bW := board.Dim().(rect.Coord).X
	result := []base.ICoord{}
	for x, y, m := 1, -1, 0; x <= bW-pX && y >= 1-pY && m < max; x, y, m = x+1, y-1, m+1 {
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
