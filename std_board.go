package mtfchess

import (
	"fmt"
)

// StdBoard is a game standard board
type StdBoard struct {
	squares       Squares
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
	for _, row := range b.squares {
		for _, square := range row {
			s += fmt.Sprintf("%s", square)
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

// createSquares returns a slice of Square for the board
func (b *StdBoard) createSquares() {
	b.squares = make(Squares, b.height)
	for y := range b.squares {
		b.squares[y] = make(Row, b.width)
		for x := range b.squares[y] {
			b.createSquare(x, y)
		}
	}
}

// createSquare creates new square at rectangular board b with coordinates x, y and row length i
func (b *StdBoard) createSquare(x, y int) {
	b.squares[y][x] = Square{
		board: b,
		num:   b.width*(b.height-y) - b.width + (x + 1),
		x:     x + 1,
		y:     b.height - y,
	}
	b.squares[y][x].Empty()
}

// Square returns a pointer to square at coords (x, y)
func (b *StdBoard) Square(x, y int) *Square {
	return &b.squares[b.Y(y)][b.X(x)]
}

// Squares returns a squares slice
func (b *StdBoard) Squares() Squares {
	return b.squares
}

// SetSquares sets squares to s
func (b *StdBoard) SetSquares(s Squares) {
	b.squares = s
}

// Piece returns a piece at coords (x,y)
func (b *StdBoard) Piece(x, y int) Piece {
	return b.Square(x, y).piece
}

// PlacePiece places piece at coords (x, y)
func (b *StdBoard) PlacePiece(x, y int, p Piece) {
	p.SetCoords(x, y)
	b.Square(x, y).piece = p
}

// Empty removes piece at coords x, y
func (b *StdBoard) Empty(x, y int) {
	b.Square(x, y).Empty()
}

// InCheck returns true if there is a check on the board for colour, otherwise it returns false
func (b StdBoard) InCheck(colour Colour) bool {
	return false // todo implement it
}

// Copy returns a pointer to a deep copy of a board
func (b *StdBoard) Copy() Board {
	newBoard := &StdBoard{}
	newBoard.SetSquares(b.Squares().Copy(newBoard))
	newBoard.SetWidth(b.Width())
	newBoard.SetHeight(b.Height())
	return newBoard
}

// Set changes b to b1
func (b *StdBoard) Set(b1 Board) {
	b.SetWidth(b1.Width())
	b.SetHeight(b1.Height())
	b.SetSquares(b1.Squares())
}

// MakeMove makes move with piece to coords (x,y)
// It returns true if move succesful (legal), otherwise it returns false.
func (b *StdBoard) MakeMove(x, y int, piece Piece) bool {
	c, offsets := piece.Coords(), piece.Offsets(b)
	for _, o := range offsets {
		if c.X+o.X == x && c.Y+o.Y == y {
			newBoard := piece.Project(x, y, b)
			piece.SetCoords(x, y)
			b.Set(newBoard)
			return true
		}
	}
	return false
}

// todo implement king
// todo implement board.InCheck()
// todo implement other pieces except knight, to implement EP captures or diag captures like pawns, use move history and board

// NewEmptyStdBoard creates new empty standard board with i cols and j rows
func NewEmptyStdBoard(i, j int) *StdBoard {
	b := &StdBoard{}
	b.width, b.height = i, j
	b.createSquares()
	return b
}
