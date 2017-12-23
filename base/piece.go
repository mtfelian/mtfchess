package base

import (
	"unicode/utf8"

	"github.com/mtfelian/cli"
	. "github.com/mtfelian/mtfchess/colour"
)

// Piece is a base piece
type Piece struct {
	colour         Colour
	coord          ICoord
	name, literals string

	promotion IPiece
}

// NewPiece creates new base piece with colour
func NewPiece(colour Colour, name, literals string) *Piece {
	if utf8.RuneCountInString(literals) != 3 {
		cli.Println("{R|Invalid literals: %s{0|", literals)
		literals = "?"
	}
	return &Piece{
		colour:   colour,
		name:     name,
		literals: literals,
	}
}

// Name returns the name of a piece
func (p *Piece) Name() string {
	return p.name
}

// String makes BasePiece to implement fmt.Stringer
func (p *Piece) String() string {
	return map[Colour]string{
		Transparent: cli.Sprintf("{0|%s", string(p.literals[0])),
		White:       cli.Sprintf("{W|%s{0|", string(p.literals[0])),
		Black:       cli.Sprintf("{A|%s{0|", string(p.literals[0])),
	}[p.Colour()]
}

// Colour returns a colour of a piece
func (p *Piece) Colour() Colour {
	return p.colour
}

// SetCoords sets piece's coords to
func (p *Piece) SetCoords(board IBoard, to ICoord) {
	p.coord = to
}

// Coord return piece coords
func (p *Piece) Coord() ICoord {
	return p.coord
}

// Copy returns a copy of a BasePiece
func (p *Piece) Copy() *Piece {
	newPiece := &Piece{
		colour:   p.colour,
		literals: p.literals,
		name:     p.name,
	}
	if p.coord != nil {
		newPiece.coord = p.coord.Copy()
	}
	if p.promotion != nil {
		newPiece.promotion = p.promotion.Copy()
	}
	return newPiece
}

// SetPromote sets a piece to promote
func (p *Piece) SetPromote(to IPiece) {
	p.promotion = to
}

// Promotion returns a piece in which p piece will be promoted
func (p *Piece) Promotion() IPiece {
	return p.promotion
}

// Equals returns true if two pieces are equal
func (p *Piece) Equals(to IPiece) bool {
	return p.Name() == to.Name() && p.Colour() == to.Colour() && p.Coord().Equals(to.Coord())
}
