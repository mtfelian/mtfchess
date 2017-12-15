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

// String makes Square implement Stringer
func (s Square) String() string {
	p := " "
	if s.piece != nil {
		p = fmt.Sprintf("%s", s.piece)
	}
	return fmt.Sprintf("%4d[%s](%d,%d)", s.num, p, s.x, s.y)
}
