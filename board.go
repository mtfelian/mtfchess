package mtfchess

import "fmt"

// Board is a game board
type Board struct {
	squares       [][]Square
	width, height int
}

// String makes Board to implement Stringer
func (b Board) String() string {
	var s string
	for _, row := range b.squares {
		for _, square := range row {
			s += fmt.Sprintf("%s", square)
		}
		s += "\n"
	}
	return s
}

// With returns a board width
func (b Board) Width() int {
	return b.width
}

// Height returns a board height
func (b Board) Height() int {
	return b.height
}

// createSquares returns a slice of Square for the board
func (b *Board) createSquares() {
	b.squares = make([][]Square, b.height)
	for y := range b.squares {
		b.squares[y] = make([]Square, b.width)
		for x := range b.squares[y] {
			b.createSquare(x, y)
		}
	}
}

// createSquare creates new square at rectangular board b with coordinates x, y and row length i
func (b *Board) createSquare(x, y int) {
	b.squares[y][x] = Square{
		board: b,
		num:   b.width*(b.height-y) - b.width + (x + 1),
		x:     x + 1,
		y:     b.height - y,
	}
}

// Square returns a pointer to square at coords (x, y)
func (b *Board) Square(x, y int) *Square {
	return &b.squares[b.height-y][x-1]
}

// PlacePiece places piece at coords (x, y)
func (b *Board) PlacePiece(x, y int, p Piece) {
	p.SetCoords(x, y)
	square := b.Square(x, y)
	square.piece = p
}

// todo b.MakeMove(...) and test, with legality check
// todo implement other pieces except knight, to implement EP captures or diag captures like pawns, use move history and board

// NewEmptyBoard creates new empty board with i cols and j rows
func NewEmptyBoard(i, j int) *Board {
	b := &Board{}
	b.width, b.height = i, j
	b.createSquares()
	return b
}
