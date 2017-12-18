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

var _ = Describe("Bishop test", func() {
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("generates moves", func() {
		wb, wn, bb := piece.NewBishop(White), piece.NewKnight(White), piece.NewBishop(Black)
		b.PlacePiece(rect.Coord{4, 2}, wb)
		b.PlacePiece(rect.Coord{4, 5}, wn)
		b.PlacePiece(rect.Coord{1, 2}, bb)

		wbDestinations := wb.Destinations(b)
		Expect(wbDestinations.Len()).To(Equal(6))
		sort.Sort(wbDestinations)
		Expect(wbDestinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{3, 1}, rect.Coord{5, 1}, rect.Coord{3, 3},
			rect.Coord{5, 3}, rect.Coord{2, 4}, rect.Coord{1, 5},
		}))).To(BeTrue())

		bbDestinations := bb.Destinations(b)
		Expect(bbDestinations.Len()).To(Equal(4))
		sort.Sort(bbDestinations)
		Expect(bbDestinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{2, 3}, rect.Coord{3, 4}, rect.Coord{4, 5},
		}))).To(BeTrue())
	})

	It("makes legal moves", func() {
		var wb, br base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wb, br = piece.NewBishop(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 3}, wb)
			b.PlacePiece(rect.Coord{1, 2}, br)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()
		boardCopy = b.Copy()
		destinations := wb.Destinations(b)

		Expect(destinations.Len()).To(Equal(7))
		for destinations.HasNext() {
			d, c := destinations.Next().(base.ICoord), wb.Coord()
			Expect(b.MakeMove(d, wb)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(c)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wb))
			if !br.Coord().Equals(d) { // if not capture
				// then there should be another piece
				Expect(b.Piece(br.Coord())).To(Equal(br))
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wb, br base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wb, br = piece.NewBishop(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 3}, wb)
			b.PlacePiece(rect.Coord{4, 5}, br)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()

		boardCopy = b.Copy()
		destinations := rect.NewCoords([]base.ICoord{rect.Coord{5, 6}, rect.Coord{2, 2}, wb.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wb.Coord()
			Expect(b.MakeMove(d, wb)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wb))

			// check that destination cell was not changed
			if boardCopy.Piece(d) == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(boardCopy.Piece(d)))
			}

			// check another cell to contain another piece
			Expect(b.Piece(br.Coord())).To(Equal(br))

			testReset()
		}
	})
})
