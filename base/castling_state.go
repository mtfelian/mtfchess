package base

import . "github.com/mtfelian/mtfchess/colour"

// CastlingEnabled
type CastlingState map[Colour][2]bool

// Copy returns a copy of c with pieces taken from board
func (c CastlingState) Copy() CastlingState {
	res := CastlingState{}
	for colour, states := range c {
		for i := range states {
			arr := res[colour]
			arr[i] = states[i]
			res[colour] = arr
		}
	}
	return res
}
