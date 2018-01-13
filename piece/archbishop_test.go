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

var _ = Describe("Archbishop test", func() {
	var b base.IBoard
	resetBoard := func() { b = rect.NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wa, wn, ba := piece.NewArchbishop(White), piece.NewKnight(White), piece.NewArchbishop(Black)
		b.PlacePiece(rect.Coord{4, 2}, wa)
		b.PlacePiece(rect.Coord{4, 5}, wn)
		b.PlacePiece(rect.Coord{1, 2}, ba)

		wbDestinations := wa.Destinations(b)
		sort.Sort(wbDestinations)
		Expect(wbDestinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{3, 1}, rect.Coord{5, 1},
			rect.Coord{2, 3}, rect.Coord{3, 3}, rect.Coord{5, 3},
			rect.Coord{2, 4}, rect.Coord{3, 4}, rect.Coord{5, 4}, rect.Coord{1, 5},
		}))).To(BeTrue())

		bbDestinations := ba.Destinations(b)
		sort.Sort(bbDestinations)
		Expect(bbDestinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{3, 1}, rect.Coord{2, 3}, rect.Coord{3, 3},
			rect.Coord{2, 4}, rect.Coord{3, 4}, rect.Coord{4, 5},
		}))).To(BeTrue())
	})

	It("attacks right cells, can release check by capture", func() {
		wa, bq := piece.NewArchbishop(White), piece.NewQueen(Black)
		bn, wr, wk := piece.NewKnight(Black), piece.NewRook(White), piece.NewKing(White)
		b.PlacePiece(rect.Coord{2, 5}, bq)
		b.PlacePiece(rect.Coord{4, 3}, wa)
		b.PlacePiece(rect.Coord{2, 3}, wk)
		b.PlacePiece(rect.Coord{5, 1}, bn)
		b.PlacePiece(rect.Coord{5, 4}, wr)

		attacking := wa.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{3, 1}, rect.Coord{5, 1},
			rect.Coord{2, 2}, rect.Coord{3, 2}, rect.Coord{5, 2},
			rect.Coord{2, 4}, rect.Coord{3, 4}, rect.Coord{5, 4},
			rect.Coord{2, 5}, rect.Coord{3, 5}, rect.Coord{5, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(rect.Coord{5, 4}, wa)).To(BeFalse(), "captured own piece, and king in check")
		Expect(b.MakeMove(rect.Coord{3, 1}, wa)).To(BeFalse(), "king still in check")
		Expect(b.MakeMove(rect.Coord{1, 6}, wa)).To(BeFalse(), "jumped over own piece, and check")

		// successful capture releasing check
		boardCopy := b.Copy()
		Expect(b.MakeMove(rect.Coord{2, 5}, wa)).To(BeTrue(), "can't capture releasing check")

		// successful pin the check
		b.Set(boardCopy)
		wa.Set(boardCopy.Piece(rect.Coord{4, 3}))
		Expect(b.MakeMove(rect.Coord{2, 4}, wa)).To(BeTrue(), "can't pin releasing check")
	})

	It("makes legal moves", func() {
		var wa, br base.IPiece
		testReset := func() {
			resetBoard()
			wa, br = piece.NewArchbishop(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 3}, wa)
			b.PlacePiece(rect.Coord{1, 2}, br)
		}
		testReset()
		destinations := wa.Destinations(b)

		wbCoord, brCoord := wa.Coord().Copy(), br.Coord().Copy()
		Expect(destinations.Len()).To(Equal(13))
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wa)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wbCoord)).To(BeNil())
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
			wa, br = piece.NewArchbishop(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 3}, wa)
			b.PlacePiece(rect.Coord{4, 5}, br)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{rect.Coord{5, 6}, rect.Coord{2, 2}, wa.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wa.Coord()
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
