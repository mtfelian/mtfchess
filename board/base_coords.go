package board

// BaseCoords is a base coords
type BaseCoords struct {
	slice []Coord
	i     int
}

// NewBaseCoords returns new base coordinates
func NewBaseCoords(s []Coord) *BaseCoords {
	return &BaseCoords{slice: s, i: 0}
}

func (s *BaseCoords) Get(i int) Coord { return s.slice[i] }

// Next returns next coordinates element
func (s *BaseCoords) Next() interface{} {
	s.i++
	return s.slice[s.i-1]
}

// HasNext returns true if an underlying slice has next element
func (s *BaseCoords) HasNext() bool { return s.i < len(s.slice) }

// I returns a current iteration index
func (s *BaseCoords) I() int        { return s.i - 1 }
func (s *BaseCoords) Len() int      { return len(s.slice) }
func (s *BaseCoords) Swap(i, j int) { s.slice[i], s.slice[j] = s.slice[j], s.slice[i] }

// Add adds an element to an underlying slice
func (s *BaseCoords) Add(c interface{}) { s.slice = append(s.slice, c.(Coord)) }

// Contains returns true if c contains in s
func (s *BaseCoords) Contains(c Coord) bool {
	for i := range s.slice {
		if s.Get(i).Equals(c) {
			return true
		}
	}
	return false
}

// Equals returns true if c equals to
func (c *BaseCoords) Equals(to Coords) bool {
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
