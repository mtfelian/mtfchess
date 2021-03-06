package base

import (
	. "github.com/mtfelian/mtfchess/colour"
)

// Settings is a game rectangular board settings
type Settings struct {
	// PawnLongMoveModifier's added to pawn's move vertical absolute offset (to the front)
	// to allow pawn to move that 1 + number of squares to the front according to this func's logic
	PawnLongMoveModifier int

	// AllowedPromotions is a list of string piece names to promote to
	AllowedPromotions []string

	// PromotionConditionFunc returns true if piece going to cell dst can be promoted to
	PromotionConditionFunc func(board IBoard, piece IPiece, dst ICoord, to IPiece) bool

	// EnPassantFunc returns coords on which a piece can do en passant capturing
	EnPassantFunc func(board IBoard, piece IPiece) ICoord

	// CastlingsFunc returns available castlings for the given colour
	CastlingsFunc func(board IBoard, colour Colour) Castlings

	// MoveOrder enables move order control if set to true
	MoveOrder bool

	// MovesToDraw specifies amount of moves without capture or pawn advance to declare a draw
	MovesToDraw int

	// PositionsToDraw specifies amount of positions repetition to declare a draw
	PositionsToDraw int
}
