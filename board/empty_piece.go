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

func (p *EmptyPiece) Offsets(b Board) Offsets {
	return Offsets{}
}

func (p *EmptyPiece) Project(x, y int, b Board) Board {
	newBoard := b.Copy()
	newBoard.Empty(x, y)
	return newBoard
}

func (p *EmptyPiece) Copy() Piece {
	return &EmptyPiece{
		BasePiece: p.BasePiece.Copy(),
	}
}
