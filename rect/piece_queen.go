package rect

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

// Queen is a chess queen
type Queen struct{ *base.Piece }

// NewQueen creates new queen with colour
func NewQueen(colour Colour) base.IPiece {
	return &Queen{Piece: base.NewPiece(colour, base.QueenName, "Q♕♛")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Queen) dst(b *Board, moving bool) base.ICoords {
	return NewCoords(append(
		reader(1, 0, p, b, moving, 0, 0, moveAny),
		reader(1, 1, p, b, moving, 0, 0, moveAny)...,
	))
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Queen) Attacks(b base.IBoard) base.ICoords { return p.dst(b.(*Board), false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Queen) Destinations(b base.IBoard) base.ICoords { return p.dst(b.(*Board), true) }

// Copy a piece
func (p *Queen) Copy() base.IPiece { return &Queen{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *Queen) Promote() base.IPiece { return p }

// Set sets a piece to p1
func (p *Queen) Set(p1 base.IPiece) { *p = *(p1.(*Queen)) }
