package rect

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queen test", func() {
	var b base.IBoard
	resetBoard := func() { b = NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wq, wn, bq := NewQueen(White), NewKnight(White), NewQueen(Black)
		b.PlacePiece(Coord{2, 3}, wq)
		b.PlacePiece(Coord{1, 2}, wn)
		b.PlacePiece(Coord{3, 3}, bq)

		wqDestinations := wq.Destinations(b)
		sort.Sort(wqDestinations)
		Expect(wqDestinations.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{4, 1}, Coord{2, 2}, Coord{3, 2},
			Coord{1, 3}, Coord{3, 3}, Coord{1, 4}, Coord{2, 4},
			Coord{3, 4}, Coord{2, 5}, Coord{4, 5}, Coord{2, 6},
			Coord{5, 6},
		}))).To(BeTrue())

		bqDestinations := bq.Destinations(b)
		sort.Sort(bqDestinations)
		Expect(bqDestinations.Equals(NewCoords([]base.ICoord{
			Coord{1, 1}, Coord{3, 1}, Coord{5, 1}, Coord{2, 2},
			Coord{3, 2}, Coord{4, 2}, Coord{2, 3}, Coord{4, 3},
			Coord{5, 3}, Coord{2, 4}, Coord{3, 4}, Coord{4, 4},
			Coord{1, 5}, Coord{3, 5}, Coord{5, 5}, Coord{3, 6},
		}))).To(BeTrue())
	})

	It("attacks right cells", func() {
		wq, bq := NewQueen(White), NewQueen(Black)
		bn, wr, wk := NewKnight(Black), NewRook(White), NewKing(White)
		b.PlacePiece(Coord{2, 5}, bq)
		b.PlacePiece(Coord{4, 3}, wq)
		b.PlacePiece(Coord{2, 3}, wk)
		b.PlacePiece(Coord{5, 1}, bn)
		b.PlacePiece(Coord{5, 4}, wr)

		attacking := wq.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{4, 1}, Coord{3, 2}, Coord{4, 2},
			Coord{5, 2}, Coord{2, 3}, Coord{3, 3}, Coord{5, 3},
			Coord{3, 4}, Coord{4, 4}, Coord{5, 4}, Coord{2, 5},
			Coord{4, 5}, Coord{4, 6},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{5, 4}, wq)).To(BeFalse(), "captured own piece, and king in check")
		Expect(b.MakeMove(Coord{3, 2}, wq)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(Coord{1, 6}, wq)).To(BeFalse(), "jumped over own piece, and check")

		// successful capture releasing check
		Expect(b.MakeMove(Coord{2, 5}, wq)).To(BeTrue(), "can't capture releasing check")
	})

	It("makes legal moves", func() {
		var wq, br, wr, bn base.IPiece
		testReset := func() {
			resetBoard()
			wq, br = NewQueen(White), NewRook(Black)
			wr, bn = NewRook(White), NewKnight(Black)
			b.PlacePiece(Coord{4, 6}, wq)
			b.PlacePiece(Coord{1, 3}, br)
			b.PlacePiece(Coord{2, 4}, wr)
			b.PlacePiece(Coord{4, 4}, bn)
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
			resetBoard()
			wq, br = NewQueen(White), NewRook(Black)
			wr, bn = NewRook(White), NewKnight(Black)
			b.PlacePiece(Coord{4, 6}, wq)
			b.PlacePiece(Coord{1, 3}, br)
			b.PlacePiece(Coord{2, 4}, wr)
			b.PlacePiece(Coord{4, 4}, bn)
		}
		testReset()

		destinations := NewCoords([]base.ICoord{Coord{2, 4}, Coord{1, 3}, Coord{4, 3}, wq.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(Coord), wq.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wq)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wq))

			// check that destination cell was not changed
			if dCellCopy.Piece() == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(dCellCopy.Piece()))
			}

			// check another cell to contain another piece
			Expect(b.Piece(br.Coord())).To(Equal(br))
			Expect(b.Piece(bn.Coord())).To(Equal(bn))
			Expect(b.Piece(wr.Coord())).To(Equal(wr))

			testReset()
		}
	})
})
