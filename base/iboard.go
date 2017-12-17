package base

// IBoard
type IBoard interface {
	Dim() ICoord
	SetDim(dim ICoord)

	Cell(c ICoord) *Cell
	Cells() ICells

	Empty(at ICoord)
	Copy() IBoard
	Set(b1 IBoard)

	Piece(at ICoord) IPiece
	PlacePiece(to ICoord, p IPiece)
	MakeMove(to ICoord, piece IPiece) bool

	FindPieces(f IPieceFilter) Pieces
	FindAttackedCellsBy(f IPieceFilter) ICoords
}
