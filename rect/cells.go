package rect

import (
	"github.com/mtfelian/mtfchess/base"
)

// Row is a row of cells
type Row []base.Cell

// Copy returns a deep copy of row
func (r Row) Copy(board base.IBoard) Row {
	newRow := make(Row, len(r))
	for i := range r {
		newRow[i] = r[i].Copy(board)
	}
	return newRow
}

// Cells is a matrix of cells
type Cells []Row

// Copy returns a deep copy of cells
func (s Cells) Copy(board base.IBoard) base.ICells {
	newCells := make(Cells, len(s))
	for i := range s {
		newCells[i] = s[i].Copy(board)
	}
	return newCells
}
