package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Queen is a chess queen
type Queen struct {
	base.Piece
}

// NewQueen creates new queen with colour
func NewQueen(colour Colour) base.IPiece {
	return &Queen{
		Piece: base.NewPiece(colour, "queen", "Q♕♛"),
	}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if excludeCheckExpose is false then pairs leading to check-exposing moves also included
func (p *Queen) dst(board base.IBoard, excludeCheckExpose bool) base.ICoords {
	d := []base.ICoord{}

	switch board.Dim().(type) {
	case rect.Coord:
		d = queenRect(p, board, excludeCheckExpose)
	default:
		panic("invalid coord type")
	}

	return rect.NewCoords(d)
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Queen) Attacks(b base.IBoard) base.ICoords {
	return p.dst(b, false)
}

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Queen) Destinations(b base.IBoard) base.ICoords {
	return p.dst(b, true)
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *Queen) Project(to base.ICoord, b base.IBoard) base.IBoard {
	newBoard := b.Copy()
	newBoard.Empty(p.Coord())
	newBoard.PlacePiece(to, p.Copy())
	return newBoard
}

// Copy a piece
func (p *Queen) Copy() base.IPiece {
	return &Queen{
		Piece: p.Piece.Copy(),
	}
}
