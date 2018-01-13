package piece

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
)

// Pawn is a chess pawn
type Pawn struct{ *base.Piece }

// NewPawn creates new pawn with colour
func NewPawn(colour Colour) base.IPiece {
	return &Pawn{Piece: base.NewPiece(colour, base.PawnName, "P♙♟")}
}

// dst returns a slice of destination cells coords, making it's legal moves
// if moving is false then pairs leading to check-exposing moves also included
func (p *Pawn) dst(board base.IBoard, moving bool) base.ICoords {
	switch b := board.(type) {
	case *rect.Board:
		long, pY := 0, p.Coord().(rect.Coord).Y
		if (p.Colour() == White && pY == 2) || (p.Colour() == Black && pY == board.Dim().(rect.Coord).Y-1) {
			long = b.Settings().PawnLongMoveModifier
		}

		d := rect.NewCoords(append(
			reader(1, 0, p, b, moving, 1+long, 1, moveNonCapture),
			leaper(1, 1, p, b, moving, 1, moveCapture)...,
		))

		if moving {
			// search through the possible en passant capturing coords and add if appropriate coords is found
			epCoord := board.Settings().EnPassantFunc(board, p)
			if epCoord != nil {
				d.Add(epCoord)
			}
		}

		return d
	default:
		panic("invalid coord type")
	}
}

// Attacks returns a slice of coords pairs of cells attacked by a piece
func (p *Pawn) Attacks(b base.IBoard) base.ICoords { return p.dst(b, false) }

// Destinations returns a slice of cells coords, making it's legal moves
func (p *Pawn) Destinations(b base.IBoard) base.ICoords { return p.dst(b, true) }

// Copy a piece
func (p *Pawn) Copy() base.IPiece { return &Pawn{Piece: p.Piece.Copy()} }

// Promote returns a promoted piece
func (p *Pawn) Promote() base.IPiece {
	promotion := p.Promotion()
	if promotion == nil {
		return p
	}
	return promotion.Copy()
}

// Set sets a piece to p1
func (p *Pawn) Set(p1 base.IPiece) { *p = *(p1.(*Pawn)) }
