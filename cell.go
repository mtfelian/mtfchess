package mtfchess

import "fmt"

// Cell is a cell on a board
type Cell struct {
	board *Board
	num   int
	x     int
	y     int
}

// String makes Cell implement Stringer
func (c Cell) String() string {
	return fmt.Sprintf("%7d(%d,%d)", c.num, c.x, c.y)
}
