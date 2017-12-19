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
// if excludeCheckExpose is false then pairs leading to check-exposing moves also included
func (p *Knight) dst(board base.IBoard, excludeCheckExpose bool) base.ICoords {
	coords := knight(p, board, excludeCheckExpose)
	switch board.Dim().(type) {
	case rect.Coord:
		return rect.NewCoords(coords)
	default:
		panic("invalid coords type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Knight) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Knight) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Project a copy of a piece to the specified coords on board, return a copy of a board
func (p *Knight) Project(to base.ICoord, b base.IBoard) base.IBoard {
	return b.Copy().Empty(p.Coord()).PlacePiece(to, p.Copy())
}

// Copy a piece
func (p *Knight) Copy() base.IPiece { return &Knight{Piece: p.Piece.Copy()} }
