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
	Next() Coord
	HasNext() bool
	I() int
	Add(c Coord)
}

// BaseIterator is a base type for iterations
type CoordsIterator struct {
	Slice []Coord
	Index int
}

func (i *CoordsIterator) Next() Coord {
	i.Index++
	return i.Slice[i.Index-1]
}

func (i *CoordsIterator) HasNext() bool { return i.Index < len(i.Slice) }

func (i *CoordsIterator) I() int { return i.Index }

func (i *CoordsIterator) Add(c Coord) { i.Slice = append(i.Slice, c) }

// Coords potentially is a slice of coords
type Coords interface {
	sort.Interface // should implement it
	Iterator
	Get(i int) Coord
	Contains(c Coord) bool
	Equals(to Coords) bool
}
