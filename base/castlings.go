package base

// Castlings is a slice of castlings
type Castlings []Castling

// Contains returns true if c contains castling
func (c Castlings) Contains(castling Castling) bool {
	for i := range c {
		if c[i].Equal(castling) {
			return true
		}
	}
	return false
}
