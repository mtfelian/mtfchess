package mtfchess

import (
	"fmt"

	. "github.com/mtfelian/mtfchess/board"
)

// RectCoord is a rectangular coordinates
type RectCoord struct {
	X, Y int
}

// String makes RectCoord to implement fmt.Stringer
func (c RectCoord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

// Add adds o to c and returns the sum as a result
func (c RectCoord) Add(o Coord) Coord {
	return RectCoord{X: c.X + o.(RectCoord).X, Y: c.Y + o.(RectCoord).Y}
}

// Out returns true if c is a coords out of board
func (c RectCoord) Out(b Board) bool {
	return c.X < 1 || c.Y < 1 || c.X > b.Dim().(RectCoord).X || c.Y > b.Dim().(RectCoord).Y
}

// Equals returns true if c equals c1
func (c RectCoord) Equals(to Coord) bool {
	return c.X == to.(RectCoord).X && c.Y == to.(RectCoord).Y
}

// Copy returns a copy of c
func (c RectCoord) Copy() Coord {
	return RectCoord{X: c.X, Y: c.Y}
}

// NewRectCoords returns new rectangular coordinates
func NewRectCoords(c []RectCoord) RectCoords {
	s := make([]Coord, len(c))
	for i := range c {
		s[i] = c[i]
	}
	return RectCoords{BaseCoords: NewBaseCoords(s)}
}

// RectCoords is a slice of rectangular coordinates
type RectCoords struct {
	*BaseCoords
}

func (s RectCoords) Less(i, j int) bool {
	siX, siY := s.Get(i).(RectCoord).X, s.Get(i).(RectCoord).Y
	sjX, sjY := s.Get(j).(RectCoord).X, s.Get(j).(RectCoord).Y
	return siY < sjY || (siY == sjY && siX < sjX)
}
