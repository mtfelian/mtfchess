package rect

import "github.com/mtfelian/mtfchess/base"

// Settings is a game rectangular board settings
type Settings struct {
	// PawnLongFunc's return value added to pawn's move vertical absolute offset (to the front)
	// to allow pawn to move that 1 + number of squares to the front according to this func's logic
	PawnLongFunc func(board base.IBoard, piece base.IPiece) int

	// AllowedPromotions is a list of string piece names to promote to
	AllowedPromotions []string

	// PromotionConditionFunc returns true if piece going to cell dst can be promoted to
	PromotionConditionFunc func(board base.IBoard, piece base.IPiece, dst base.ICoord, to base.IPiece) bool
}
