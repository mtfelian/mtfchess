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

var _ = Describe("King test", func() {
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("generates moves", func() {
		wk, wn, bn := piece.NewKing(White), piece.NewKnight(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 2}, wk)
		b.PlacePiece(rect.Coord{2, 3}, wn)
		b.PlacePiece(rect.Coord{1, 1}, bn)
		d := wk.Destinations(b)
		Expect(d.Len()).To(Equal(6))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{1, 1}, rect.Coord{2, 1}, rect.Coord{3, 1},
			rect.Coord{1, 2}, rect.Coord{1, 3}, rect.Coord{3, 3},
		}))).To(BeTrue())
	})

	It("makes legal moves", func() {
		var wk, bk, br base.IPiece
		var boardCopy base.IBoard
		testReset := func() {
			wk, bk, br = piece.NewKing(White), piece.NewKing(Black), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{2, 3}, wk)
			b.PlacePiece(rect.Coord{3, 4}, br)
			b.PlacePiece(rect.Coord{4, 2}, bk)
			if boardCopy != nil && !reflect.ValueOf(boardCopy).IsNil() {
				b.Set(boardCopy)
			}
		}
		testReset()
		boardCopy = b.Copy()
		destinations := wk.Destinations(b)

		wkCoord, brCoord := wk.Coord().Copy(), br.Coord().Copy()
		fmt.Println(destinations)
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wk)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wkCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wk))
			fmt.Println("%%", wk.Coord(), br.Coord())
			if !brCoord.Equals(d) { // if not capture
				// then there should be another piece
				Expect(b.Piece(brCoord)).To(Equal(br))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wk.Coord()).To(Equal(d))
				// captured piece's coords is nil
				//Expect(br.Coord()).To(BeNil()) todo fix
			}

			testReset()
		}
	})

	It("can't go under check of rook", func() {
		wk, bk, br := piece.NewKing(White), piece.NewKing(Black), piece.NewRook(Black)
		b.PlacePiece(rect.Coord{2, 4}, wk)
		b.PlacePiece(rect.Coord{4, 3}, bk)
		b.PlacePiece(rect.Coord{4, 5}, br)

		d, c := rect.Coord{2, 5}, wk.Coord()
		Expect(b.MakeMove(d, wk)).To(BeFalse(), "gone under rook's check!")
		Expect(b.Piece(c)).To(Equal(wk))
		Expect(b.Piece(d)).To(BeNil())
		Expect(b.Piece(br.Coord())).To(Equal(br))
		Expect(b.Piece(bk.Coord())).To(Equal(bk))

		wkDestinations := wk.Destinations(b)
		Expect(wkDestinations.Len()).To(Equal(3))
		sort.Sort(wkDestinations)
		Expect(wkDestinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{1, 3}, rect.Coord{2, 3}, rect.Coord{1, 4},
		}))).To(BeTrue())
	})

	It("can't go to a opponent's king neighbour cell", func() {
		wk, bk, br := piece.NewKing(White), piece.NewKing(Black), piece.NewRook(Black)
		b.PlacePiece(rect.Coord{2, 4}, wk)
		b.PlacePiece(rect.Coord{4, 3}, bk)
		b.PlacePiece(rect.Coord{4, 5}, br)

		d, c := rect.Coord{3, 4}, wk.Coord()
		Expect(b.MakeMove(d, wk)).To(BeFalse(), "two kings on a neighbour cells!")
		Expect(b.Piece(c)).To(Equal(wk))
		Expect(b.Piece(d)).To(BeNil())
		Expect(b.Piece(br.Coord())).To(Equal(br))
		Expect(b.Piece(bk.Coord())).To(Equal(bk))
	})

	It("can't go into same cell (skip move)", func() {
		wk, bk, br := piece.NewKing(White), piece.NewKing(Black), piece.NewRook(Black)
		b.PlacePiece(rect.Coord{2, 4}, wk)
		b.PlacePiece(rect.Coord{4, 3}, bk)
		b.PlacePiece(rect.Coord{4, 5}, br)

		c := wk.Coord()
		Expect(b.MakeMove(c, wk)).To(BeFalse(), "performed a move into same cell!")
		Expect(b.Piece(c)).To(Equal(wk))
		Expect(b.Piece(br.Coord())).To(Equal(br))
		Expect(b.Piece(bk.Coord())).To(Equal(bk))
	})
})
