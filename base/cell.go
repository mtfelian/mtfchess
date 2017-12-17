package base

import "fmt"

// Cell is a cell on a board
type Cell struct {
	board Board // reverse link to a board
	num   int   // sequential number on the board
	coord Coord // cell coords
	piece Piece // contains piece
}

// NewCell returns a new cell
func NewCell(board Board, num int, coord Coord) Cell {
	return Cell{
		board: board,
		num:   num,
		coord: coord,
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
	s.piece = nil
}

// Copy returns a copy of a cell
func (s *Cell) Copy(board Board) Cell {
	newCell := Cell{
		board: board,
		num:   s.num,
		coord: s.coord.Copy(),
	}
	if s.piece != nil {
		newCell.piece = s.piece.Copy()
	}
	return newCell
}

// String makes Cell implement Stringer
func (s Cell) String() string {
	p := " "
	if s.piece != nil {
		p = fmt.Sprintf("%s", s.piece)
	}
	return fmt.Sprintf("%4d[%s](%s)", s.num, p, s.coord)
}
