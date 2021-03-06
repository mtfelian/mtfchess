package base

import (
	"fmt"
	"sort"

	"github.com/mtfelian/iterator"
)

// ICoords is an interface to implement like slice of coordinates
type ICoords interface {
	fmt.Stringer
	sort.Interface
	iterator.Interface
	// Get should return i-th element
	Get(i int) ICoord
	// Contains should return true if coordinates bulk contains the given element c
	Contains(c ICoord) bool
	// Equals should return true if coordinates bulks are equal
	Equals(to ICoords) bool
}
