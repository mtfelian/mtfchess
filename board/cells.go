package board

// Cells is a matrix of cells
type Cells []Row

// Copy returns a deep copy of cells
func (s Cells) Copy(board Board) Cells {
	newCells := make(Cells, len(s))
	for i := range s {
		newCells[i] = s[i].Copy(board)
	}
	return newCells
}
