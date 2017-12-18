package piece

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
)

// knightRect launches piece's (knight) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func knightRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	// todo
	return d
}

// rookRect launches piece's (rook) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func rookRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	d = append(d, eastRect(piece, board, excludeCheckExpose)...)
	d = append(d, westRect(piece, board, excludeCheckExpose)...)
	d = append(d, northRect(piece, board, excludeCheckExpose)...)
	d = append(d, southRect(piece, board, excludeCheckExpose)...)
	return d
}

// bishopRect launches piece's (bishop) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func bishopRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	d = append(d, nwRect(piece, board, excludeCheckExpose)...)
	d = append(d, neRect(piece, board, excludeCheckExpose)...)
	d = append(d, swRect(piece, board, excludeCheckExpose)...)
	d = append(d, seRect(piece, board, excludeCheckExpose)...)
	return d
}

// queenRect launches piece's (queen) arrows on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func queenRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
	d := []base.ICoord{}
	d = append(d, rookRect(piece, board, excludeCheckExpose)...)
	d = append(d, bishopRect(piece, board, excludeCheckExpose)...)
	return d
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

// eastRect launches piece's arrow to the east (x increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func eastRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// westRect launches piece's arrow to the west (x decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func westRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// northRect launches piece's arrow to the north (y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func northRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// southRect launches piece's arrow to the south (y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func southRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// nwRect launches piece's arrow to the north-west (x decreasing, y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func nwRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// neRect launches piece's arrow to the north-east (x increasing, y increasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func neRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// swRect launches piece's arrow to the south-west (x decreasing, y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func swRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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

// seRect launches piece's arrow to the south-east (x increasing, y decreasing) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func seRect(piece base.IPiece, board base.IBoard, excludeCheckExpose bool) []base.ICoord {
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
