package mtfchess

import (
	. "github.com/mtfelian/mtfchess/board"
)

// Row is a row of cells
type RectRow []Cell

// Copy returns a deep copy of row
func (r RectRow) Copy(board Board) RectRow {
	newRow := make(RectRow, len(r))
	for i := range r {
		newRow[i] = r[i].Copy(board)
	}
	return newRow
}

// Cells is a matrix of cells
type RectCells []RectRow

// Copy returns a deep copy of cells
func (s RectCells) Copy(board Board) Cells {
	newCells := make(RectCells, len(s))
	for i := range s {
		newCells[i] = s[i].Copy(board)
	}
	return newCells
}
