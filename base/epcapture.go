package base

// EPCapture
type EPCapture struct {
	// From and To is a source and dst coords of a piece which can be captured by EP
	From, To ICoord
}

// Copy returns a copy of c with pieces taken from board
func (c EPCapture) Copy(board IBoard) EPCapture {
	return EPCapture{From: c.From.Copy(), To: c.To.Copy()}
}
