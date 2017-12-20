package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Rook is a chess rook
type Rook struct{ *base.Piece }

// NewRook creates new rook with colour
func NewRook(colour Colour) base.IPiece { return &Rook{Piece: base.NewPiece(colour, "rook", "R♖♜")} }

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Rook) dst(board base.IBoard, moving bool) base.ICoords {
	switch b := board.(type) {
	case *rect.Board:
		return rect.NewCoords(reader(1, 0, p, b, moving, 0, 0))
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Rook) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Rook) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *Rook) Project(to base.ICoord, b base.IBoard) base.IBoard {
	return b.Copy().Empty(p.Coord()).PlacePiece(to, p.Copy())
}

// Copy a piece
func (p *Rook) Copy() base.IPiece { return &Rook{Piece: p.Piece.Copy()} }
