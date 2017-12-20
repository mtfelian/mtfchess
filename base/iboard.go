package base

import (
	. "github.com/mtfelian/mtfchess/colour"
)

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

	King(of Colour) IPiece
	SetKing(of Colour, to IPiece)

	FindPieces(f IPieceFilter) Pieces
	FindAttackedCellsBy(f IPieceFilter) ICoords
}
