package mtfchess

import (
	"github.com/mtfelian/cli"
)

type Knight struct {
	board  *Board
	colour Colour
	x, y   int
}

// NewKnight creates new knight on x, y coords of a board with colour
func NewKnight(board *Board, colour Colour) Piece {
	return &Knight{
		board:  board,
		colour: colour,
	}
}

func (p *Knight) Board() *Board {
	return p.board
}

func (p *Knight) Name() string {
	return "knight"
}

func (p *Knight) CanJump() bool {
	return true
}

func (p *Knight) Colour() Colour {
	return p.colour
}

func (p *Knight) String() string { // ♘♞
	return map[Colour]string{White: cli.Sprintf("{W|N{0|"), Black: cli.Sprintf("{A|N{0|")}[p.Colour()]
}

func (p *Knight) SetCoords(x, y int) {
	p.x, p.y = x, y
}

func (p *Knight) Offsets() Offsets {
	o := []Pair{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	b := p.Board()
	for i := 0; i < len(o); i++ {
		if p.x+o[i].X < 1 || p.y+o[i].Y < 1 || p.x+o[i].X > b.width || p.y+o[i].Y > b.height {
			o = append(o[:i], o[i+1:]...)
			i--
			continue
		}
		// check thet destination square isn't contains a piece of same colour
		if dstPiece, ok := b.Square(p.y+o[i].Y, p.x+o[i].X).piece.(*Knight); ok && dstPiece != nil && dstPiece.Colour() == p.Colour() {
			o = append(o[:i], o[i+1:]...)
			i--
			continue
		}
	}
	return o
}
