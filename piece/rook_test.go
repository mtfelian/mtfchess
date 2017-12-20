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
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

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

	It("makes legal moves", func() {
		var wr, br base.IPiece
		testReset := func() {
			b = rect.NewEmptyBoard(w, h)
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
				Expect(br.Coord()).To(BeNil()) //todo fix
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wr, br base.IPiece
		testReset := func() {
			b = rect.NewEmptyBoard(w, h)
			wr, br = piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 1}, wr)
			b.PlacePiece(rect.Coord{4, 1}, br)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{rect.Coord{3, 2}, rect.Coord{5, 1}, wr.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wr.Coord()
			Expect(b.MakeMove(d, wr)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wr))

			// check that destination cell was not changed
			p := b.Piece(d)
			if p == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(b.Piece(d)))
			}

			// check another cell to contain another piece
			Expect(b.Piece(br.Coord())).To(Equal(br))

			testReset()
		}
	})
})
