package mtfchess

type Empty struct {
	BasePiece
}

// NewEmpty creates new empty figure
func NewEmpty(x, y int) Piece {
	empty := &Empty{
		BasePiece: NewBasePiece(Transparent, "", "   "),
	}
	empty.SetCoords(x, y)
	return empty
}

func (p *Empty) Offsets(b Board) Offsets {
	return Offsets{}
}

func (p *Empty) Project(x, y int, b Board) Board {
	newBoard := b.Copy()
	newBoard.Empty(x, y)
	return newBoard
}

func (p *Empty) Copy() Piece {
	return &Empty{
		BasePiece: p.BasePiece.Copy(),
	}
}
