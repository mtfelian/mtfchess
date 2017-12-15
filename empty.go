package mtfchess

type Empty struct {
	x, y int
}

// NewEmpty creates new empty figure
func NewEmpty(x, y int) Piece {
	empty := &Empty{}
	empty.SetCoords(x, y)
	return empty
}

func (p *Empty) Name() string {
	return ""
}

func (p *Empty) CanJump() bool {
	return true
}

func (p *Empty) Colour() Colour {
	return Transparent
}

func (p *Empty) String() string {
	return " "
}

func (p *Empty) SetCoords(x, y int) {
	p.x, p.y = x, y
}

func (p *Empty) Offsets(b *Board) Offsets {
	return Offsets{}
}

func (p *Empty) Project(x, y int, b *Board) *Board {
	newBoard := b.Copy()
	newBoard.Empty(x, y)
	return newBoard
}

func (p *Empty) Coords() Pair {
	return Pair{X: p.x, Y: p.y}
}

func (p *Empty) Copy() Piece {
	return &Empty{
		x: p.x,
		y: p.y,
	}
}
