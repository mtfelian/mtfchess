package base

// INotation for writing game moves
type INotation interface {
	// EncodeCoord to string
	EncodeCoord() string

	// EncodeMove on board with piece to dst coord
	EncodeMove(IBoard, IPiece, ICoord) string

	// EncodeCastling
	EncodeCastling(i int) string

	// DecodeCoord from string
	DecodeCoord(string) error

	// DecodeMove from string into making move func, see IBoard.MakeMove()
	DecodeMove(IBoard, string) (func() bool, error)

	// SetCoord sets coord to
	SetCoord(ICoord) INotation
}
