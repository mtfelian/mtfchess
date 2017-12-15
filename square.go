package mtfchess

import "fmt"

// Square is a square on a board
type Square struct {
	board *Board
	num   int
	x     int // [1;board.width]
	y     int // [1;max board.height]
	piece Piece
}

// Empty makes square empty
func (s *Square) Empty() {
	s.piece = NewEmpty(s.x, s.y)
}

// Copy returns a copy of a square
func (s *Square) Copy(board *Board) Square {
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
