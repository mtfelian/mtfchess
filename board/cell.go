package board

import "fmt"

// Cell is a cell on a board
type Cell struct {
	board Board // reverse link to a board
	num   int   // sequential number on the board
	x     int   // [1;board.width]
	y     int   // [1;board.height]
	piece Piece // contains piece
}

// NewCell returns a new cell
func NewCell(board Board, num, x, y int) Cell {
	return Cell{
		board: board,
		num:   num,
		x:     x,
		y:     y,
	}
}

// Piece returns a piece on a cell
func (s *Cell) Piece() Piece {
	return s.piece
}

// SetPiece to p
func (s *Cell) SetPiece(p Piece) {
	s.piece = p
}

// Empty makes cell empty
func (s *Cell) Empty() {
	s.piece = NewEmpty(s.x, s.y)
}

// Copy returns a copy of a cell
func (s *Cell) Copy(board Board) Cell {
	return Cell{
		board: board,
		num:   s.num,
		x:     s.x,
		y:     s.y,
		piece: s.piece.Copy(),
	}
}

// String makes Cell implement Stringer
func (s Cell) String() string {
	p := " "
	if s.piece != nil {
		p = fmt.Sprintf("%s", s.piece)
	}
	return fmt.Sprintf("%4d[%s](%d,%d)", s.num, p, s.x, s.y)
}
