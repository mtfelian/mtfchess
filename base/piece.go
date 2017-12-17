package base

import (
	"unicode/utf8"

	"github.com/mtfelian/cli"
)

// BasePiece
type BasePiece struct {
	colour         Colour
	coord          Coord
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

// Name returns the name of a piece
func (p *BasePiece) Name() string {
	return p.name
}

// String makes BasePiece to implement fmt.Stringer
func (p *BasePiece) String() string {
	return map[Colour]string{
		Transparent: cli.Sprintf("{0|%s", string(p.literals[0])),
		White:       cli.Sprintf("{W|%s{0|", string(p.literals[0])),
		Black:       cli.Sprintf("{A|%s{0|", string(p.literals[0])),
	}[p.Colour()]
}

// Colour returns a colour of a piece
func (p *BasePiece) Colour() Colour {
	return p.colour
}

// SetCoords sets piece's coords to
func (p *BasePiece) SetCoords(to Coord) {
	p.coord = to
}

// Coord return piece coords
func (p *BasePiece) Coord() Coord {
	return p.coord
}

// Copy returns a copy of a BasePiece
func (p *BasePiece) Copy() BasePiece {
	return BasePiece{
		colour:   p.colour,
		coord:    p.coord.Copy(),
		literals: p.literals,
		name:     p.name,
	}
}
