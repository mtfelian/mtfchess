package rect

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

// Bishop is a chess bishop
type Bishop struct{ *base.Piece }

// NewBishop creates new bishop with colour
func NewBishop(colour Colour) base.IPiece {
	return &Bishop{Piece: base.NewPiece(colour, base.BishopName, "B♗♝")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Bishop) dst(b *Board, moving bool) base.ICoords {
	return NewCoords(reader(1, 1, p, b, moving, 0, 0, moveAny))
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Bishop) Attacks(b base.IBoard) base.ICoords { return p.dst(b.(*Board), false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Bishop) Destinations(b base.IBoard) base.ICoords { return p.dst(b.(*Board), true) }

// Copy a piece
func (p *Bishop) Copy() base.IPiece { return &Bishop{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *Bishop) Promote() base.IPiece { return p }

// Set sets a piece to p1
func (p *Bishop) Set(p1 base.IPiece) { *p = *(p1.(*Bishop)) }
