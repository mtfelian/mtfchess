package base

// Board
type Board interface {
	Dim() Coord
	SetDim(dim Coord)

	Cell(c Coord) *Cell
	Cells() Cells

	InCheck(colour Colour) bool

	Empty(at Coord)
	Copy() Board
	Set(b1 Board)

	Piece(at Coord) Piece
	PlacePiece(to Coord, p Piece)
	MakeMove(to Coord, piece Piece) bool

	FindPieces(f PieceFilter) Pieces
	FindAttackedCellsBy(f PieceFilter) Coords
}
