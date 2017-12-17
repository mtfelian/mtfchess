package mtfchess

import (
	"fmt"

	. "github.com/mtfelian/mtfchess/board"
	. "github.com/mtfelian/utils"
)

// StdBoard is a game standard board
type StdBoard struct {
	cells         Cells
	width, height int
}

// X converts x1 to slice index
func (b StdBoard) X(x int) int {
	return x - 1
}

// Y convers y1 to slice index
func (b StdBoard) Y(y int) int {
	return b.height - y
}

// String makes Board to implement Stringer
func (b StdBoard) String() string {
	var s string
	for _, row := range b.cells {
		for _, cell := range row {
			s += fmt.Sprintf("%s", cell)
		}
		s += "\n"
	}
	return s
}

// With returns a board width
func (b StdBoard) Width() int {
	return b.width
}

// Height returns a board height
func (b StdBoard) Height() int {
	return b.height
}

// SetHeight sets board height
func (b *StdBoard) SetHeight(height int) {
	b.height = height
}

// SetWidth sets board width
func (b *StdBoard) SetWidth(width int) {
	b.width = width
}

// createCells returns a slice of Cell for the board
func (b *StdBoard) createCells() {
	b.cells = make(Cells, b.height)
	for y := range b.cells {
		b.cells[y] = make(Row, b.width)
		for x := range b.cells[y] {
			b.createCell(x, y)
		}
	}
}

// createCell creates new cell at rectangular board b with coordinates x, y and row length i
func (b *StdBoard) createCell(x, y int) {
	b.cells[y][x] = NewCell(b, b.width*(b.height-y)-b.width+(x+1), x+1, b.height-y)
	b.cells[y][x].Empty()
}

// Cell returns a pointer to cell at coords (x, y)
func (b *StdBoard) Cell(x, y int) *Cell {
	return &b.cells[b.Y(y)][b.X(x)]
}

// Cells returns a cells slice
func (b *StdBoard) Cells() Cells {
	return b.cells
}

// SetCells sets cells to s
func (b *StdBoard) SetCells(s Cells) {
	b.cells = s
}

// Piece returns a piece at coords (x,y)
func (b *StdBoard) Piece(x, y int) Piece {
	return b.Cell(x, y).Piece()
}

// PlacePiece places piece at coords (x, y)
func (b *StdBoard) PlacePiece(x, y int, p Piece) {
	p.SetCoords(x, y)
	b.Cell(x, y).SetPiece(p)
}

// Empty removes piece at coords x, y
func (b *StdBoard) Empty(x, y int) {
	b.Cell(x, y).Empty()
}

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func (b *StdBoard) InCheck(colour Colour) bool {
	return len(b.FindPieces(PieceFilter{
		Colours: []Colour{colour},
		Names:   []string{NewKingPiece(Transparent).Name()},
		Condition: func(p Piece) bool {
			opponentPieces := PieceFilter{Colours: []Colour{colour.Invert()}}
			return SliceContains(Pair{X: p.X(), Y: p.Y()}, b.FindAttackedCellsBy(opponentPieces))
		},
	})) > 0
}

// Copy returns a pointer to a deep copy of a board
func (b *StdBoard) Copy() Board {
	newBoard := &StdBoard{}
	newBoard.SetCells(b.Cells().Copy(newBoard))
	newBoard.SetWidth(b.width)
	newBoard.SetHeight(b.height)
	return newBoard
}

// Set changes b to b1
func (b *StdBoard) Set(b1 Board) {
	b.SetWidth(b1.Width())
	b.SetHeight(b1.Height())
	b.SetCells(b1.Cells())
}

// MakeMove makes move with piece to coords (x,y)
// It returns true if move succesful (legal), otherwise it returns false.
func (b *StdBoard) MakeMove(x, y int, piece Piece) bool {
	destinations := piece.Destinations(b)
	for _, d := range destinations {
		if d.X == x && d.Y == y {
			newBoard := piece.Project(x, y, b)
			piece.SetCoords(x, y)
			b.Set(newBoard)
			return true
		}
	}
	return false
}

// FindPieces finds and returns pieces by filter
func (b *StdBoard) FindPieces(f PieceFilter) Pieces {
	pieces := Pieces{}
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
			if len(f.X) > 0 && !SliceContains(p.X(), f.X) {
				continue
			}
			if len(f.Y) > 0 && !SliceContains(p.Y(), f.Y) {
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
func (b *StdBoard) FindAttackedCellsBy(f PieceFilter) Pairs {
	pieces, pairs := b.FindPieces(f), Pairs{}
	for _, piece := range pieces {
		for _, pair := range piece.Attacks(b) {
			if !SliceContains(pair, pairs) {
				pairs = append(pairs, pair)
			}
		}
	}
	return pairs
}

// todo better testing that Piece.Offsets() don't permit check-exposing moves in situations like (BN's move): WR|BN|BK
// todo implement other pieces except knight, to implement EP captures or diag captures like pawns, use move history and board

// NewEmptyStdBoard creates new empty standard board with i cols and j rows
func NewEmptyStdBoard(i, j int) *StdBoard {
	b := &StdBoard{}
	b.width, b.height = i, j
	b.createCells()
	return b
}
