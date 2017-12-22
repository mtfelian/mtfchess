package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Queen is a chess queen
type Queen struct{ *base.Piece }

// NewQueen creates new queen with colour
func NewQueen(colour Colour) base.IPiece {
	return &Queen{Piece: base.NewPiece(colour, "queen", "Q♕♛")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Queen) dst(board base.IBoard, moving bool) base.ICoords {
	switch b := board.(type) {
	case *rect.Board:
		return rect.NewCoords(append(
			reader(1, 0, p, b, moving, 0, 0, moveAny),
			reader(1, 1, p, b, moving, 0, 0, moveAny)...,
		))
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Queen) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Queen) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Copy a piece
func (p *Queen) Copy() base.IPiece { return &Queen{Piece: p.Piece.Copy()} }
