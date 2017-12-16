package board

// Row is a row of cells
type Row []Cell

// Copy returns a deep copy of row
func (r Row) Copy(board Board) Row {
	newRow := make(Row, len(r))
	for i := range r {
		newRow[i] = r[i].Copy(board)
	}
	return newRow
}
