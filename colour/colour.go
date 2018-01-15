package colour

import "github.com/mtfelian/cli"

// Colour is a side colour
type Colour int

const (
	Transparent Colour = iota
	White
	Black
)

// Name returns a side name
func (c Colour) Name() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	}
	return ""
}

// String make Colour to implement fmt.Stringer
func (c Colour) String() string {
	switch c {
	case White:
		return cli.Sprintf("{W|W{0|")
	case Black:
		return cli.Sprintf("{A|B{0|")
	}
	return cli.Sprintf("{R|?{0|")
}

// Invert returns an inverted colour
func (c Colour) Invert() Colour {
	switch c {
	case White:
		return Black
	case Black:
		return White
	}
	return Transparent
}
