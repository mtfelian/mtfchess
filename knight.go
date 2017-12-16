package mtfchess

import (
	"github.com/mtfelian/cli"
)

type Knight struct {
	BasePiece
}

// NewKnight creates new knight with colour
func NewKnight(colour Colour) Piece {
	return &Knight{
		BasePiece: NewBasePiece(colour),
	}
}

func (p *Knight) Name() string {
	return "knight"
}

func (p *Knight) String() string { // ♘♞
	return map[Colour]string{White: cli.Sprintf("{W|N{0|"), Black: cli.Sprintf("{A|N{0|")}[p.Colour()]
}

func (p *Knight) Offsets(b *Board) Offsets {
	o := []Pair{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	for i := 0; i < len(o); i++ {
		remove := func() {
			o = append(o[:i], o[i+1:]...)
			i--
		}
		x1, y1 := p.x+o[i].X, p.y+o[i].Y
		if x1 < 1 || y1 < 1 || x1 > b.width || y1 > b.height {
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

func (p *Knight) Project(x, y int, b *Board) *Board {
	newBoard := b.Copy()
	newBoard.Empty(p.x, p.y)
	newBoard.PlacePiece(x, y, p.Copy())
	return newBoard
}

func (p *Knight) Copy() Piece {
	return &Knight{
		BasePiece: p.BasePiece.Copy(),
	}
}
