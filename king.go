package mtfchess

import (
	. "github.com/mtfelian/mtfchess/board"
)

// King is a chess king
type King struct {
	BasePiece
}

// NewKingPiece creates new king with colour
func NewKingPiece(colour Colour) Piece {
	return &King{
		BasePiece: NewBasePiece(colour, "king", "K♔♚"),
	}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if excludeCheckExpose is false then pairs leading to check-exposing moves also included
func (p *King) dst(b Board, excludeCheckExpose bool) Pairs {
	o := Pairs{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	d := Pairs{}
	for i := 0; i < len(o); i++ {
		x1, y1 := p.X()+o[i].X, p.Y()+o[i].Y
		if x1 < 1 || y1 < 1 || x1 > b.Width() || y1 > b.Height() {
			continue
		}
		// check thet destination cell isn't contains a piece of same colour
		if dstPiece, ok := b.Cell(x1, y1).Piece().(*Knight); ok && dstPiece != nil && dstPiece.Colour() == p.Colour() {
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
func (p *King) Attacks(b Board) Pairs {
	return p.dst(b, false)
}

// Destinations returns a slice of cells coords, making it's legal moves
func (p *King) Destinations(b Board) Pairs {
	return p.dst(b, true)
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *King) Project(x, y int, b Board) Board {
	newBoard := b.Copy()
	newBoard.Empty(p.X(), p.Y())
	newBoard.PlacePiece(x, y, p.Copy())
	return newBoard
}

// Copy a piece
func (p *King) Copy() Piece {
	return &Knight{
		BasePiece: p.BasePiece.Copy(),
	}
}
