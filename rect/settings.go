package rect

// Settings is a game rectangular board settings
type Settings struct {
	// PawnLongModifier added to pawn's move vertical absolute offset (to the front)
	// to allow pawn to move PawnLongModifier+1 squares to the front from
	// 2nd horizontal for White and from (board.height-1)th horizontal for Black
	PawnLongModifier int

	// AllowedPromotions is a list of string piece names to promote to
	AllowedPromotions []string
}
