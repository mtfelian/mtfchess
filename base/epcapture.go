package base

// EPCapture
type EPCapture struct {
	From, To ICoord
	Piece    IPiece
}

// Copy returns a copy of c with pieces taken from board
func (c EPCapture) Copy(board IBoard) EPCapture {
	return EPCapture{From: c.From.Copy(), To: c.To.Copy(), Piece: c.Piece.Copy()}
}
