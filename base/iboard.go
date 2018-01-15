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

	SetSettings(s *Settings)
	Settings() *Settings

	Piece(at ICoord) IPiece
	PlacePiece(to ICoord, p IPiece) IBoard
	MakeMove(to ICoord, piece IPiece) bool

	InCheck(colour Colour) bool
	InCheckMate(colour Colour) bool
	InStaleMate(colour Colour) bool

	Outcome() Outcome
	SetOutcome(to Outcome)

	Castlings(colour Colour) Castlings
	MakeCastling(castling Castling) bool

	SetCanCaptureEnPassantAt(dst ICoord)
	CanCaptureEnPassantAt() ICoord

	SetRookInitialCoords(colour Colour, i int, coord ICoord)
	RookInitialCoords(colour Colour) [2]ICoord

	// Project a piece to coords, returns a pointer to a new copy of a board, don't check legality
	// this don't change coords of a piece
	Project(piece IPiece, to ICoord) IBoard

	King(of Colour) IPiece
	SetKing(of Colour, to IPiece)

	FindPieces(f IPieceFilter) Pieces
	FindAttackedCellsBy(f IPieceFilter) ICoords

	SideToMove() Colour
	SetSideToMove(to Colour)

	MoveNumber() int
	SetMoveNumber(n int)

	HalfMoveCount() int
	SetHalfMoveCount(n int)

	HasMoves(colour Colour) bool
	LegalMoves(notation INotation) []string
}
