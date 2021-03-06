package base

import (
	"unicode/utf8"

	"github.com/mtfelian/cli"
	. "github.com/mtfelian/mtfchess/colour"
)

const (
	PawnName       = "pawn"
	KnightName     = "knight"
	BishopName     = "bishop"
	RookName       = "rook"
	QueenName      = "queen"
	ArchbishopName = "archbishop"
	ChancellorName = "chancellor"
	KingName       = "king"
)

// Piece is a base piece
type Piece struct {
	colour         Colour
	coord          ICoord
	name, literals string

	promotion IPiece
	moved     bool
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
func (p *Piece) Name() string { return p.name }

// Capital returns a piece's capital letter
func (p *Piece) Capital() rune { return rune(p.literals[0]) }

// String makes BasePiece to implement fmt.Stringer
func (p *Piece) String() string {
	return map[Colour]string{
		Transparent: cli.Sprintf("{0|%c", p.Capital()),
		White:       cli.Sprintf("{W|%c{0|", p.Capital()),
		Black:       cli.Sprintf("{A|%c{0|", p.Capital()),
	}[p.Colour()]
}

// Colour returns a colour of a piece
func (p *Piece) Colour() Colour { return p.colour }

// SetCoords sets piece's coords to
func (p *Piece) SetCoords(board IBoard, to ICoord) { p.coord = to }

// Coord return piece coords
func (p *Piece) Coord() ICoord { return p.coord }

// Copy returns a copy of a BasePiece
func (p *Piece) Copy() *Piece {
	newPiece := &Piece{
		colour:   p.colour,
		literals: p.literals,
		name:     p.name,
		moved:    p.moved,
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
func (p *Piece) SetPromote(to IPiece) { p.promotion = to }

// Promotion returns a piece in which p piece will be promoted
func (p *Piece) Promotion() IPiece { return p.promotion }

// WasMoved returns true if a piece was moved from it's starting position
func (p *Piece) WasMoved() bool { return p.moved }

// MarkMoved marks piece as moved
func (p *Piece) MarkMoved() { p.moved = true }

// Equals returns true if two pieces are equal
func (p *Piece) Equals(to IPiece) bool {
	return p.Name() == to.Name() && p.Colour() == to.Colour() && p.Coord().Equals(to.Coord())
}
