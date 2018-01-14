package base

// INotation for writing game moves
type INotation interface {
	// EncodeCoord to string
	EncodeCoord() string

	// EncodeMove on board with piece to dst coord
	EncodeMove(board IBoard, piece IPiece, dst ICoord) string

	// EncodeCastling on board
	EncodeCastling(board IBoard, i int) string

	// DecodeCoord from string
	DecodeCoord(string) error

	// SetCoord sets coord to
	SetCoord(ICoord) INotation
}
