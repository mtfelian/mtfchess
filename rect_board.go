package mtfchess

import (
	"fmt"

	. "github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/utils"
)

// RectBoard is a game rectangular board
type RectBoard struct {
	cells         RectCells
	width, height int
}

// X converts x1 to slice index
func (b RectBoard) X(x int) int {
	return x - 1
}

// Y convers y1 to slice index
func (b RectBoard) Y(y int) int {
	return b.height - y
}

// String makes Board to implement Stringer
func (b RectBoard) String() string {
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
func (b RectBoard) Dim() Coord {
	return RectCoord{X: b.width, Y: b.height}
}

// SetDim sets board dimensions to dim
func (b *RectBoard) SetDim(dim Coord) {
	b.width, b.height = dim.(RectCoord).X, dim.(RectCoord).Y
}

// createCells returns a slice of Cell for the board
func (b *RectBoard) createCells() {
	b.cells = make(RectCells, b.height)
	for y := range b.cells {
		b.cells[y] = make(RectRow, b.width)
		for x := range b.cells[y] {
			b.createCell(x, y)
		}
	}
}

// createCell creates new cell at rectangular board b with coordinates x, y and row length i
func (b *RectBoard) createCell(x, y int) {
	b.cells[y][x] = NewCell(b, b.width*(b.height-y)-b.width+(x+1), RectCoord{X: x + 1, Y: b.height - y})
	b.cells[y][x].Empty()
}

// Cell returns a pointer to cell at coords
func (b *RectBoard) Cell(at Coord) *Cell {
	c := at.(RectCoord)
	return &b.cells[b.Y(c.Y)][b.X(c.X)]
}

// Cells returns a cells slice
func (b *RectBoard) Cells() Cells {
	return b.cells
}

// SetCells sets cells to s
func (b *RectBoard) SetCells(s Cells) {
	b.cells = s.(RectCells)
}

// Piece returns a piece at coords
func (b *RectBoard) Piece(at Coord) Piece {
	return b.Cell(at).Piece()
}

// PlacePiece places piece at coords (x, y)
func (b *RectBoard) PlacePiece(to Coord, p Piece) {
	p.SetCoords(to)
	b.Cell(to).SetPiece(p)
}

// Empty removes piece at coords x, y
func (b *RectBoard) Empty(at Coord) {
	b.Cell(at).Empty()
}

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func (b *RectBoard) InCheck(colour Colour) bool {
	return len(b.FindPieces(RectPieceFilter{
		BasePieceFilter: BasePieceFilter{
			Colours: []Colour{colour},
			Names:   []string{NewKingPiece(Transparent).Name()},
			Condition: func(p Piece) bool {
				opponentPieces := RectPieceFilter{
					BasePieceFilter: BasePieceFilter{Colours: []Colour{colour.Invert()}},
				}
				return b.FindAttackedCellsBy(opponentPieces).Contains(p.Coord())
			},
		},
	})) > 0
}

// Copy returns a pointer to a deep copy of a board
func (b *RectBoard) Copy() Board {
	newBoard := &RectBoard{}
	newBoard.SetCells(b.Cells().Copy(newBoard))
	newBoard.SetDim(RectCoord{X: b.width, Y: b.height})
	return newBoard
}

// Set changes b to b1
func (b *RectBoard) Set(b1 Board) {
	b.SetDim(b1.Dim())
	b.SetCells(b1.Cells())
}

// MakeMove makes move with piece to coords (x,y)
// It returns true if move succesful (legal), otherwise it returns false.
func (b *RectBoard) MakeMove(to Coord, piece Piece) bool {
	destinations := piece.Destinations(b)
	for destinations.HasNext() {
		d := destinations.Next().(Coord)
		if to.Equals(d) {
			newBoard := piece.Project(to, b)
			piece.SetCoords(to)
			b.Set(newBoard)
			return true
		}
	}
	return false
}

// FindPieces finds and returns pieces by filter
func (b *RectBoard) FindPieces(pf PieceFilter) Pieces {
	pieces := Pieces{}
	f := pf.(RectPieceFilter)
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
			if len(f.X) > 0 && !SliceContains(p.Coord().(RectCoord).X, f.X) {
				continue
			}
			if len(f.Y) > 0 && !SliceContains(p.Coord().(RectCoord).Y, f.Y) {
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
func (b *RectBoard) FindAttackedCellsBy(f PieceFilter) Coords {
	pieces, pairs := b.FindPieces(f.(RectPieceFilter)), NewRectCoords([]Coord{})
	for _, piece := range pieces {
		attackedCoords := piece.Attacks(b)
		for attackedCoords.HasNext() {
			pair := attackedCoords.Next().(Coord)
			if !pairs.Contains(pair) {
				pairs.Add(pair)
			}
		}
	}
	return pairs
}

// todo better testing that Piece.Offsets() don't permit check-exposing moves in situations like (BN's move): WR|BN|BK
// todo implement other pieces except knight, to implement EP captures or diag captures like pawns, use move history and board

// NewEmptyRectBoard creates new empty rectangular board with i cols and j rows
func NewEmptyRectBoard(i, j int) *RectBoard {
	b := &RectBoard{}
	b.width, b.height = i, j
	b.createCells()
	return b
}
