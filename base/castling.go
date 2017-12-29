package base

// Castling
type Castling struct {
	// Piece, [0] - king piece in standard chess, [1] - rook piece in std chess
	Piece [2]IPiece
	// To, [0] - dst coord of king piece in standard chess, [1] - dst coord of rook piece in std chess
	To [2]ICoord
	// Enabled is true if castling is possible
	Enabled bool
}

// Equal returns true if c equals to castling
func (c Castling) Equal(castling Castling) bool {
	eq := c.Enabled == castling.Enabled
	eq = eq && c.Piece[0].Equals(castling.Piece[0])
	eq = eq && c.Piece[1].Equals(castling.Piece[1])
	eq = eq && c.To[0].Equals(castling.To[0])
	eq = eq && c.To[1].Equals(castling.To[1])
	return eq
}
