package base

// BasePieceFilter
type BasePieceFilter struct {
	Names     []string
	Colours   []Colour
	Condition func(Piece) bool
}
