package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// King is a chess king
type King struct{ *base.Piece }

// NewKing creates new king with colour
func NewKing(colour Colour) base.IPiece {
	return &King{Piece: base.NewPiece(colour, "king", "K♔♚")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *King) dst(board base.IBoard, moving bool) base.ICoords {
	switch b := board.(type) {
	case *rect.Board:
		return rect.NewCoords(append(leaper(1, 0, p, b, moving, 0, moveAny), leaper(1, 1, p, b, moving, 0, moveAny)...))
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *King) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *King) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// SetCoords sets piece's coords to
func (p *King) SetCoords(board base.IBoard, to base.ICoord) {
	p.Piece.SetCoords(board, to)
	board.SetKing(p.Colour(), p) // when king moves, set it to a board for faster check detection
}

// Copy a piece
func (p *King) Copy() base.IPiece { return &King{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *King) Promote() base.IPiece { return p }
