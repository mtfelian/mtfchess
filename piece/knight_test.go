package piece_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Knight test", func() {
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("generates moves", func() {
		wn1, wn2, bn := piece.NewKnight(White), piece.NewKnight(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 1}, wn1)
		b.PlacePiece(rect.Coord{3, 3}, wn2)
		b.PlacePiece(rect.Coord{4, 2}, bn)

		d := wn1.Destinations(b)
		Expect(d.Len()).To(Equal(2))
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{4, 2}}))).To(BeTrue())
	})

	It("makes legal moves", func() {
		var wn, bn base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wn, bn = piece.NewKnight(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 1}, wn)
			b.PlacePiece(rect.Coord{4, 2}, bn)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()
		boardCopy = b.Copy()
		destinations := wn.Destinations(b)

		for destinations.HasNext() {
			d, c := destinations.Next().(base.ICoord), wn.Coord()
			Expect(b.MakeMove(d, wn)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(c)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wn))
			if !bn.Coord().Equals(d) { // if not capture
				// then there should be another piece
				Expect(b.Piece(bn.Coord())).To(Equal(bn))
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wn, bn base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wn, bn = piece.NewKnight(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 1}, wn)
			b.PlacePiece(rect.Coord{4, 2}, bn)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()

		boardCopy = b.Copy()
		destinations := rect.NewCoords([]base.ICoord{rect.Coord{5, 2}, rect.Coord{1, 4}, wn.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wn.Coord()
			Expect(b.MakeMove(d, wn)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wn))

			// check that destination cell was not changed
			if boardCopy.Piece(d) == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(boardCopy.Piece(d)))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bn.Coord())).To(Equal(bn))

			testReset()
		}
	})
})
