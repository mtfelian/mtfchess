package mtfchess

type King struct {
	BasePiece
}

// NewKingPiece creates new king with colour
func NewKingPiece(colour Colour) Piece {
	return &King{
		BasePiece: NewBasePiece(colour, "king", "K♔♚"),
	}
}

func (p *King) Offsets(b IBoard) Offsets {
	o := []Pair{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < len(o); i++ {
		remove := func() {
			o = append(o[:i], o[i+1:]...)
			i--
		}
		x1, y1 := p.x+o[i].X, p.y+o[i].Y
		if x1 < 1 || y1 < 1 || x1 > b.Width() || y1 > b.Height() {
			remove()
			continue
		}
		// check thet destination square isn't contains a piece of same colour
		if dstPiece, ok := b.Square(x1, y1).piece.(*Knight); ok && dstPiece != nil && dstPiece.Colour() == p.Colour() {
			remove()
			continue
		}

		if p.Project(x1, y1, b).InCheck(p.Colour()) {
			remove()
			continue
		}
	}
	return o
}

func (p *King) Project(x, y int, b IBoard) IBoard {
	newBoard := b.Copy()
	newBoard.Empty(p.x, p.y)
	newBoard.PlacePiece(x, y, p.Copy())
	return newBoard
}

func (p *King) Copy() Piece {
	return &Knight{
		BasePiece: p.BasePiece.Copy(),
	}
}
