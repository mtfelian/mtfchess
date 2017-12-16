package board

// EmptyPiece is an empty piece
type EmptyPiece struct {
	BasePiece
}

// NewEmpty creates new empty figure
func NewEmpty(x, y int) Piece {
	empty := &EmptyPiece{
		BasePiece: NewBasePiece(Transparent, "", "   "),
	}
	empty.SetCoords(x, y)
	return empty
}

// Offsets return an empty slice for an empty piece
func (p *EmptyPiece) Offsets(b Board) Offsets {
	return Offsets{}
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *EmptyPiece) Project(x, y int, b Board) Board {
	newBoard := b.Copy()
	newBoard.Empty(x, y)
	return newBoard
}

// Copy a piece
func (p *EmptyPiece) Copy() Piece {
	return &EmptyPiece{
		BasePiece: p.BasePiece.Copy(),
	}
}
