package base

// Cells is a board cells
type Cells interface {
	Copy(board Board) Cells
}
