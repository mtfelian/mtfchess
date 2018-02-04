package rect

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chancellor test", func() {
	var b base.IBoard
	resetBoard := func() { b = NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wc, wn, bc := NewChancellor(White), NewKnight(White), NewChancellor(Black)
		b.PlacePiece(Coord{4, 2}, wc)
		b.PlacePiece(Coord{4, 4}, wn)
		b.PlacePiece(Coord{4, 1}, bc)

		d := wc.Destinations(b)
		sort.Sort(d)
		Expect(d.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{4, 1}, Coord{1, 2}, Coord{2, 2},
			Coord{3, 2}, Coord{5, 2}, Coord{2, 3}, Coord{4, 3},
			Coord{3, 4}, Coord{5, 4},
		}))).To(BeTrue())
	})

	It("attacks right cells", func() {
		wc, bq := NewChancellor(White), NewQueen(Black)
		bn, wn, wk := NewKnight(Black), NewKnight(White), NewKing(White)
		b.PlacePiece(Coord{1, 2}, bq)
		b.PlacePiece(Coord{2, 2}, wc)
		b.PlacePiece(Coord{3, 2}, wk)
		b.PlacePiece(Coord{2, 1}, bn)
		b.PlacePiece(Coord{2, 5}, wn)

		attacking := wc.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{4, 1}, Coord{1, 2}, Coord{3, 2},
			Coord{2, 3}, Coord{4, 3}, Coord{1, 4}, Coord{2, 4},
			Coord{3, 4}, Coord{2, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{2, 5}, wc)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(Coord{3, 2}, wc)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(Coord{2, 6}, wc)).To(BeFalse(), "jumped over own piece")
		Expect(b.MakeMove(Coord{1, 2}, wc)).To(BeTrue(), "can't capture")
	})

	It("makes legal moves", func() {
		var wc, bc base.IPiece
		testReset := func() {
			resetBoard()
			wc, bc = NewChancellor(White), NewChancellor(Black)
			b.PlacePiece(Coord{2, 1}, wc)
			b.PlacePiece(Coord{4, 1}, bc)
		}
		testReset()

		destinations := wc.Destinations(b)
		sort.Sort(destinations)
		Expect(destinations.Equals(NewCoords([]base.ICoord{
			Coord{1, 1}, Coord{3, 1}, Coord{4, 1},
			Coord{2, 2}, Coord{4, 2}, Coord{1, 3},
			Coord{2, 3}, Coord{3, 3}, Coord{2, 4},
			Coord{2, 5}, Coord{2, 6},
		}))).To(BeTrue())

		bcCoord, wcCoord := bc.Coord().Copy(), wc.Coord().Copy()
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wc)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wcCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wc))
			if !bcCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(bcCoord)).To(Equal(bc))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wc.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(bc.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wc, bc base.IPiece
		testReset := func() {
			resetBoard()
			wc, bc = NewChancellor(White), NewChancellor(Black)
			b.PlacePiece(Coord{2, 1}, wc)
			b.PlacePiece(Coord{4, 1}, bc)
		}
		testReset()

		destinations := NewCoords([]base.ICoord{Coord{3, 2}, Coord{5, 1}, wc.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(Coord), wc.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wc)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wc))

			// check that destination cell was not changed
			if dCellCopy.Piece() == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(dCellCopy.Piece()))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bc.Coord())).To(Equal(bc))

			testReset()
		}
	})
})
