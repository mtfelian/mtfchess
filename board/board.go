package board

// Board
type Board interface {
	Width() int
	Height() int
	SetWidth(width int)
	SetHeight(height int)

	Cell(x, y int) *Cell
	Cells() Cells

	InCheck(colour Colour) bool

	Empty(x, y int)
	Copy() Board
	Set(b1 Board)

	Piece(x, y int) Piece
	PlacePiece(x, y int, p Piece)
	MakeMove(x, y int, piece Piece) bool

	FindPieces(f PieceFilter) Pieces
	FindAttackedCellsBy(f PieceFilter) Pairs
}
