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

// offsets returns a slice of offsets relative to piece coords, making it's legal moves
// if excludeCheckExpose is false then offsets leading to check-exposing moves also included
func (p *King) offsets(b Board, excludeCheckExpose bool) Offsets {
	o := Offsets{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < len(o); i++ {
		remove := func() {
			o = append(o[:i], o[i+1:]...)
			i--
		}
		x1, y1 := p.X()+o[i].X, p.Y()+o[i].Y
		if x1 < 1 || y1 < 1 || x1 > b.Width() || y1 > b.Height() {
			remove()
			continue
		}
		// check thet destination cell isn't contains a piece of same colour
		if dstPiece, ok := b.Cell(x1, y1).Piece().(*Knight); ok && dstPiece != nil && dstPiece.Colour() == p.Colour() {
			remove()
			continue
		}

		if excludeCheckExpose && p.Project(x1, y1, b).InCheck(p.Colour()) {
			remove()
			continue
		}
	}
	return o
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *King) Attacks(b Board) Pairs {
	o := p.offsets(b, false)
	return o.Pairs(p)
}

// Offsets returns a slice of offsets relative to piece coords, making it's legal moves
func (p *King) Offsets(b Board) Offsets {
	return p.offsets(b, true)
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
