package mtfchess

import (
	"fmt"
)

// Board is a game board
type Board struct {
	squares       Squares
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
	b.squares = make(Squares, b.height)
	for y := range b.squares {
		b.squares[y] = make(Row, b.width)
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
	b.squares[y][x].Empty()
}

// Square returns a pointer to square at coords (x, y)
func (b *Board) Square(x, y int) *Square {
	return &b.squares[b.Y(y)][b.X(x)]
}

// Piece returns a piece at coords (x,y)
func (b *Board) Piece(x, y int) Piece {
	return b.Square(x, y).piece
}

// PlacePiece places piece at coords (x, y)
func (b *Board) PlacePiece(x, y int, p Piece) {
	p.SetCoords(x, y)
	b.Square(x, y).piece = p
}

// Empty removes piece at coords x, y
func (b *Board) Empty(x, y int) {
	b.Square(x, y).Empty()
}

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func (b Board) InCheck(colour Colour) bool {
	return false // todo implement it
}

// Copy returns a pointer to a deep copy of a board
func (b *Board) Copy() *Board {
	newBoard := &Board{}
	newBoard.squares = b.squares.Copy(newBoard)
	newBoard.width, newBoard.height = b.width, b.height
	return newBoard
}

// MakeMove makes move with piece to coords (x,y)
// It returns true if move succesful (legal), otherwise it returns false.
func (b *Board) MakeMove(x, y int, piece Piece) bool {
	c, offsets := piece.Coords(), piece.Offsets(b)
	for _, o := range offsets {
		if c.X+o.X == x && c.Y+o.Y == y {
			newBoard := piece.Project(x, y, b)
			piece.SetCoords(x, y)
			*b = *newBoard
			return true
		}
	}
	return false
}

// todo b.MakeMove(...) and test, with legality check
// todo implement king
// todo implement board.InCheck()
// todo implement legal moves filtering if board.InCheck(), and test for check the next state after move
// todo implement other pieces except knight, to implement EP captures or diag captures like pawns, use move history and board

// NewEmptyBoard creates new empty board with i cols and j rows
func NewEmptyBoard(i, j int) *Board {
	b := &Board{}
	b.width, b.height = i, j
	b.createSquares()
	return b
}
