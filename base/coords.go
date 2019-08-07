package base

import (
	"strings"
)

// Coords is a base coords
type Coords struct {
	slice []ICoord
	i     int
}

// NewCoords returns new base coordinates
func NewCoords(s []ICoord) *Coords {
	return &Coords{slice: s, i: 0}
}

func (s *Coords) Get(i int) ICoord { return s.slice[i] }

// Next returns next coordinates element
func (s *Coords) Next() interface{} {
	s.i++
	return s.slice[s.i-1]
}

// HasNext returns true if an underlying slice has next element
func (s *Coords) HasNext() bool { return s.i < len(s.slice) }

// SetI sets current iteration index
func (s *Coords) SetI(i int) { s.i = i }

// I returns a current iteration index
func (s *Coords) I() int        { return s.i - 1 }
func (s *Coords) Len() int      { return len(s.slice) }
func (s *Coords) Swap(i, j int) { s.slice[i], s.slice[j] = s.slice[j], s.slice[i] }

// Add adds an element to an underlying slice
func (s *Coords) Add(c interface{}) { s.slice = append(s.slice, c.(ICoord)) }

// Contains returns true if c contains in s
func (s *Coords) Contains(c ICoord) bool {
	for i := range s.slice {
		if s.Get(i).Equals(c) {
			return true
		}
	}
	return false
}

// Equals returns true if c equals to
func (c *Coords) Equals(to ICoords) bool {
	if c.Len() != to.Len() {
		return false
	}
	for i := range c.slice {
		if !c.Get(i).Equals(to.Get(i)) {
			return false
		}
	}
	return true
}

// String makes Coords to implement fmt.Stringer
func (c *Coords) String() string {
	s := "("
	for i := range c.slice {
		s += c.slice[i].String() + ","
	}
	return strings.TrimRight(s, ",") + ")"
}
