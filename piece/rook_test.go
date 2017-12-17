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
		wr := piece.NewRook(White)
		wn := piece.NewKnight(White)
		br := piece.NewRook(Black)
		b.PlacePiece(rect.Coord{X: 4, Y: 2}, wr)
		b.PlacePiece(rect.Coord{X: 4, Y: 4}, wn)
		b.PlacePiece(rect.Coord{X: 4, Y: 1}, br)

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
		var boardCopy base.IBoard
		testReset := func() {
			wr, br = piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{X: 2, Y: 1}, wr)
			b.PlacePiece(rect.Coord{X: 4, Y: 2}, br)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()
		boardCopy = b.Copy()
		destinations := wr.Destinations(b)

		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			c := wr.Coord()
			Expect(b.MakeMove(d, wr)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(c)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wr))
			if !br.Coord().Equals(d) { // if not capture
				// then there should be another piece
				Expect(b.Piece(br.Coord())).To(Equal(br))
			}

			testReset()
		}
	})

	It("makes illegal moves", func() {
		var wr, br base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wr, br = piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{X: 2, Y: 1}, wr)
			b.PlacePiece(rect.Coord{X: 4, Y: 1}, br)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()

		boardCopy = b.Copy()
		destinations := rect.NewCoords([]base.ICoord{rect.Coord{3, 2}, rect.Coord{5, 1}})
		for destinations.HasNext() {
			d := destinations.Next().(rect.Coord)
			c := wr.Coord()
			Expect(b.MakeMove(d, wr)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wr))
			// check destination cell to be empty
			Expect(b.Piece(d)).To(BeNil())
			// check another cell to contain another piece
			Expect(b.Piece(br.Coord())).To(Equal(br))

			testReset()
		}
	})
})
