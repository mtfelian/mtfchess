package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Chancellor is a chess chancellor
type Chancellor struct{ *base.Piece }

// NewChancellor creates new chancellor with colour
func NewChancellor(colour Colour) base.IPiece {
	return &Chancellor{Piece: base.NewPiece(colour, base.ChancellorName, "CCc")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Chancellor) dst(board base.IBoard, moving bool) base.ICoords {
	switch b := board.(type) {
	case *rect.Board:
		return rect.NewCoords(append(
			reader(1, 0, p, b, moving, 0, 0, moveAny),
			leaper(1, 2, p, b, moving, 0, moveAny)...,
		))
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Chancellor) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Chancellor) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Copy a piece
func (p *Chancellor) Copy() base.IPiece { return &Chancellor{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *Chancellor) Promote() base.IPiece { return p }

// Set sets a piece to p1
func (p *Chancellor) Set(p1 base.IPiece) { *p = *(p1.(*Chancellor)) }
