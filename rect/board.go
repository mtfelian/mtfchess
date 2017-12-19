package rect

import (
	"fmt"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/utils"
)

// Board is a game rectangular board
type Board struct {
	cells         Cells
	width, height int
}

// X converts x1 to slice index
func (b Board) X(x int) int {
	return x - 1
}

// Y convers y1 to slice index
func (b Board) Y(y int) int {
	return b.height - y
}

// String makes Board to implement Stringer
func (b Board) String() string {
	var s string
	for _, row := range b.cells {
		for _, cell := range row {
			s += fmt.Sprintf("%s", cell)
		}
		s += "\n"
	}
	return s
}

// Dim returns a board dimensions
func (b Board) Dim() base.ICoord {
	return Coord{X: b.width, Y: b.height}
}

// SetDim sets board dimensions to dim
func (b *Board) SetDim(dim base.ICoord) {
	b.width, b.height = dim.(Coord).X, dim.(Coord).Y
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
func (b *Board) Cells() base.ICells {
	return b.cells.Copy(b)
}

// SetCells sets cells to s
func (b *Board) SetCells(s base.ICells) {
	b.cells = s.(Cells)
}

// Piece returns a piece at coords
func (b *Board) Piece(at base.ICoord) base.IPiece {
	return b.Cell(at).Piece()
}

// PlacePiece places piece at coords (x, y)
func (b *Board) PlacePiece(to base.ICoord, p base.IPiece) base.IBoard {
	p.SetCoords(to)
	b.Cell(to).SetPiece(p)
	return b
}

// Empty removes piece at coords x, y
func (b *Board) Empty(at base.ICoord) base.IBoard {
	b.Cell(at).Empty()
	return b
}

// Copy returns a pointer to a deep copy of a board
func (b *Board) Copy() base.IBoard {
	newBoard := &Board{}
	newBoard.SetCells(b.Cells().Copy(newBoard))
	newBoard.SetDim(Coord{X: b.width, Y: b.height})
	return newBoard
}

// Set changes b to b1
func (b *Board) Set(b1 base.IBoard) {
	b.SetDim(b1.Dim())
	b.SetCells(b1.Cells())
}

// MakeMove makes move with piece to coords (x,y)
// It returns true if move succesful (legal), otherwise it returns false.
func (b *Board) MakeMove(to base.ICoord, piece base.IPiece) bool {
	destinations := piece.Destinations(b)
	for destinations.HasNext() {
		d := destinations.Next().(base.ICoord)
		if !to.Equals(d) {
			continue
		}
		wasPiece := b.Cell(to).Piece()
		if wasPiece != nil {
			wasPiece.SetCoords(nil)
		}
		newBoard := piece.Project(to, b)
		piece.SetCoords(to)
		b.Set(newBoard)
		return true
	}
	return false
}

// baseFindPieces finds and returns pieces by base.PieceFilter
func (b *Board) baseFindPieces(f base.PieceFilter) base.Pieces {
	pieces := base.Pieces{}
	for _, row := range b.cells {
		for _, cell := range row {
			p := cell.Piece()
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
	for _, p := range pieces {
		if len(f.X) > 0 && !SliceContains(p.Coord().(Coord).X, f.X) {
			continue
		}
		if len(f.Y) > 0 && !SliceContains(p.Coord().(Coord).Y, f.Y) {
			continue
		}
		r = append(r, p)
	}
	return r
}

// FindAttackedCellsBy returns a slice of coords of cells attacked by filter of pieces.
// For ex., call b.FindAttackedCells(White) to get cell coords attacked by white pieces.
func (b *Board) FindAttackedCellsBy(f base.IPieceFilter) base.ICoords {
	pieces, pairs := b.FindPieces(f), NewCoords([]base.ICoord{})
	for _, piece := range pieces {
		attackedCoords := piece.Attacks(b)
		for attackedCoords.HasNext() {
			pair := attackedCoords.Next().(base.ICoord)
			if !pairs.Contains(pair) {
				pairs.Add(pair)
			}
		}
	}
	return pairs
}

// NewEmptyBoard creates new empty rectangular board with i cols and j rows
func NewEmptyBoard(i, j int) *Board {
	b := &Board{}
	b.width, b.height = i, j
	b.createCells()
	return b
}

/*
todo to implement:
  - directions_rect.go: front(), frontDiag() based on piece colour
  - pawn piece (one cell front, capture by one front diag);
  - archbishop piece;
  - chancellor piece;
  - with board options:
    - pawn moving by two cells;
    - EP, when a pawn goes by two cells, if at end of that move there is an opponent's pawn(s) by the side of the
      destination cell, memorize side pawns' coords in board, and clean it at next move (can set again to different
      pawns at the end of the same move);
      - rework InCheck() detection by simply keeping kings' coords always in board like with pawns' EP;
    - castling (some work with rooks memorizing like in EP and side cells check);
    - pawn promotion (exchanging one piece with another) to one from the list of allowed pieces;
    - 3-fold repetition draw rule;
    - 50 moves draw rule;
  - returning legal moves in algebraic notation;
  - stalemate detection (no check and no legal moves);
  - checkmate detection (check and no legal moves);
*/
