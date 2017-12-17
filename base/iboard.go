package base

// IBoard
type IBoard interface {
	Dim() Coord
	SetDim(dim Coord)

	Cell(c Coord) *Cell
	Cells() ICells

	Empty(at Coord)
	Copy() IBoard
	Set(b1 IBoard)

	Piece(at Coord) IPiece
	PlacePiece(to Coord, p IPiece)
	MakeMove(to Coord, piece IPiece) bool

	FindPieces(f IPieceFilter) Pieces
	FindAttackedCellsBy(f IPieceFilter) ICoords
}
