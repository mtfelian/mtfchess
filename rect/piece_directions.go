package rect

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/mtfelian/utils"
)

// inOneStep returns legal moves for pieces which move in one step, like knight and king
func inOneStep(piece base.IPiece, board base.IBoard, moving bool, o []base.ICoord, moveType int) []base.ICoord {
	result := []base.ICoord{}
	for i := range o {
		to := piece.Coord().Add(o[i])
		if to.OutOf(board) {
			continue
		}
		if moving && board.Project(piece, to).InCheck(piece.Colour()) {
			continue
		}
		stroke(to, moving, board, piece, &result, moveType) // here should not break even if true!
	}
	return result
}

// stroke returns true if mine imaginary beam strokes some piece on coords on board, memorizing it's path
// it returns false if an imaginary beam is still going meating no barrier
// to is a destination cell coords
// moving - set it to true if the func should return possible legal moves, set it to false to return attacked cells
// on is a board on which piece is moving
// mine is a moving piece
// path is a pointer to a slice of coords to add
// moveType is a type of move: only capturing, only non-capturing, or any
func stroke(to base.ICoord, moving bool, on base.IBoard, mine base.IPiece, path *[]base.ICoord, moveType int) bool {
	dstPiece := on.Cell(to).Piece()
	// destination cell contains another piece
	if dstPiece != nil {
		// if we are only calculating attacking cells, or if can capture
		if SliceContains(moveType, []int{moveAny, moveCapture}) && (!moving || dstPiece.Colour() != mine.Colour()) {
			*path = append(*path, to)
		}
		return true
	}

	// dstPiece == nil, empty cell
	if moveType == moveAny || (moving && moveType == moveNonCapture) || (!moving && moveType == moveCapture) {
		*path = append(*path, to)
	}
	return false
}

const (
	moveAny        = iota // capture / non-capturing move
	moveCapture           // only capture
	moveNonCapture        // only non-capturing move
)

// leaper launches piece's beam like knight (+/- m/n, rot90, +/- n,m) on a board.
// Set moving to true to exclude check exposing path and defending own piece path.
// Set f (front) to 1 to allow movement only forward.
// Set f to -1 to allow only backward movement.
// Set f to 0 to allow both forward and backward piece movement.
// Returns a slice of destination coords.
func leaper(m, n int, piece base.IPiece, board *Board, moving bool, f int, moveType int) []base.ICoord {
	if piece.Colour() == Black {
		f *= -1
	}
	if m == 0 { // only n can be 0
		if n == 0 {
			panic("reader can't be (0,0)")
		}
		m, n = n, 0
	}
	offsets := []Coord{}
	switch f {
	case 1, -1: // only front or only back
		switch {
		case n == 0: // horizontally and vertically
			offsets = []Coord{{0, f * m}} // removed {{m, 0}, {-m, 0}}, it's side movements
		case m == n: // diagonally
			offsets = []Coord{{m, f * m}, {-m, f * m}}
		default: // m != n case
			offsets = []Coord{{n, f * m}, {-n, f * m}, {m, f * n}, {-m, f * n}}
		}
	case 0: // both front and back
		switch {
		case n == 0: // horizontally and vertically, for example, king is an (1,0)-leaper + (1,1)-leaper
			offsets = []Coord{{m, 0}, {-m, 0}, {0, m}, {0, -m}}
		case m == n: // diagonally
			offsets = []Coord{{m, m}, {-m, m}, {m, -m}, {-m, -m}}
		default: // m != n case, for example, knight is (1,2)-leaper
			offsets = []Coord{
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

	return inOneStep(piece, board, moving, iOffsets, moveType)
}

// inManySteps returns legal moves for pieces which move in many steps, like rook and bishop
func inManySteps(piece base.IPiece, board *Board, moving bool, o []Coord, max int, moveType int) []base.ICoord {
	bW, pX := board.Dim().(Coord).X, piece.Coord().(Coord).X
	bH, pY := board.Dim().(Coord).Y, piece.Coord().(Coord).Y
	// oX, oY - offsets, step - current step of a reader
	notOut := func(oX, oY, step int) bool {
		return oX >= 1-pX && oY >= 1-pY && oX <= bW-pX && oY <= bH-pY && (step < max || max == 0)
	}
	result := []base.ICoord{}
directions:
	for i := range o {
		for oX, oY, step := o[i].X, o[i].Y, 0; notOut(oX, oY, step); oX, oY, step = oX+o[i].X, oY+o[i].Y, step+1 {
			to := piece.Coord().Add(Coord{oX, oY})
			if moving && board.Project(piece, to).InCheck(piece.Colour()) {
				continue // should continue in same direction (may be further capture releases check?)
			}
			if stroke(to, moving, board, piece, &result, moveType) {
				continue directions // capture occurred, don't go further in that direction
			}
		}
	}
	return result
}

// reader launches (m,n)-reader piece's beam on a board.
// Set moving to true to exclude check exposing path and defending own pieces path.
// Set max to non-0 value to restrict the maximum steps to move in each one direction.
// max value of 0 means no maximum steps restriction.
// Set f (front) to 1 to allow movement only forward.
// Set f to -1 to allow only backward movement.
// Set f to 0 to allow both forward and backward piece movement.
// Returns a slice of destination coords.
func reader(m, n int, piece base.IPiece, board *Board, moving bool, max int, f int, moveType int) []base.ICoord {
	if piece.Colour() == Black {
		f *= -1
	}
	if m == 0 { // only n can be 0
		if n == 0 {
			panic("reader can't be (0,0)")
		}
		m, n = n, 0
	}

	offsets := []Coord{}
	switch f {
	case 1, -1: // only front or only back
		switch {
		case n == 0: // horizontally and vertically
			offsets = []Coord{{0, f * m}} // removed {{m, 0}, {-m, 0}}, it's side movements
		case m == n: // diagonally
			offsets = []Coord{{m, f * m}, {-m, f * m}}
		default: // on a special offsets (see "nightreader" - (1,2)-reader)
			offsets = []Coord{{n, f * m}, {-n, f * m}, {m, f * n}, {-m, f * n}}
		}
	case 0: // both front and back
		switch {
		case n == 0: // horizontally and vertically
			offsets = []Coord{{m, 0}, {-m, 0}, {0, m}, {0, -m}}
		case m == n: // diagonally
			offsets = []Coord{{m, m}, {-m, m}, {m, -m}, {-m, -m}}
		default: // on a special offsets (see "nightreader" - (1,2)-reader)
			offsets = []Coord{
				{n, m}, {-n, m}, {n, -m}, {-n, -m},
				{m, n}, {-m, n}, {m, -n}, {-m, -n},
			}
		}
	default:
		panic("wrong front value")
	}

	return inManySteps(piece, board, moving, offsets, max, moveType)
}
