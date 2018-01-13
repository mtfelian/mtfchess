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

var _ = Describe("Rook test", func() {
	var b base.IBoard
	resetBoard := func() { b = rect.NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wr, wn, br := piece.NewRook(White), piece.NewKnight(White), piece.NewRook(Black)
		b.PlacePiece(rect.Coord{4, 2}, wr)
		b.PlacePiece(rect.Coord{4, 4}, wn)
		b.PlacePiece(rect.Coord{4, 1}, br)

		d := wr.Destinations(b)
		Expect(d.Len()).To(Equal(6))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{4, 1}, rect.Coord{1, 2}, rect.Coord{2, 2},
			rect.Coord{3, 2}, rect.Coord{5, 2}, rect.Coord{4, 3},
		}))).To(BeTrue())
	})

	It("attacks right cells", func() {
		wr, bq := piece.NewRook(White), piece.NewQueen(Black)
		bn, wn, wk := piece.NewKnight(Black), piece.NewKnight(White), piece.NewKing(White)
		b.PlacePiece(rect.Coord{1, 2}, bq)
		b.PlacePiece(rect.Coord{2, 2}, wr)
		b.PlacePiece(rect.Coord{3, 2}, wk)
		b.PlacePiece(rect.Coord{2, 1}, bn)
		b.PlacePiece(rect.Coord{2, 5}, wn)

		attacking := wr.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{1, 2}, rect.Coord{3, 2},
			rect.Coord{2, 3}, rect.Coord{2, 4}, rect.Coord{2, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(rect.Coord{2, 5}, wr)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(rect.Coord{3, 2}, wr)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(rect.Coord{2, 6}, wr)).To(BeFalse(), "jumped over own piece")
		Expect(b.MakeMove(rect.Coord{1, 2}, wr)).To(BeTrue(), "can't capture")
	})

	It("makes legal moves", func() {
		var wr, br base.IPiece
		testReset := func() {
			resetBoard()
			wr, br = piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 1}, wr)
			b.PlacePiece(rect.Coord{4, 1}, br)
		}
		testReset()

		destinations := wr.Destinations(b)
		sort.Sort(destinations)
		Expect(destinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{1, 1}, rect.Coord{3, 1}, rect.Coord{4, 1},
			rect.Coord{2, 2}, rect.Coord{2, 3}, rect.Coord{2, 4},
			rect.Coord{2, 5}, rect.Coord{2, 6},
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
			wr, br = piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 1}, wr)
			b.PlacePiece(rect.Coord{4, 1}, br)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{rect.Coord{3, 2}, rect.Coord{5, 1}, wr.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wr.Coord()
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
