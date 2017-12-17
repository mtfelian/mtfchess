package base

import (
	"sort"

	"github.com/mtfelian/iterator"
)

// Coords is an interface to implement like slice of coordinates
type Coords interface {
	sort.Interface
	iterator.Interface
	// Get should return i-th element
	Get(i int) Coord
	// Contains should return true if coordinates bulk contains the given element c
	Contains(c Coord) bool
	// Equals should return true if coordinates bulks are equal
	Equals(to Coords) bool
}
