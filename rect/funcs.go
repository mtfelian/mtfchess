package rect

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/mtfelian/utils"
)

// NoPawnLongMoveFunc always disable pawn long forward move
func NoPawnLongMoveFunc(board base.IBoard, piece base.IPiece) int { return 0 }

// StandardPawnLongMoveFunc is a condition for pawn long forward move for standard chess
func StandardPawnLongMoveFunc(board base.IBoard, piece base.IPiece) int {
	bh := board.Dim().(Coord).Y
	if (piece.Colour() == White && piece.Coord().(Coord).Y == 2) ||
		(piece.Colour() == Black && piece.Coord().(Coord).Y == bh-1) {
		return 1
	}
	return 0
}

// StandardAllowedPromotions returns allowed pawn promotions pieces names list for standard chess
func StandardAllowedPromotions() []string { return []string{"knight", "bishop", "rook", "queen"} }

// StandardPromotionConditionFunc is a pawn promotion condition for standard chess
func StandardPromotionConditionFunc(board base.IBoard, piece base.IPiece, dst base.ICoord, to base.IPiece) bool {
	bh, fromY, dstY := board.(*Board).height, piece.Coord().(Coord).Y, dst.(Coord).Y
	return piece.Name() == "pawn" && // only pawn can be promoted
		to.Colour() == piece.Colour() && // only to self-colour
		SliceContains(to.Name(), board.(*Board).Settings().AllowedPromotions) && // to piece from list
		(piece.Colour() == White && fromY == bh-1 && dstY == bh || // for white from pre-last horizontal to the last
			piece.Colour() == Black && fromY == 2 && dstY == 1) // for black from 2nd horizontal to the 1st
}
