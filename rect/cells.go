package rect

import (
	"github.com/mtfelian/mtfchess/base"
)

// Row is a row of cells
type RectRow []base.Cell

// Copy returns a deep copy of row
func (r RectRow) Copy(board base.IBoard) RectRow {
	newRow := make(RectRow, len(r))
	for i := range r {
		newRow[i] = r[i].Copy(board)
	}
	return newRow
}

// Cells is a matrix of cells
type RectCells []RectRow

// Copy returns a deep copy of cells
func (s RectCells) Copy(board base.IBoard) base.ICells {
	newCells := make(RectCells, len(s))
	for i := range s {
		newCells[i] = s[i].Copy(board)
	}
	return newCells
}
