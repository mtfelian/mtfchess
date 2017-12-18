package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Rook is a chess rook
type Rook struct {
	base.Piece
}

// NewRook creates new rook with colour
func NewRook(colour Colour) base.IPiece {
	return &Rook{
		Piece: base.NewPiece(colour, "rook", "R♖♜"),
	}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if excludeCheckExpose is false then pairs leading to check-exposing moves also included
func (p *Rook) dst(board base.IBoard, excludeCheckExpose bool) base.ICoords {
	d := []base.ICoord{}

	switch board.Dim().(type) {
	case rect.Coord:
		d = rook(p, board, excludeCheckExpose)
	default:
		panic("invalid coord type")
	}

	return rect.NewCoords(d)
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Rook) Attacks(b base.IBoard) base.ICoords {
	return p.dst(b, false)
}

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Rook) Destinations(b base.IBoard) base.ICoords {
	return p.dst(b, true)
}

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *Rook) Project(to base.ICoord, b base.IBoard) base.IBoard {
	newBoard := b.Copy()
	newBoard.Empty(p.Coord())
	newBoard.PlacePiece(to, p.Copy())
	return newBoard
}

// Copy a piece
func (p *Rook) Copy() base.IPiece {
	return &Rook{
		Piece: p.Piece.Copy(),
	}
}
