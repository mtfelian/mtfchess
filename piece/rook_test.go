package piece_test

import (
	"sort"

	"fmt"
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

	FIt("makes legal moves", func() {
		var wr, br base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wr, br = piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 1}, wr)
			b.PlacePiece(rect.Coord{4, 1}, br)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()
		boardCopy = b.Copy()
		destinations := wr.Destinations(b)

		brCoord := br.Coord().Copy()
		for destinations.HasNext() {
			d, c := destinations.Next().(base.ICoord), wr.Coord()
			Expect(b.MakeMove(d, wr)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(c)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wr))
			fmt.Println("@@", wr.Coord(), br.Coord())
			if !brCoord.Equals(d) { // if not capture
				// then there should be another piece
				Expect(b.Piece(br.Coord())).To(Equal(br))
			} else {
				Expect(br.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wr, br base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wr, br = piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 1}, wr)
			b.PlacePiece(rect.Coord{4, 1}, br)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()

		boardCopy = b.Copy()
		destinations := rect.NewCoords([]base.ICoord{rect.Coord{3, 2}, rect.Coord{5, 1}, wr.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wr.Coord()
			Expect(b.MakeMove(d, wr)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wr))

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
