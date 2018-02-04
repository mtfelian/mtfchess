package rect

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

// Rook is a chess rook
type Rook struct{ *base.Piece }

// NewRook creates new rook with colour
func NewRook(colour Colour) base.IPiece {
	return &Rook{Piece: base.NewPiece(colour, base.RookName, "R♖♜")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Rook) dst(b *Board, moving bool) base.ICoords {
	return NewCoords(reader(1, 0, p, b, moving, 0, 0, moveAny))
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Rook) Attacks(b base.IBoard) base.ICoords { return p.dst(b.(*Board), false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Rook) Destinations(b base.IBoard) base.ICoords { return p.dst(b.(*Board), true) }

// Copy a piece
func (p *Rook) Copy() base.IPiece { return &Rook{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *Rook) Promote() base.IPiece { return p }

// Set sets a piece to p1
func (p *Rook) Set(p1 base.IPiece) { *p = *(p1.(*Rook)) }
