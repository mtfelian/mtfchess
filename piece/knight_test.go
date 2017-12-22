package piece_test

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Knight test", func() {
	var b base.IBoard
	resetBoard := func() { b = rect.NewTestEmptyBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wn1, wn2, bn := piece.NewKnight(White), piece.NewKnight(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 1}, wn1)
		b.PlacePiece(rect.Coord{3, 3}, wn2)
		b.PlacePiece(rect.Coord{4, 2}, bn)

		d := wn1.Destinations(b)
		Expect(d.Len()).To(Equal(2))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{4, 2}, rect.Coord{1, 3}}))).To(BeTrue())
	})

	It("attacks right cells, can release check by capture", func() {
		wn, bq := piece.NewKnight(White), piece.NewQueen(Black)
		bn, wr, wk := piece.NewKnight(Black), piece.NewRook(White), piece.NewKing(White)
		b.PlacePiece(rect.Coord{2, 5}, bq)
		b.PlacePiece(rect.Coord{4, 4}, wn)
		b.PlacePiece(rect.Coord{2, 3}, wk)
		b.PlacePiece(rect.Coord{5, 1}, bn)
		b.PlacePiece(rect.Coord{5, 4}, wr)

		attacking := wn.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{3, 2}, rect.Coord{5, 2}, rect.Coord{2, 3},
			rect.Coord{2, 5}, rect.Coord{3, 6}, rect.Coord{5, 6},
		}))).To(BeTrue())

		Expect(b.MakeMove(rect.Coord{2, 3}, wn)).To(BeFalse(), "captured own piece, and king in check")
		Expect(b.MakeMove(rect.Coord{3, 2}, wn)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(rect.Coord{6, 3}, wn)).To(BeFalse(), "jumped out of board, and check")

		// successful capture releasing check
		Expect(b.MakeMove(rect.Coord{2, 5}, wn)).To(BeTrue(), "can't capture releasing check")
	})

	It("attacks right cells, can release check by pinning self", func() {
		wn, bq := piece.NewKnight(White), piece.NewQueen(Black)
		bn, wr, wk := piece.NewKnight(Black), piece.NewRook(White), piece.NewKing(White)
		b.PlacePiece(rect.Coord{2, 5}, bq)
		b.PlacePiece(rect.Coord{4, 3}, wn)
		b.PlacePiece(rect.Coord{2, 3}, wk)
		b.PlacePiece(rect.Coord{5, 1}, bn)
		b.PlacePiece(rect.Coord{5, 4}, wr)

		attacking := wn.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{3, 1}, rect.Coord{5, 1}, rect.Coord{2, 2},
			rect.Coord{2, 4}, rect.Coord{3, 5}, rect.Coord{5, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(rect.Coord{3, 1}, wn)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(rect.Coord{5, 1}, wn)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(rect.Coord{6, 2}, wn)).To(BeFalse(), "jumped out of board, and check")

		// successful capture releasing check
		Expect(b.MakeMove(rect.Coord{2, 4}, wn)).To(BeTrue(), "can't pin self releasing check")
	})

	It("makes legal moves", func() {
		var wn, bq base.IPiece
		testReset := func() {
			resetBoard()
			wn, bq = piece.NewKnight(White), piece.NewQueen(Black)
			b.PlacePiece(rect.Coord{2, 1}, wn)
			b.PlacePiece(rect.Coord{4, 2}, bq)
		}
		testReset()
		destinations := wn.Destinations(b)
		sort.Sort(destinations)
		Expect(destinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{4, 2}, rect.Coord{1, 3}, rect.Coord{3, 3},
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
			wn, bn = piece.NewKnight(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 1}, wn)
			b.PlacePiece(rect.Coord{4, 2}, bn)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{rect.Coord{5, 2}, rect.Coord{1, 4}, wn.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wn.Coord()
			Expect(b.MakeMove(d, wn)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wn))

			// check that destination cell was not changed
			p := b.Piece(d)
			if p == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(b.Piece(d)))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bn.Coord())).To(Equal(bn))

			testReset()
		}
	})
})
