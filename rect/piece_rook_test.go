package rect

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rook test", func() {
	var b base.IBoard
	resetBoard := func() { b = NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wr, wn, br := NewRook(White), NewKnight(White), NewRook(Black)
		b.PlacePiece(Coord{4, 2}, wr)
		b.PlacePiece(Coord{4, 4}, wn)
		b.PlacePiece(Coord{4, 1}, br)

		d := wr.Destinations(b)
		Expect(d.Len()).To(Equal(6))
		sort.Sort(d)
		Expect(d.Equals(NewCoords([]base.ICoord{
			Coord{4, 1}, Coord{1, 2}, Coord{2, 2},
			Coord{3, 2}, Coord{5, 2}, Coord{4, 3},
		}))).To(BeTrue())
	})

	It("attacks right cells", func() {
		wr, bq := NewRook(White), NewQueen(Black)
		bn, wn, wk := NewKnight(Black), NewKnight(White), NewKing(White)
		b.PlacePiece(Coord{1, 2}, bq)
		b.PlacePiece(Coord{2, 2}, wr)
		b.PlacePiece(Coord{3, 2}, wk)
		b.PlacePiece(Coord{2, 1}, bn)
		b.PlacePiece(Coord{2, 5}, wn)

		attacking := wr.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{1, 2}, Coord{3, 2},
			Coord{2, 3}, Coord{2, 4}, Coord{2, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{2, 5}, wr)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(Coord{3, 2}, wr)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(Coord{2, 6}, wr)).To(BeFalse(), "jumped over own piece")
		Expect(b.MakeMove(Coord{1, 2}, wr)).To(BeTrue(), "can't capture")
	})

	It("makes legal moves", func() {
		var wr, br base.IPiece
		testReset := func() {
			resetBoard()
			wr, br = NewRook(White), NewRook(Black)
			b.PlacePiece(Coord{2, 1}, wr)
			b.PlacePiece(Coord{4, 1}, br)
		}
		testReset()

		destinations := wr.Destinations(b)
		sort.Sort(destinations)
		Expect(destinations.Equals(NewCoords([]base.ICoord{
			Coord{1, 1}, Coord{3, 1}, Coord{4, 1},
			Coord{2, 2}, Coord{2, 3}, Coord{2, 4},
			Coord{2, 5}, Coord{2, 6},
		}))).To(BeTrue())

		brCoord, wrCoord := br.Coord().Copy(), wr.Coord().Copy()
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wr)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wrCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wr))
			if !brCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(brCoord)).To(Equal(br))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wr.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(br.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wr, br base.IPiece
		testReset := func() {
			resetBoard()
			wr, br = NewRook(White), NewRook(Black)
			b.PlacePiece(Coord{2, 1}, wr)
			b.PlacePiece(Coord{4, 1}, br)
		}
		testReset()

		destinations := NewCoords([]base.ICoord{Coord{3, 2}, Coord{5, 1}, wr.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(Coord), wr.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wr)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wr))

			// check that destination cell was not changed
			if dCellCopy.Piece() == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(dCellCopy.Piece()))
			}

			// check another cell to contain another piece
			Expect(b.Piece(br.Coord())).To(Equal(br))

			testReset()
		}
	})
})
