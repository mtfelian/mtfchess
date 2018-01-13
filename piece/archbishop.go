package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Archbishop is a chess archbishop
type Archbishop struct{ *base.Piece }

// NewArchbishop creates new archbishop with colour
func NewArchbishop(colour Colour) base.IPiece {
	return &Archbishop{Piece: base.NewPiece(colour, base.ArchbishopName, "AAa")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Archbishop) dst(board base.IBoard, moving bool) base.ICoords {
	switch b := board.(type) {
	case *rect.Board:
		return rect.NewCoords(append(
			reader(1, 1, p, b, moving, 0, 0, moveAny),
			leaper(1, 2, p, b, moving, 0, moveAny)...,
		))
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Archbishop) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Archbishop) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Copy a piece
func (p *Archbishop) Copy() base.IPiece { return &Archbishop{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *Archbishop) Promote() base.IPiece { return p }

// Set sets a piece to p1
func (p *Archbishop) Set(p1 base.IPiece) { *p = *(p1.(*Archbishop)) }
