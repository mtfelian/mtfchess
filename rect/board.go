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
	return b.cells
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
func (b *Board) PlacePiece(to base.ICoord, p base.IPiece) {
	p.SetCoords(to)
	b.Cell(to).SetPiece(p)
}

// Empty removes piece at coords x, y
func (b *Board) Empty(at base.ICoord) {
	b.Cell(at).Empty()
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
	*b = *b1.Copy().(*Board)
}

// MakeMove makes move with piece to coords (x,y)
// It returns true if move succesful (legal), otherwise it returns false.
func (b *Board) MakeMove(to base.ICoord, piece base.IPiece) bool {
	destinations := piece.Destinations(b)
	for destinations.HasNext() {
		d := destinations.Next().(base.ICoord)
		if to.Equals(d) {
			wasPiece := b.Cell(to).Piece()
			if wasPiece != nil {
				wasPiece.SetCoords(nil)
			}
			newBoard := piece.Project(to, b)
			piece.SetCoords(to)
			b.Set(newBoard)
			return true
		}
	}
	return false
}

// FindPieces finds and returns pieces by filter
func (b *Board) FindPieces(pf base.IPieceFilter) base.Pieces {
	pieces := base.Pieces{}
	f := pf.(PieceFilter)
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
			if len(f.X) > 0 && !SliceContains(p.Coord().(Coord).X, f.X) {
				continue
			}
			if len(f.Y) > 0 && !SliceContains(p.Coord().(Coord).Y, f.Y) {
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

// FindAttackedCellsBy returns a slice of coords of cells attacked by filter of pieces.
// For ex., call b.FindAttackedCells(White) to get cell coords attacked by white pieces.
func (b *Board) FindAttackedCellsBy(f base.IPieceFilter) base.ICoords {
	pieces, pairs := b.FindPieces(f.(PieceFilter)), NewCoords([]base.ICoord{})
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

// todo implement other pieces except knight, to implement EP captures or diag captures like pawns, use move history and board

// NewEmptyBoard creates new empty rectangular board with i cols and j rows
func NewEmptyBoard(i, j int) *Board {
	b := &Board{}
	b.width, b.height = i, j
	b.createCells()
	return b
}
