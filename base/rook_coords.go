package base

import . "github.com/mtfelian/mtfchess/colour"

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
