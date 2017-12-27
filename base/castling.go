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
