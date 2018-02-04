package rect

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Archbishop test", func() {
	var b base.IBoard
	resetBoard := func() { b = NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wa, wn, ba := NewArchbishop(White), NewKnight(White), NewArchbishop(Black)
		b.PlacePiece(Coord{4, 2}, wa)
		b.PlacePiece(Coord{4, 5}, wn)
		b.PlacePiece(Coord{1, 2}, ba)

		wbDestinations := wa.Destinations(b)
		sort.Sort(wbDestinations)
		Expect(wbDestinations.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{3, 1}, Coord{5, 1},
			Coord{2, 3}, Coord{3, 3}, Coord{5, 3},
			Coord{2, 4}, Coord{3, 4}, Coord{5, 4}, Coord{1, 5},
		}))).To(BeTrue())

		bbDestinations := ba.Destinations(b)
		sort.Sort(bbDestinations)
		Expect(bbDestinations.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{3, 1}, Coord{2, 3}, Coord{3, 3},
			Coord{2, 4}, Coord{3, 4}, Coord{4, 5},
		}))).To(BeTrue())
	})

	It("attacks right cells, can release check by capture", func() {
		wa, bq := NewArchbishop(White), NewQueen(Black)
		bn, wr, wk := NewKnight(Black), NewRook(White), NewKing(White)
		b.PlacePiece(Coord{2, 5}, bq)
		b.PlacePiece(Coord{4, 3}, wa)
		b.PlacePiece(Coord{2, 3}, wk)
		b.PlacePiece(Coord{5, 1}, bn)
		b.PlacePiece(Coord{5, 4}, wr)

		attacking := wa.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{2, 1}, Coord{3, 1}, Coord{5, 1},
			Coord{2, 2}, Coord{3, 2}, Coord{5, 2},
			Coord{2, 4}, Coord{3, 4}, Coord{5, 4},
			Coord{2, 5}, Coord{3, 5}, Coord{5, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{5, 4}, wa)).To(BeFalse(), "captured own piece, and king in check")
		Expect(b.MakeMove(Coord{3, 1}, wa)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(Coord{1, 6}, wa)).To(BeFalse(), "jumped over own piece, and check")

		// successful capture releasing check
		boardCopy := b.Copy()
		Expect(b.MakeMove(Coord{2, 5}, wa)).To(BeTrue(), "can't capture releasing check")

		// successful pin the check
		b.Set(boardCopy)
		wa.Set(boardCopy.Piece(Coord{4, 3}))
		Expect(b.MakeMove(Coord{2, 4}, wa)).To(BeTrue(), "can't pin releasing check")
	})

	It("makes legal moves", func() {
		var wa, br base.IPiece
		testReset := func() {
			resetBoard()
			wa, br = NewArchbishop(White), NewRook(Black)
			b.PlacePiece(Coord{2, 3}, wa)
			b.PlacePiece(Coord{1, 2}, br)
		}
		testReset()
		destinations := wa.Destinations(b)

		waCoord, brCoord := wa.Coord().Copy(), br.Coord().Copy()
		Expect(destinations.Len()).To(Equal(13))
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wa)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(waCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wa))
			if !brCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(brCoord)).To(Equal(br))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wa.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(br.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wa, br base.IPiece
		testReset := func() {
			resetBoard()
			wa, br = NewArchbishop(White), NewRook(Black)
			b.PlacePiece(Coord{2, 3}, wa)
			b.PlacePiece(Coord{4, 5}, br)
		}
		testReset()

		destinations := NewCoords([]base.ICoord{Coord{5, 6}, Coord{2, 2}, wa.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(Coord), wa.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wa)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wa))

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
