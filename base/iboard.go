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

	Equals(to IBoard) bool

	Empty(at ICoord) IBoard
	Copy() IBoard
	Set(b1 IBoard)

	SetSettings(s Settings)
	Settings() Settings

	Piece(at ICoord) IPiece
	PlacePiece(to ICoord, p IPiece) IBoard
	MakeMove(to ICoord, piece IPiece) bool

	InCheck(colour Colour) bool
	Castlings(colour Colour) Castlings
	MakeCastling(castling Castling) bool

	SetCanCaptureEnPassant(p IPiece)
	CanCaptureEnPassant() IPiece

	// Project a piece to coords, returns a pointer to a new copy of a board, don't check legality
	// this don't change coords of a piece
	Project(piece IPiece, to ICoord) IBoard

	King(of Colour) IPiece
	SetKing(of Colour, to IPiece)

	FindPieces(f IPieceFilter) Pieces
	FindAttackedCellsBy(f IPieceFilter) ICoords
}
