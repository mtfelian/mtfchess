package board

// Board
type Board interface {
	Width() int
	Height() int
	SetWidth(width int)
	SetHeight(height int)

	Square(x, y int) *Square
	Squares() Squares

	InCheck(colour Colour) bool

	Empty(x, y int)
	Copy() Board
	Set(b1 Board)

	Piece(x, y int) Piece
	PlacePiece(x, y int, p Piece)
	MakeMove(x, y int, piece Piece) bool
}
