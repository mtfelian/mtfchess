package rect

import (
	"fmt"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/mtfelian/utils"
)

// Board is a game rectangular board
type Board struct {
	cells                 Cells
	width, height         int
	king                  map[Colour]base.IPiece
	canCaptureEnPassantAt base.ICoord
	rookCoords            base.RookCoords
	settings              base.Settings
}

// X converts x1 to slice index
func (b Board) X(x int) int { return x - 1 }

// Y convers y1 to slice index
func (b Board) Y(y int) int { return b.height - y }

// String makes Board to implement Stringer
func (b Board) String() string {
	var s string
	for i := range b.cells {
		for j := range b.cells[i] {
			s += fmt.Sprintf("%s", b.cells[i][j])
		}
		s += "\n"
	}
	return s
}

// Dim returns a board dimensions
func (b Board) Dim() base.ICoord { return Coord{X: b.width, Y: b.height} }

// SetSettings of a board to s
func (b *Board) SetSettings(s base.Settings) { b.settings = s }

// Settings returns board settings
func (b Board) Settings() base.Settings { return b.settings }

// SetDim sets board dimensions to dim
func (b *Board) SetDim(dim base.ICoord) { b.width, b.height = dim.(Coord).X, dim.(Coord).Y }

// initializeKing initializes board king
func (b *Board) initializeKing() {
	if b.king == nil {
		b.king = map[Colour]base.IPiece{White: nil, Black: nil}
	}
}

// initializeRookCoords initializes board rook coords
func (b *Board) initializeRookCoords() {
	if b.rookCoords == nil {
		b.rookCoords = base.RookCoords{White: [2]base.ICoord{}, Black: [2]base.ICoord{}}
	}
}

// SetKing sets a board king
func (b *Board) SetKing(of Colour, to base.IPiece) {
	b.initializeKing()
	b.king[of] = to
}

// SetCanCaptureEnPassantAt sets a piece dst coords which can be captured en passant
func (b *Board) SetCanCaptureEnPassantAt(dst base.ICoord) { b.canCaptureEnPassantAt = dst }

// CanCaptureEnPassantAt returns a piece dst coords which can be captured en passant
func (b *Board) CanCaptureEnPassantAt() base.ICoord { return b.canCaptureEnPassantAt }

// SetRookInitialCoords sets the rook initial coords
// set parameter i to 0 for most aSide rook and set i to 1 for most zSide rook
func (b *Board) SetRookInitialCoords(colour Colour, i int, coord base.ICoord) {
	if i != 0 && i != 1 {
		panic(fmt.Sprintf("SetRookInitialCoords(): it should be 0 or 1, i: %d", i))
	}
	states := b.rookCoords[colour]
	states[i] = coord.Copy()
	b.rookCoords[colour] = states
}

// RookCanCastle returns true if rook of colour can castle
// set parameter i to 0 for most aSide rook and set i to 1 for most zSide rook
func (b *Board) RookCanCastle(colour Colour, i int) bool {
	if i != 0 && i != 1 {
		panic(fmt.Sprintf("RookCanCastle(): it should be 0 or 1, i: %d", i))
	}
	states := b.rookCoords[colour]
	return states[i] != nil
}

// RookInitialCoords returns rooks coords for castlings as array where
// 0 index means most aSide rook and 1 index means most zSide rook, coord contains nil if no such rook
func (b *Board) RookInitialCoords(colour Colour) [2]base.ICoord {
	return b.rookCoords[colour]
}

// createCells returns a slice of Cell for the board
func (b *Board) createCells() {
	b.cells = make(Cells, b.height)
	for y := range b.cells {
		b.cells[y] = make(Row, b.width)
		for x := range b.cells[y] {
			b.createCell(x, y)
		}
	}
}

// createCell creates new cell at rectangular board b with coordinates x, y and row length i
func (b *Board) createCell(x, y int) {
	b.cells[y][x] = base.NewCell(b, b.width*(b.height-y)-b.width+(x+1), Coord{X: x + 1, Y: b.height - y})
	b.cells[y][x].Empty()
}

// Cell returns a pointer to cell at coords
func (b *Board) Cell(at base.ICoord) *base.Cell {
	c := at.(Coord)
	return &b.cells[b.Y(c.Y)][b.X(c.X)]
}

// Cells returns a cells slice
func (b *Board) Cells() base.ICells { return b.cells }

// SetCells sets cells to s
func (b *Board) SetCells(s base.ICells) { b.cells = s.(Cells) }

// Piece returns a piece at coords
func (b *Board) Piece(at base.ICoord) base.IPiece { return b.Cell(at).Piece() }

// PlacePiece places piece at coords (x, y)
func (b *Board) PlacePiece(to base.ICoord, p base.IPiece) base.IBoard {
	if to.OutOf(b) {
		panic("out of board")
		//return b
	}
	p.SetCoords(b, to)
	b.Cell(to).SetPiece(p)
	return b
}

// Empty removes piece at coords x, y
func (b *Board) Empty(at base.ICoord) base.IBoard {
	piece := b.Cell(at).Piece()
	if piece != nil {
		piece.SetCoords(b, nil)
	}
	b.Cell(at).Empty()
	return b
}

// King returns a king of specified colour
func (b *Board) King(of Colour) base.IPiece { return b.king[of] }

// copyKing returns a copy of a board king
func (b *Board) copyKings() map[Colour]base.IPiece {
	newKing := map[Colour]base.IPiece{}
	for colour := range b.king {
		king := b.King(colour)
		if king != nil {
			newKing[colour] = king.Copy()
		}
	}
	return newKing
}

// Copy returns a pointer to a deep copy of a board
func (b *Board) Copy() base.IBoard {
	newBoard := &Board{}
	newBoard.SetCells(b.Cells().Copy(newBoard))
	newBoard.SetDim(Coord{X: b.width, Y: b.height})
	newBoard.king = b.copyKings()
	newBoard.rookCoords = b.rookCoords.Copy()
	newBoard.SetSettings(b.Settings())
	newBoard.SetCanCaptureEnPassantAt(b.CanCaptureEnPassantAt())
	return newBoard
}

// Set changes b to b1
func (b *Board) Set(b1 base.IBoard) { *b = *(b1.(*Board)) }

// Projects returns a copy of board with projected piece copy to given coords
func (b *Board) Project(piece base.IPiece, to base.ICoord) base.IBoard {
	return b.Copy().Empty(piece.Coord()).PlacePiece(to, piece.Copy())
}

// MakeMove makes move with piece to coords (x,y)
// It returns true if move succesful (legal), otherwise it returns false.
func (b *Board) MakeMove(to base.ICoord, piece base.IPiece) bool {
	if to.OutOf(b) {
		return false
	}
	destinations, capturedPiece := piece.Destinations(b), b.Piece(to)

	if !destinations.Contains(to) {
		return false
	}

	if piece.Promotion() != nil {
		oldCoords := piece.Coord().Copy()
		newPiece := piece.Promote()
		if !b.Settings().PromotionConditionFunc(b, piece, to, newPiece) {
			return false
		}
		piece = newPiece
		b.Empty(oldCoords)
		piece.SetCoords(b, oldCoords)
	}

	if capturedPiece != nil {
		capturedPiece.SetCoords(b, nil)
	}

	if piece.Name() == "pawn" {
		epCaptureAt := b.CanCaptureEnPassantAt()
		if epCaptureAt != nil && to.(Coord).X == epCaptureAt.(Coord).X {
			b.Empty(epCaptureAt)
		}

		b.SetCanCaptureEnPassantAt(nil)
		pY, toY := piece.Coord().(Coord).Y, to.(Coord).Y
		diff := pY - toY
		if diff != 1 && diff != -1 { // long pawn move
			b.SetCanCaptureEnPassantAt(to)
		}
	}

	piece.MarkMoved()
	b.Set(b.Project(piece, to))
	// first project (and empty source piece square, and only then set piece)
	piece.Set(b.Piece(to)) // set piece to copy of itself on the new board

	return true
}

// MakeCastling makes a castling.
// It returns true if castling succesful (legal), otherwise it returns false.
func (b *Board) MakeCastling(castling base.Castling) bool {
	castlings := b.Castlings(castling.Piece[0].Colour())

	if !castlings.Contains(castling) {
		return false
	}

	castling.Piece[0].MarkMoved()
	castling.Piece[1].MarkMoved()

	kingCopy, rookCopy := b.Piece(castling.Piece[0].Coord()).Copy(), b.Piece(castling.Piece[1].Coord()).Copy()
	b.Empty(kingCopy.Coord())
	b.Empty(rookCopy.Coord())
	b.PlacePiece(castling.To[0], kingCopy)
	b.PlacePiece(castling.To[1], rookCopy)
	castling.Piece[0].Set(b.Piece(castling.To[0]))
	castling.Piece[1].Set(b.Piece(castling.To[1]))

	return true
}

// baseFindPieces finds and returns pieces by base.PieceFilter
func (b *Board) baseFindPieces(f base.PieceFilter) base.Pieces {
	pieces := base.Pieces{}
	for i := range b.cells {
		for j := range b.cells[i] {
			p := b.cells[i][j].Piece()
			if p == nil {
				continue
			}
			if len(f.Colours) > 0 && !SliceContains(p.Colour(), f.Colours) {
				continue
			}
			if len(f.Names) > 0 && !SliceContains(p.Name(), f.Names) {
				continue
			}
			if f.Condition != nil && !f.Condition(p) {
				continue
			}
			pieces = append(pieces, p)
		}
	}
	return pieces
}

// FindPieces finds and returns pieces by base.PieceFilter or rect.PieceFilter
func (b *Board) FindPieces(pf base.IPieceFilter) base.Pieces {
	pieces := base.Pieces{}
	switch filter := pf.(type) {
	case base.PieceFilter:
		return b.baseFindPieces(filter)
	case PieceFilter:
		pieces = b.baseFindPieces(filter.PieceFilter)
	}

	f := pf.(PieceFilter)
	r := base.Pieces{}
	for i := range pieces {
		if len(f.X) > 0 && !SliceContains(pieces[i].Coord().(Coord).X, f.X) {
			continue
		}
		if len(f.Y) > 0 && !SliceContains(pieces[i].Coord().(Coord).Y, f.Y) {
			continue
		}
		r = append(r, pieces[i])
	}
	return r
}

// FindAttackedCellsBy returns a slice of coords of cells attacked by filter of pieces.
// For ex., call b.FindAttackedCells(White) to get cell coords attacked by white pieces.
func (b *Board) FindAttackedCellsBy(f base.IPieceFilter) base.ICoords {
	pieces, pairs := b.FindPieces(f), NewCoords([]base.ICoord{})
	for i := range pieces {
		attackedCoords := pieces[i].Attacks(b)
		for attackedCoords.HasNext() {
			pair := attackedCoords.Next().(base.ICoord)
			if !pairs.Contains(pair) {
				pairs.Add(pair)
			}
		}
	}
	return pairs
}

// Equals returns true if two boards are equal
func (b *Board) Equals(to base.IBoard) bool {
	b1 := to.(*Board)
	if b.width != b1.width || b.height != b1.height {
		return false
	}
	for y := 1; y <= b.height; y++ {
		for x := 1; x <= b.width; x++ {
			p1, p2 := b.Cell(Coord{x, y}).Piece(), b1.Cell(Coord{x, y}).Piece()
			if (p1 == nil && p2 != nil) || (p1 != nil && p2 == nil) {
				return false
			}
			if p1 != nil && p2 != nil && !p1.Equals(p2) {
				return false
			}
		}
	}
	return true
}

// Castlings returns available castlings for colour
func (b *Board) Castlings(colour Colour) base.Castlings { return b.Settings().CastlingsFunc(b, colour) }

// InChecks returns true if king of colour is in check
func (b *Board) InCheck(colour Colour) bool {
	king := b.King(colour)
	return king != nil && b.FindAttackedCellsBy(base.PieceFilter{Colours: []Colour{colour.Invert()}}).Contains(king.Coord())
}

// NewEmptyStandardChessBoard creates new empty board for standard chess
func NewEmptyStandardChessBoard() *Board { return NewEmptyBoard(8, 8, StandardChessBoardSettings()) }

// NewEmptyTestBoard creates new empty board for tests
func NewEmptyTestBoard() *Board { return NewEmptyBoard(5, 6, testBoardSettings()) }

// NewEmptyBoard creates new empty rectangular board with i cols and j rows
func NewEmptyBoard(i, j int, settings base.Settings) *Board {
	b := &Board{}
	b.width, b.height = i, j
	b.createCells()
	b.initializeKing()
	b.initializeRookCoords()
	b.settings = settings
	return b
}

/*
todo to implement:
  - archbishop piece;
  - chancellor piece;
  - with board options:
    - 3-fold repetition draw rule;
    - 50 moves draw rule;
  - returning legal moves in algebraic notation;
  - stalemate detection (no check and no legal moves);
  - checkmate detection (check and no legal moves);
*/