package base

import (
	"fmt"
	"reflect"
)

// Cell is a cell on a board
type Cell struct {
	board IBoard // reverse link to a board
	num   int    // sequential number on the board
	coord ICoord // cell coords
	piece IPiece // contains piece
}

// NewCell returns a new cell
func NewCell(board IBoard, num int, coord ICoord) Cell {
	return Cell{
		board: board,
		num:   num,
		coord: coord,
	}
}

// Piece returns a piece on a cell
func (s *Cell) Piece() IPiece {
	return s.piece
}

// SetPiece to p
func (s *Cell) SetPiece(p IPiece) {
	s.piece = p
}

// Empty makes cell empty
func (s *Cell) Empty() {
	s.piece = nil
}

// Copy returns a copy of a cell
func (s *Cell) Copy(board IBoard) Cell {
	newCell := Cell{
		board: board,
		num:   s.num,
		coord: s.coord.Copy(),
		piece: nil,
	}
	if s.piece != nil && !reflect.ValueOf(s.piece).IsNil() {
		newCell.piece = s.piece.Copy()
	}
	return newCell
}

// String makes Cell implement Stringer
func (s Cell) String() string {
	p := " "
	if s.piece != nil && !reflect.ValueOf(s.piece).IsNil() {
		p = fmt.Sprintf("%s", s.piece)
	}
	return fmt.Sprintf("%4d[%s]%s", s.num, p, s.coord)
}
