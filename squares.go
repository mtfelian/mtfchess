package mtfchess

// Squares is a matrix of squares
type Squares []Row

// Copy returns a deep copy of squares
func (s Squares) Copy(board Board) Squares {
	newSquares := make(Squares, len(s))
	for i := range s {
		newSquares[i] = s[i].Copy(board)
	}
	return newSquares
}
