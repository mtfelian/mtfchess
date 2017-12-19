package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// leaper launches piece's beam like knight (+/- m/n, rot90, +/- n,m) on a board.
// Set excludeCheckExpose to true to exclude check exposing path.
// Set f (front) to 1 to allow movement only forward.
// Set f to -1 to allow only backward movement.
// Set f to 0 to allow both forward and backward piece movement.
// Returns a slice of destination coords.
func leaper(m, n int, piece base.IPiece, board *rect.Board, excludeCheckExpose bool, f int) []base.ICoord {
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
		default: // m != n case
			offsets = []rect.Coord{{n, f * m}, {-n, f * m}, {m, f * n}, {-m, f * n}}
		}
	case 0: // both front and back
		switch {
		case n == 0: // horizontally and vertically, for example, king is an (1,0)-leaper + (1,1)-leaper
			offsets = []rect.Coord{{m, 0}, {-m, 0}, {0, m}, {0, -m}}
		case m == n: // diagonally
			offsets = []rect.Coord{{m, m}, {-m, m}, {m, -m}, {-m, -m}}
		default: // m != n case, for example, knight is (1,2)-leaper
			offsets = []rect.Coord{
				{n, m}, {-n, m}, {n, -m}, {-n, -m},
				{m, n}, {-m, n}, {m, -n}, {-m, -n},
			}
		}
	default:
		panic("wrong front value")
	}

	iOffsets := make([]base.ICoord, len(offsets))
	for i := range offsets {
		iOffsets[i] = offsets[i]
	}
	return inOneStep(piece, board, excludeCheckExpose, iOffsets)
}

// inManySteps returns legal moves for pieces which move in many steps, like rook and bishop
func inManySteps(piece base.IPiece, board base.IBoard, excludeCheckExpose bool, o []rect.Coord, max int) []base.ICoord {
	bW, pX := board.Dim().(rect.Coord).X, piece.Coord().(rect.Coord).X
	bH, pY := board.Dim().(rect.Coord).Y, piece.Coord().(rect.Coord).Y
	// oX, oY - offsets, step - current step of a reader
	notOut := func(oX, oY, step int) bool {
		return oX >= 1-pX && oY >= 1-pY && oX <= bW-pX && oY <= bH-pY && (step < max || max == 0)
	}
	result := []base.ICoord{}
offsets:
	for i := range o {
		for oX, oY, step := o[i].X, o[i].Y, 0; notOut(oX, oY, step); oX, oY, step = oX+o[i].X, oY+o[i].Y, step+1 {
			to := piece.Coord().Add(rect.Coord{oX, oY})
			if excludeCheckExpose && InCheck(piece.Project(to, board), piece.Colour()) {
				continue offsets // check exposed, don't go further in that direction
			}
			if stroke(to, board, piece, &result) {
				break offsets // capture occured, don't go further in that direction
			}
		}
	}
	return result
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

	return inManySteps(piece, board, excludeCheckExpose, offsets, max)
}
