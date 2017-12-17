package mtfchess

import (
	. "github.com/mtfelian/mtfchess/board"
)

// Knight is a chess knight
type Knight struct {
	BasePiece
}

// NewKnightPiece creates new knight with colour
func NewKnightPiece(colour Colour) Piece {
	return &Knight{
		BasePiece: NewBasePiece(colour, "knight", "N♘♞"),
	}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if excludeCheckExpose is false then pairs leading to check-exposing moves also included
func (p *Knight) dst(b Board, excludeCheckExpose bool) Pairs {
	o := Pairs{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	d := Pairs{}
	for i := 0; i < len(o); i++ {
		x1, y1 := p.X()+o[i].X, p.Y()+o[i].Y
		if x1 < 1 || y1 < 1 || x1 > b.Width() || y1 > b.Height() {
			continue
		}
		// check that destination cell isn't contains a piece of same colour
		if dstPiece := b.Cell(x1, y1).Piece(); dstPiece != nil && dstPiece.Colour() == p.Colour() {
			continue
		}

		if excludeCheckExpose && p.Project(x1, y1, b).InCheck(p.Colour()) {
			continue
		}
		d = append(d, Pair{X: x1, Y: y1})
	}
	return d
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Knight) Attacks(b Board) Pairs {
	return p.dst(b, false)
}

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Knight) Destinations(b Board) Pairs {
	return p.dst(b, true)
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *Knight) Project(x, y int, b Board) Board {
	newBoard := b.Copy()
	newBoard.Empty(p.X(), p.Y())
	newBoard.PlacePiece(x, y, p.Copy())
	return newBoard
}

// Copy a piece
func (p *Knight) Copy() Piece {
	return &Knight{
		BasePiece: p.BasePiece.Copy(),
	}
}
