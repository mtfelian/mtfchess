package base

import . "github.com/mtfelian/mtfchess/colour"

// INotation for writing game moves
type INotation interface {
	// EncodeCoord to string
	EncodeCoord() string

	// EncodeMove on board with piece to dst coord
	EncodeMove(board IBoard, piece IPiece, dst ICoord) string

	// EncodeCastling on board for sideToMove
	EncodeCastling(board IBoard, sideToMove Colour, i int) string

	// DecodeCoord from string
	DecodeCoord(string) error

	// SetCoord sets coord to
	SetCoord(ICoord) INotation
}
