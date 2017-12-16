package board

import "fmt"

// Square is a square on a board
type Square struct {
	board Board // reverse link to a board
	num   int   // sequential number on the board
	x     int   // [1;board.width]
	y     int   // [1;board.height]
	piece Piece // contains piece
}

// NewSquare returns a new square
func NewSquare(board Board, num, x, y int) Square {
	return Square{
		board: board,
		num:   num,
		x:     x,
		y:     y,
	}
}

// Piece returns a piece on a square
func (s *Square) Piece() Piece {
	return s.piece
}

// SetPiece to p
func (s *Square) SetPiece(p Piece) {
	s.piece = p
}

// Empty makes square empty
func (s *Square) Empty() {
	s.piece = NewEmpty(s.x, s.y)
}

// Copy returns a copy of a square
func (s *Square) Copy(board Board) Square {
	return Square{
		board: board,
		num:   s.num,
		x:     s.x,
		y:     s.y,
		piece: s.piece.Copy(),
	}
}

// String makes Square implement Stringer
func (s Square) String() string {
	p := " "
	if s.piece != nil {
		p = fmt.Sprintf("%s", s.piece)
	}
	return fmt.Sprintf("%4d[%s](%d,%d)", s.num, p, s.x, s.y)
}
