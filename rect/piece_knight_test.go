package rect

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Knight test", func() {
	var b base.IBoard
	resetBoard := func() { b = NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wn1, wn2, bn := NewKnight(White), NewKnight(White), NewKnight(Black)
		b.PlacePiece(Coord{2, 1}, wn1)
		b.PlacePiece(Coord{3, 3}, wn2)
		b.PlacePiece(Coord{4, 2}, bn)

		d := wn1.Destinations(b)
		sort.Sort(d)
		Expect(d.Equals(NewCoords([]base.ICoord{Coord{4, 2}, Coord{1, 3}}))).To(BeTrue())
	})

	It("attacks right cells, can release check by capture", func() {
		wn, bq := NewKnight(White), NewQueen(Black)
		bn, wr, wk := NewKnight(Black), NewRook(White), NewKing(White)
		b.PlacePiece(Coord{2, 5}, bq)
		b.PlacePiece(Coord{4, 4}, wn)
		b.PlacePiece(Coord{2, 3}, wk)
		b.PlacePiece(Coord{5, 1}, bn)
		b.PlacePiece(Coord{5, 4}, wr)

		attacking := wn.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{3, 2}, Coord{5, 2}, Coord{2, 3},
			Coord{2, 5}, Coord{3, 6}, Coord{5, 6},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{2, 3}, wn)).To(BeFalse(), "captured own piece, and king in check")
		Expect(b.MakeMove(Coord{3, 2}, wn)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(Coord{6, 3}, wn)).To(BeFalse(), "jumped out of board, and check")

		// successful capture releasing check
		Expect(b.MakeMove(Coord{2, 5}, wn)).To(BeTrue(), "can't capture releasing check")
	})

	It("attacks right cells, can release check by pinning self", func() {
		wn, bq := NewKnight(White), NewQueen(Black)
		bn, wr, wk := NewKnight(Black), NewRook(White), NewKing(White)
		b.PlacePiece(Coord{2, 5}, bq)
		b.PlacePiece(Coord{4, 3}, wn)
		b.PlacePiece(Coord{2, 3}, wk)
		b.PlacePiece(Coord{5, 1}, bn)
		b.PlacePiece(Coord{5, 4}, wr)

		attacking := wn.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{3, 1}, Coord{5, 1}, Coord{2, 2},
			Coord{2, 4}, Coord{3, 5}, Coord{5, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{3, 1}, wn)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(Coord{5, 1}, wn)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(Coord{6, 2}, wn)).To(BeFalse(), "jumped out of board, and check")

		// successful capture releasing check
		Expect(b.MakeMove(Coord{2, 4}, wn)).To(BeTrue(), "can't pin self releasing check")
	})

	It("makes legal moves", func() {
		var wn, bq base.IPiece
		testReset := func() {
			resetBoard()
			wn, bq = NewKnight(White), NewQueen(Black)
			b.PlacePiece(Coord{2, 1}, wn)
			b.PlacePiece(Coord{4, 2}, bq)
		}
		testReset()
		destinations := wn.Destinations(b)
		sort.Sort(destinations)
		Expect(destinations.Equals(NewCoords([]base.ICoord{
			Coord{4, 2}, Coord{1, 3}, Coord{3, 3},
		}))).To(BeTrue())

		bqCoord, wnCoord := bq.Coord().Copy(), wn.Coord().Copy()
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wn)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wnCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wn))
			if !bqCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(bqCoord)).To(Equal(bq))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wn.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(bq.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wn, bn base.IPiece
		testReset := func() {
			wn, bn = NewKnight(White), NewKnight(Black)
			b.PlacePiece(Coord{2, 1}, wn)
			b.PlacePiece(Coord{4, 2}, bn)
		}
		testReset()

		destinations := NewCoords([]base.ICoord{Coord{5, 2}, Coord{1, 4}, wn.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(Coord), wn.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wn)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wn))

			// check that destination cell was not changed
			if dCellCopy.Piece() == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(dCellCopy.Piece()))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bn.Coord())).To(Equal(bn))

			testReset()
		}
	})
})
