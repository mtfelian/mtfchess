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

	It("makes legal moves", func() {
		var wq, br, wr, bn base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wq, br = piece.NewQueen(White), piece.NewRook(Black)
			wr, bn = piece.NewRook(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{4, 6}, wq)
			b.PlacePiece(rect.Coord{1, 3}, br)
			b.PlacePiece(rect.Coord{2, 4}, wr)
			b.PlacePiece(rect.Coord{4, 4}, bn)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()
		boardCopy = b.Copy()
		destinations := wq.Destinations(b)

		bnCoord := bn.Coord().Copy()
		Expect(destinations.Len()).To(Equal(8))
		for destinations.HasNext() {
			d, c := destinations.Next().(base.ICoord), wq.Coord()
			Expect(b.MakeMove(d, wq)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(c)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wq))
			fmt.Println("$$", wq.Coord(), br.Coord())
			if !bnCoord.Equals(d) { // if not capture
				// then there should be another piece
				Expect(b.Piece(bn.Coord())).To(Equal(bn))
			} else {
				Expect(bn.Coord()).To(BeNil())
			}

			Expect(b.Piece(wr.Coord())).To(Equal(wr))

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wq, br, wr, bn base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wq, br = piece.NewQueen(White), piece.NewRook(Black)
			wr, bn = piece.NewRook(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{4, 6}, wq)
			b.PlacePiece(rect.Coord{1, 3}, br)
			b.PlacePiece(rect.Coord{2, 4}, wr)
			b.PlacePiece(rect.Coord{4, 4}, bn)
			if boardCopy != nil {
				b.Set(boardCopy)
			}
		}
		testReset()

		boardCopy = b.Copy()
		destinations := rect.NewCoords([]base.ICoord{rect.Coord{2, 4}, rect.Coord{1, 3}, rect.Coord{4, 3}, wq.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wq.Coord()
			Expect(b.MakeMove(d, wq)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wq))

			// check that destination cell was not changed
			if boardCopy.Piece(d) == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(boardCopy.Piece(d)))
			}

			// check another cell to contain another piece
			Expect(b.Piece(br.Coord())).To(Equal(br))
			Expect(b.Piece(bn.Coord())).To(Equal(bn))
			Expect(b.Piece(wr.Coord())).To(Equal(wr))

			testReset()
		}
	})
})
