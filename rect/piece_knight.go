package rect

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

// Knight is a chess knight
type Knight struct{ *base.Piece }

// NewKnight creates new knight with colour
func NewKnight(colour Colour) base.IPiece {
	return &Knight{Piece: base.NewPiece(colour, base.KnightName, "N♘♞")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Knight) dst(b *Board, moving bool) base.ICoords {
	return NewCoords(leaper(1, 2, p, b, moving, 0, moveAny))
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Knight) Attacks(b base.IBoard) base.ICoords { return p.dst(b.(*Board), false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Knight) Destinations(b base.IBoard) base.ICoords { return p.dst(b.(*Board), true) }

// Copy a piece
func (p *Knight) Copy() base.IPiece { return &Knight{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *Knight) Promote() base.IPiece { return p }

// Set sets a piece to p1
func (p *Knight) Set(p1 base.IPiece) { *p = *(p1.(*Knight)) }
