package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Knight is a chess knight
type Knight struct{ *base.Piece }

// NewKnight creates new knight with colour
func NewKnight(colour Colour) base.IPiece {
	return &Knight{Piece: base.NewPiece(colour, "knight", "N♘♞")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Knight) dst(board base.IBoard, moving bool) base.ICoords {
	switch b := board.(type) {
	case *rect.Board:
		return rect.NewCoords(leaper(1, 2, p, b, moving, 0))
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Knight) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Knight) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Copy a piece
func (p *Knight) Copy() base.IPiece { return &Knight{Piece: p.Piece.Copy()} }
