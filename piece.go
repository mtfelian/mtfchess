package mtfchess

import (
	"fmt"
	"unicode/utf8"

	"github.com/mtfelian/cli"
)

// Pair is a coordinate pair
type Pair struct {
	X, Y int
}

// Offsets is a slice of pair offsets
type Offsets []Pair

type Board interface {
	Width() int
	Height() int
	SetWidth(width int)
	SetHeight(height int)
	Square(x, y int) *Square
	InCheck(colour Colour) bool
	Squares() Squares
	Copy() Board
	Empty(x, y int)
	PlacePiece(x, y int, p Piece)
	Set(b1 Board)
	MakeMove(x, y int, piece Piece) bool
	Piece(x, y int) Piece
}

// Piece
type Piece interface {
	fmt.Stringer
	// Name returns piece name
	Name() string
	// Offsets returns a slice of offsets to possible moves
	Offsets(b Board) Offsets
	// Colour returns piece colour
	Colour() Colour
	// SetCoords sets the figure coords to (x,y)
	SetCoords(x, y int)
	// Cords returns a pair of coords
	Coords() Pair
	// Copy returns a deep copy of a piece
	Copy() Piece
	// Project a piece to coords (x,y), returns a pointer to a new copy of a board, don't check legality
	// this don't change coords of a piece
	Project(x, y int, b Board) Board
}

// BasePiece
type BasePiece struct {
	colour         Colour
	x, y           int
	name, literals string
}

// NewBasePiece creates new base piece with colour
func NewBasePiece(colour Colour, name, literals string) BasePiece {
	if utf8.RuneCountInString(literals) != 3 {
		cli.Println("{R|Invalid literals: %s{0|", literals)
		literals = "?"
	}
	return BasePiece{
		colour:   colour,
		name:     name,
		literals: literals,
	}
}

func (p *BasePiece) Name() string {
	return p.name
}

func (p *BasePiece) String() string {
	return map[Colour]string{
		Transparent: cli.Sprintf("{0|%s", string(p.literals[0])),
		White:       cli.Sprintf("{W|%s{0|", string(p.literals[0])),
		Black:       cli.Sprintf("{A|%s{0|", string(p.literals[0])),
	}[p.Colour()]
}

func (p *BasePiece) Colour() Colour {
	return p.colour
}

func (p *BasePiece) SetCoords(x, y int) {
	p.x, p.y = x, y
}

func (p *BasePiece) Coords() Pair {
	return Pair{X: p.x, Y: p.y}
}

func (p *BasePiece) Copy() BasePiece {
	return BasePiece{
		colour:   p.colour,
		x:        p.x,
		y:        p.y,
		literals: p.literals,
		name:     p.name,
	}
}
