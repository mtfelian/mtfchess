package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// King is a chess king
type King struct {
	base.Piece
}

// NewKing creates new king with colour
func NewKing(colour Colour) base.IPiece {
	return &King{
		Piece: base.NewPiece(colour, "king", "K♔♚"),
	}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if excludeCheckExpose is false then pairs leading to check-exposing moves also included
func (p *King) dst(board base.IBoard, excludeCheckExpose bool) base.ICoords {
	o, d := []base.Coord{}, []base.Coord{}

	switch board.Dim().(type) {
	case rect.Coord:
		o = []base.Coord{
			rect.Coord{-1, -1}, rect.Coord{-1, 0}, rect.Coord{-1, 1}, rect.Coord{0, -1},
			rect.Coord{0, 1}, rect.Coord{1, -1}, rect.Coord{1, 0}, rect.Coord{1, 1},
		}
	default:
		panic("invalid coord type")
	}

	for i := 0; i < len(o); i++ {
		to := p.Coord().Add(o[i])
		if to.Out(board) {
			continue
		}
		// check thet destination cell isn't contains a piece of same colour
		if dstPiece := board.Cell(to).Piece(); dstPiece != nil && dstPiece.Colour() == p.Colour() {
			continue
		}

		if excludeCheckExpose && InCheck(p.Project(to, board), p.Colour()) {
			continue
		}
		d = append(d, to)
	}

	return rect.NewCoords(d)
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *King) Attacks(b base.IBoard) base.ICoords {
	return p.dst(b, false)
}

// Destinations returns a slice of cells coords, making it's legal moves
func (p *King) Destinations(b base.IBoard) base.ICoords {
	return p.dst(b, true)
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *King) Project(to base.Coord, b base.IBoard) base.IBoard {
	newBoard := b.Copy()
	newBoard.Empty(to)
	newBoard.PlacePiece(to, p.Copy())
	return newBoard
}

// Copy a piece
func (p *King) Copy() base.IPiece {
	return &Knight{
		Piece: p.Piece.Copy(),
	}
}
