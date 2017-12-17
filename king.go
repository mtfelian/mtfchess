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
func (p *King) dst(b Board, excludeCheckExpose bool) Coords {
	switch b.Dim().(type) {
	case RectCoord:
		o := []RectCoord{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		d := []RectCoord{}
		for i := 0; i < len(o); i++ {
			c1 := p.Coord().Add(o[i])
			if c1.Out(b) {
				continue
			}
			// check thet destination cell isn't contains a piece of same colour
			if dstPiece := b.Cell(c1).Piece(); dstPiece != nil && dstPiece.Colour() == p.Colour() {
				continue
			}

			if excludeCheckExpose && p.Project(c1, b).InCheck(p.Colour()) {
				continue
			}
			d = append(d, c1.(RectCoord))
		}
		return NewRectCoords(d)
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *King) Attacks(b Board) Coords {
	return p.dst(b, false)
}

// Destinations returns a slice of cells coords, making it's legal moves
func (p *King) Destinations(b Board) Coords {
	return p.dst(b, true)
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *King) Project(to Coord, b Board) Board {
	newBoard := b.Copy()
	newBoard.Empty(to)
	newBoard.PlacePiece(to, p.Copy())
	return newBoard
}

// Copy a piece
func (p *King) Copy() Piece {
	return &Knight{
		BasePiece: p.BasePiece.Copy(),
	}
}
