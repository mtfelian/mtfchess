package base

import (
	. "github.com/mtfelian/mtfchess/colour"
)

// RookCoords maps colour to array of rook coords
type RookCoords map[Colour][2]ICoord

// NewRookCoords returns new rook coords
func NewRookCoords() RookCoords { return RookCoords{White: [2]ICoord{}, Black: [2]ICoord{}} }

// Copy returns a copy of c
func (c RookCoords) Copy() RookCoords {
	res := RookCoords{}
	for colour, state := range c {
		coords := [2]ICoord{}
		for i := range state {
			if state[i] != nil {
				coords[i] = state[i].Copy()
			}
		}
		res[colour] = coords
	}
	return res
}

// Equals returns true if c equals to, and returns false otherwise
func (c RookCoords) Equals(to RookCoords) bool {
	for colour := range c {
		if _, exists := to[colour]; !exists || len(c[colour]) != len(to[colour]) {
			return false
		}
		for i := range c[colour] {
			if (c[colour][i] == nil) != (to[colour][i] == nil) {
				return false
			}
			if c[colour][i] != nil && to[colour][i] == nil && !c[colour][i].Equals(to[colour][i]) {
				return false
			}
		}
	}
	for colour := range to {
		if _, exists := c[colour]; !exists {
			return false
		}
	}
	return true
}
