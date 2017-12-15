package mtfchess

import (
	"fmt"
)

// Row is a row of squares
type Row []Square

// Copy returns a copy of row
func (r Row) Copy(board *Board) Row {
	newRow := make(Row, len(r))
	for i := range r {
		newRow[i] = r[i].Copy(board)
	}
	return newRow
}

// Squares is a matrix of squares
type Squares []Row

// Copy returns a copy of squares
func (s Squares) Copy(board *Board) Squares {
	newSquares := make(Squares, len(s))
	for i := range s {
		newSquares[i] = s[i].Copy(board)
	}
	return newSquares
}

// Board is a game board
type Board struct {
	squares       Squares
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

func (b *Board) Copy() *Board {
	newBoard := &Board{}
	newBoard.squares = b.squares.Copy(newBoard)
	newBoard.width, newBoard.height = b.width, b.height
	return newBoard
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
