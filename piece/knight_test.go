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
	"reflect"
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
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{4, 2}, rect.Coord{1, 3}}))).To(BeTrue())
	})

	It("makes legal moves", func() {
		var wn, bq base.IPiece
		testReset := func() {
			b = rect.NewEmptyBoard(w, h)
			wn, bq = piece.NewKnight(White), piece.NewQueen(Black)
			b.PlacePiece(rect.Coord{2, 1}, wn)
			b.PlacePiece(rect.Coord{4, 2}, bq)
		}
		testReset()
		destinations := wn.Destinations(b)
		sort.Sort(destinations)
		Expect(destinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{4, 2}, rect.Coord{1, 3}, rect.Coord{3, 3},
		}))).To(BeTrue())

		bqCoord, wnCoord := bq.Coord().Copy(), wn.Coord().Copy()
		for destinations.HasNext() {
			fmt.Println(b)
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wn)).To(BeTrue(), "failed at destination %d", destinations.I())
			fmt.Println(b)
			// check source cell to be empty
			Expect(b.Piece(wnCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wn))
			if !bqCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(bqCoord)).To(Equal(bq))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wn.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(bq.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wn, bn base.IPiece
		testReset := func() {
			wn, bn = piece.NewKnight(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 1}, wn)
			b.PlacePiece(rect.Coord{4, 2}, bn)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{rect.Coord{5, 2}, rect.Coord{1, 4}, wn.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wn.Coord()
			Expect(b.MakeMove(d, wn)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wn))

			// check that destination cell was not changed
			p := b.Piece(d)
			if p == nil || reflect.ValueOf(p).IsNil() {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(b.Piece(d)))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bn.Coord())).To(Equal(bn))

			testReset()
		}
	})
})
