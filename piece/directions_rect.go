package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// likeKnight21 launches piece's beam like knight (+/- 2, rot90, +/- 1) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Returns a slice of destination coords.
func likeKnight21(me base.IPiece, board *rect.Board, excludeCheckExpose bool) []base.ICoord {
	offsets := []base.ICoord{
		rect.Coord{-2, -1}, rect.Coord{-2, 1}, rect.Coord{-1, -2}, rect.Coord{-1, 2},
		rect.Coord{1, -2}, rect.Coord{1, 2}, rect.Coord{2, -1}, rect.Coord{2, 1},
	}
	return inOneStep(me, board, excludeCheckExpose, offsets)
}

// likeKing launches piece's beam like king (+/-1 around) on a board.
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

// reader launches (m,n)-reader piece's beam on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Set max to non-0 value to restrict the maximum steps to move in each one direction.
// max value of 0 means no maximum steps restriction.
// Set f (front) to 1 to allow movement only forward.
// Set f to -1 to allow only backward movement.
// Set f to 0 to allow both forward and backward piece movement.
// Returns a slice of destination coords.
func reader(m, n int, piece base.IPiece, board *rect.Board, excludeCheckExpose bool, max int, f int) []base.ICoord {
	bW, pX := board.Dim().(rect.Coord).X, piece.Coord().(rect.Coord).X
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y

	if piece.Colour() == Black {
		f *= -1
	}

	if m == 0 { // only n can be 0
		if n == 0 {
			panic("reader can't be (0,0)")
		}
		m, n = n, 0
	}

	offsets := []rect.Coord{}
	switch f {
	case 1, -1: // only front or only back
		switch {
		case n == 0: // horizontally and vertically
			offsets = []rect.Coord{{m, 0}, {-m, 0}, {0, f * m}}
		case m == n: // diagonally
			offsets = []rect.Coord{{m, f * m}, {-m, f * m}}
		default: // on a special offsets (see "nightreader" - (1,2)-reader)
			offsets = []rect.Coord{{n, f * m}, {-n, f * m}, {m, f * n}, {-m, f * n}}
		}
	case 0: // both front and back
		switch {
		case n == 0: // horizontally and vertically
			offsets = []rect.Coord{{m, 0}, {-m, 0}, {0, m}, {0, -m}}
		case m == n: // diagonally
			offsets = []rect.Coord{{m, m}, {-m, m}, {m, -m}, {-m, -m}}
		default: // on a special offsets (see "nightreader" - (1,2)-reader)
			offsets = []rect.Coord{
				{n, m}, {-n, m}, {n, -m}, {-n, -m},
				{m, n}, {-m, n}, {m, -n}, {-m, -n},
			}
		}
	default:
		panic("wrong front value")
	}

	// oX, oY - offsets, step - current step of a reader
	notOut := func(oX, oY, step int) bool {
		return oX >= 1-pX && oY >= 1-pY && oX <= bW-pX && oY <= bH-pY && (step < max || max == 0)
	}
	result := []base.ICoord{}
	for _, o := range offsets {
		for oX, oY, step := o.X, o.Y, 0; notOut(oX, oY, step); oX, oY, step = oX+o.X, oY+o.Y, step+1 {
			to := piece.Coord().Add(rect.Coord{oX, oY})
			if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
				continue
			}
			if stroke(to, board, piece, &result) {
				break
			}
		}
	}
	return result
}
