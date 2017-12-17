package rect

import (
	"fmt"

	"github.com/mtfelian/mtfchess/base"
)

// Coord is a rectangular coordinates
type Coord struct {
	X, Y int
}

// String makes RectCoord to implement fmt.Stringer
func (c Coord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

// Add adds o to c and returns the sum as a result
func (c Coord) Add(o base.ICoord) base.ICoord {
	return Coord{X: c.X + o.(Coord).X, Y: c.Y + o.(Coord).Y}
}

// Out returns true if c is a coords out of board
func (c Coord) Out(b base.IBoard) bool {
	return c.X < 1 || c.Y < 1 || c.X > b.Dim().(Coord).X || c.Y > b.Dim().(Coord).Y
}

// Equals returns true if c equals c1
func (c Coord) Equals(to base.ICoord) bool {
	return c.X == to.(Coord).X && c.Y == to.(Coord).Y
}

// Copy returns a copy of c
func (c Coord) Copy() base.ICoord {
	return Coord{X: c.X, Y: c.Y}
}

// NewCoords returns new rectangular coordinates
func NewCoords(c []base.ICoord) Coords {
	return Coords{Coords: base.NewCoords(c)}
}

// Coords is a slice of rectangular coordinates
type Coords struct {
	*base.Coords
}

func (s Coords) Less(i, j int) bool {
	siX, siY := s.Get(i).(Coord).X, s.Get(i).(Coord).Y
	sjX, sjY := s.Get(j).(Coord).X, s.Get(j).(Coord).Y
	return siY < sjY || (siY == sjY && siX < sjX)
}
