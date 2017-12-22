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

var _ = Describe("Queen test", func() {
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("generates moves", func() {
		wq, wn, bq := piece.NewQueen(White), piece.NewKnight(White), piece.NewQueen(Black)
		b.PlacePiece(rect.Coord{2, 3}, wq)
		b.PlacePiece(rect.Coord{1, 2}, wn)
		b.PlacePiece(rect.Coord{3, 3}, bq)

		wqDestinations := wq.Destinations(b)
		Expect(wqDestinations.Len()).To(Equal(13))
		sort.Sort(wqDestinations)
		Expect(wqDestinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{4, 1}, rect.Coord{2, 2}, rect.Coord{3, 2},
			rect.Coord{1, 3}, rect.Coord{3, 3}, rect.Coord{1, 4}, rect.Coord{2, 4},
			rect.Coord{3, 4}, rect.Coord{2, 5}, rect.Coord{4, 5}, rect.Coord{2, 6},
			rect.Coord{5, 6},
		}))).To(BeTrue())

		bqDestinations := bq.Destinations(b)
		Expect(bqDestinations.Len()).To(Equal(16))
		sort.Sort(bqDestinations)
		Expect(bqDestinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{1, 1}, rect.Coord{3, 1}, rect.Coord{5, 1}, rect.Coord{2, 2},
			rect.Coord{3, 2}, rect.Coord{4, 2}, rect.Coord{2, 3}, rect.Coord{4, 3},
			rect.Coord{5, 3}, rect.Coord{2, 4}, rect.Coord{3, 4}, rect.Coord{4, 4},
			rect.Coord{1, 5}, rect.Coord{3, 5}, rect.Coord{5, 5}, rect.Coord{3, 6},
		}))).To(BeTrue())
	})

	It("attacks right cells", func() {
		wq, bq := piece.NewQueen(White), piece.NewQueen(Black)
		bn, wr, wk := piece.NewKnight(Black), piece.NewRook(White), piece.NewKing(White)
		b.PlacePiece(rect.Coord{2, 5}, bq)
		b.PlacePiece(rect.Coord{4, 3}, wq)
		b.PlacePiece(rect.Coord{2, 3}, wk)
		b.PlacePiece(rect.Coord{5, 1}, bn)
		b.PlacePiece(rect.Coord{5, 4}, wr)

		attacking := wq.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{4, 1}, rect.Coord{3, 2}, rect.Coord{4, 2},
			rect.Coord{5, 2}, rect.Coord{2, 3}, rect.Coord{3, 3}, rect.Coord{5, 3},
			rect.Coord{3, 4}, rect.Coord{4, 4}, rect.Coord{5, 4}, rect.Coord{2, 5},
			rect.Coord{4, 5}, rect.Coord{4, 6},
		}))).To(BeTrue())

		Expect(b.MakeMove(rect.Coord{5, 4}, wq)).To(BeFalse(), "captured own piece, and king in check")
		Expect(b.MakeMove(rect.Coord{3, 2}, wq)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(rect.Coord{1, 6}, wq)).To(BeFalse(), "jumped over own piece, and check")

		// successful capture releasing check
		Expect(b.MakeMove(rect.Coord{2, 5}, wq)).To(BeTrue(), "can't capture releasing check")
	})

	It("makes legal moves", func() {
		var wq, br, wr, bn base.IPiece
		testReset := func() {
			b = rect.NewEmptyBoard(w, h)
			wq, br = piece.NewQueen(White), piece.NewRook(Black)
			wr, bn = piece.NewRook(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{4, 6}, wq)
			b.PlacePiece(rect.Coord{1, 3}, br)
			b.PlacePiece(rect.Coord{2, 4}, wr)
			b.PlacePiece(rect.Coord{4, 4}, bn)
		}
		testReset()
		destinations := wq.Destinations(b)

		wqCoord, bnCoord := wq.Coord().Copy(), bn.Coord().Copy()
		Expect(destinations.Len()).To(Equal(8))
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wq)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wqCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wq))
			if !bnCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(bnCoord)).To(Equal(bn))
			} else {
				// capturing piece's coords is destination
				Expect(wq.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(bn.Coord()).To(BeNil())
			}

			Expect(b.Piece(wr.Coord())).To(Equal(wr))
			Expect(b.Piece(br.Coord())).To(Equal(br))

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wq, br, wr, bn base.IPiece
		testReset := func() {
			b = rect.NewEmptyBoard(w, h)
			wq, br = piece.NewQueen(White), piece.NewRook(Black)
			wr, bn = piece.NewRook(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{4, 6}, wq)
			b.PlacePiece(rect.Coord{1, 3}, br)
			b.PlacePiece(rect.Coord{2, 4}, wr)
			b.PlacePiece(rect.Coord{4, 4}, bn)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{rect.Coord{2, 4}, rect.Coord{1, 3}, rect.Coord{4, 3}, wq.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wq.Coord()
			Expect(b.MakeMove(d, wq)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wq))

			// check that destination cell was not changed
			p := b.Piece(d)
			if p == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(b.Piece(d)))
			}

			// check another cell to contain another piece
			Expect(b.Piece(br.Coord())).To(Equal(br))
			Expect(b.Piece(bn.Coord())).To(Equal(bn))
			Expect(b.Piece(wr.Coord())).To(Equal(wr))

			testReset()
		}
	})
})
