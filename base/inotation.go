package base

// INotation for writing game moves
type INotation interface {
	// Encode to string
	Encode() string

	// EncodeMove on board with piece to dst coord
	EncodeMove(board IBoard, piece IPiece, dst ICoord) string

	// Decode from string
	Decode(string) error

	// SetCoord sets coord to
	SetCoord(ICoord) INotation
}
