package base

// IBoard
type IBoard interface {
	Dim() ICoord
	SetDim(dim ICoord)

	Cell(c ICoord) *Cell
	Cells() ICells

	Empty(at ICoord) IBoard
	Copy() IBoard
	Set(b1 IBoard)

	Piece(at ICoord) IPiece
	PlacePiece(to ICoord, p IPiece) IBoard
	MakeMove(to ICoord, piece IPiece) bool

	FindPieces(f IPieceFilter) Pieces
	FindAttackedCellsBy(f IPieceFilter) ICoords
}
