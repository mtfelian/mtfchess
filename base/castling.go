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
	return c.Enabled == castling.Enabled &&
		c.Piece[0].Equals(castling.Piece[0]) &&
		c.Piece[1].Equals(castling.Piece[1]) &&
		c.To[0].Equals(castling.To[0]) &&
		c.To[1].Equals(castling.To[1])
}

// Copy returns a copy of c with pieces taken from board
func (c Castling) Copy(board IBoard) Castling {
	return Castling{
		Piece:   [2]IPiece{board.Piece(c.Piece[0].Coord()), board.Piece(c.Piece[1].Coord())},
		To:      [2]ICoord{c.To[0].Copy(), c.To[1].Copy()},
		Enabled: c.Enabled,
	}
}
