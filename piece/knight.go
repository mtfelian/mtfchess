package piece

import (
	. "github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/rect"
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
func (p *Knight) dst(b Board, excludeCheckExpose bool) Coords {
	o, d := []Coord{}, []Coord{}

	switch b.Dim().(type) {
	case RectCoord:
		o = []Coord{
			RectCoord{-2, -1}, RectCoord{-2, 1}, RectCoord{-1, -2}, RectCoord{-1, 2},
			RectCoord{1, -2}, RectCoord{1, 2}, RectCoord{2, -1}, RectCoord{2, 1},
		}
	default:
		panic("invalid coord type")
	}

	for i := 0; i < len(o); i++ {
		c1 := p.Coord().Add(o[i])
		if c1.Out(b) {
			continue
		}
		// check that destination cell isn't contains a piece of same colour
		if dstPiece := b.Cell(c1).Piece(); dstPiece != nil && dstPiece.Colour() == p.Colour() {
			continue
		}

		if excludeCheckExpose && InCheck(p.Project(c1, b), p.Colour()) {
			continue
		}
		d = append(d, c1)
	}

	return NewRectCoords(d)
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Knight) Attacks(b Board) Coords {
	return p.dst(b, false)
}

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Knight) Destinations(b Board) Coords {
	return p.dst(b, true)
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *Knight) Project(to Coord, b Board) Board {
	newBoard := b.Copy()
	newBoard.Empty(p.Coord())
	newBoard.PlacePiece(to, p.Copy())
	return newBoard
}

// Copy a piece
func (p *Knight) Copy() Piece {
	return &Knight{
		BasePiece: p.BasePiece.Copy(),
	}
}
