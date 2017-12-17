package board

import (
	"fmt"
	"sort"

	. "github.com/mtfelian/mtfchess/iterator"
)

// Coord
type Coord interface {
	fmt.Stringer
	// Add returns a sum of coords
	Add(c Coord) Coord
	// Out returns true if coords is out of board
	Out(b Board) bool
	// Equals returns true if coords are equal
	Equals(to Coord) bool
	// Copy returns a copy of coord
	Copy() Coord
}

// BaseCoords is a base coords
type BaseCoords struct {
	// Slice is an underlying slice
	Slice []Coord
	// Index for iterations
	Index int
}

// Next returns next coordinates element
func (i *BaseCoords) Next() interface{} {
	i.Index++
	return i.Slice[i.Index-1]
}

// HasNext returns true if an underlying slice has next element
func (i *BaseCoords) HasNext() bool { return i.Index < len(i.Slice) }

// I returns a current iteration index
func (i *BaseCoords) I() int { return i.Index }

// Add adds an element to an underlying slice
func (i *BaseCoords) Add(c interface{}) { i.Slice = append(i.Slice, c.(Coord)) }

// Coords is an interface to implement like slice of coordinates
type Coords interface {
	sort.Interface // should implement it
	Iterator
	// Get should return i-th element
	Get(i int) Coord
	// Contains should return true if coordinates bulk contains the given element c
	Contains(c Coord) bool
	// Equals should return true if coordinates bulks are equal
	Equals(to Coords) bool
}
