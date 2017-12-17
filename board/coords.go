package board

import (
	"fmt"
	"sort"
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

// Iterator is an interface for iterator
type Iterator interface {
	// Next should return the next element
	Next() Coord
	// HasNext should return true if we have next element
	HasNext() bool
	// I should return an iteration index
	I() int
	// Add should add c to an underlying storage
	Add(c Coord)
}

// CoordsIterator is a coordinates iterator
type CoordsIterator struct {
	// Slice is an underlying slice
	Slice []Coord
	// Index for iterations
	Index int
}

// Next returns next coordinates element
func (i *CoordsIterator) Next() Coord {
	i.Index++
	return i.Slice[i.Index-1]
}

// HasNext returns true if an underlying slice has next element
func (i *CoordsIterator) HasNext() bool { return i.Index < len(i.Slice) }

// I returns a current iteration index
func (i *CoordsIterator) I() int { return i.Index }

// Add adds an element to an underlying slice
func (i *CoordsIterator) Add(c Coord) { i.Slice = append(i.Slice, c) }

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
