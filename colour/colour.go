package colour

// Colour is a side colour
type Colour int

const (
	Transparent Colour = iota
	White
	Black
)

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
