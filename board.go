package mtfchess

import "fmt"

// Board is a game board
type Board struct {
	cells [][]Cell
}

// String makes Board to implement stringer
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

// newCells returns a slice of Cell sized i (cols) x j (rows)
func (b *Board) newCells(i, j int) {
	b.cells = make([][]Cell, j)
	for y := range b.cells {
		b.cells[y] = make([]Cell, i)
		for x := range b.cells[y] {
			//fmt.Println("new cell call", x, y, i)
			b.newCell(x, y, i)
		}
	}
}

// newCell creates new cell at rectangle board b with coordinates x, y and row length i
func (b *Board) newCell(x, y, i int) {
	fmt.Println(y, x)
	b.cells[y][x] = Cell{
		board: b,
		num:   i*(i-y+1) - i + (x + 1),
		x:     x + 1,
		y:     i - y + 1,
	}
}

// NewEmptyBoard creates new empty board
func NewEmptyBoard(i, j int) Board {
	b := Board{}
	b.newCells(i, j)
	return b
}
