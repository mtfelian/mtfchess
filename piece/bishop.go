package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Bishop is a chess bishop
type Bishop struct{ *base.Piece }

// NewBishop creates new bishop with colour
func NewBishop(colour Colour) base.IPiece {
	return &Bishop{Piece: base.NewPiece(colour, "bishop", "B♗♝")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if excludeCheckExpose is false then pairs leading to check-exposing moves also included
func (p *Bishop) dst(board base.IBoard, excludeCheckExpose bool) base.ICoords {
	coords := bishop(p, board, excludeCheckExpose)
	switch board.Dim().(type) {
	case rect.Coord:
		return rect.NewCoords(coords)
	default:
		panic("invalid coords type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Bishop) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Bishop) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *Bishop) Project(to base.ICoord, b base.IBoard) base.IBoard {
	return b.Copy().Empty(p.Coord()).PlacePiece(to, p.Copy())
}

// Copy a piece
func (p *Bishop) Copy() base.IPiece { return &Bishop{Piece: p.Piece.Copy()} }
