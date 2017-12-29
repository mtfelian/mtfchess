package base

import (
	. "github.com/mtfelian/mtfchess/colour"
)

// Settings is a game rectangular board settings
type Settings struct {
	// PawnLongFunc's return value added to pawn's move vertical absolute offset (to the front)
	// to allow pawn to move that 1 + number of squares to the front according to this func's logic
	PawnLongFunc func(board IBoard, piece IPiece) int

	// AllowedPromotions is a list of string piece names to promote to
	AllowedPromotions []string

	// PromotionConditionFunc returns true if piece going to cell dst can be promoted to
	PromotionConditionFunc func(board IBoard, piece IPiece, dst ICoord, to IPiece) bool

	// CastlingsFunc returns available castlings for the given colour
	CastlingsFunc func(board IBoard, colour Colour) Castlings
}
