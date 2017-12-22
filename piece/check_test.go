package piece_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check test", func() {
	var b base.IBoard
	BeforeEach(func() { b = rect.NewTestEmptyBoard() })

	It("is white in check", func() {
		wn, bn := piece.NewKnight(White), piece.NewKnight(Black)
		wk, bk := piece.NewKing(White), piece.NewKing(Black)
		b.PlacePiece(rect.Coord{1, 1}, wk)
		b.PlacePiece(rect.Coord{3, 2}, bn)
		b.PlacePiece(rect.Coord{5, 4}, bk)
		b.PlacePiece(rect.Coord{4, 4}, wn)

		Expect(piece.InCheck(b, White)).To(BeTrue())
		Expect(piece.InCheck(b, Black)).To(BeFalse())
	})

	It("is black in check", func() {
		wn, bn := piece.NewKnight(White), piece.NewKnight(Black)
		wk, bk := piece.NewKing(White), piece.NewKing(Black)
		b.PlacePiece(rect.Coord{1, 1}, bk)
		b.PlacePiece(rect.Coord{3, 2}, wn)
		b.PlacePiece(rect.Coord{5, 4}, wk)
		b.PlacePiece(rect.Coord{4, 4}, bn)

		Expect(piece.InCheck(b, White)).To(BeFalse())
		Expect(piece.InCheck(b, Black)).To(BeTrue())
	})

	It("can't expose check", func() {
		wn, br, wk := piece.NewKnight(White), piece.NewRook(Black), piece.NewKing(White)
		b.PlacePiece(rect.Coord{3, 2}, br)
		b.PlacePiece(rect.Coord{4, 2}, wn)
		b.PlacePiece(rect.Coord{5, 2}, wk)

		d, c := rect.Coord{2, 1}, wn.Coord().Copy()
		Expect(b.MakeMove(d, wn)).To(BeFalse(), "check exposed!")
		Expect(b.Piece(c)).To(Equal(wn))
		Expect(b.Piece(d)).To(BeNil())
		Expect(b.Piece(br.Coord())).To(Equal(br))

		// white knight should have no possible moves
		Expect(wn.Destinations(b).Len()).To(Equal(0))
	})

	It("can capture at pin", func() {
		wr, bn, wk := piece.NewRook(White), piece.NewKnight(Black), piece.NewKing(White)
		b.PlacePiece(rect.Coord{3, 2}, bn)
		b.PlacePiece(rect.Coord{4, 2}, wr)
		b.PlacePiece(rect.Coord{5, 2}, wk)

		d, c := rect.Coord{3, 2}, wr.Coord()
		Expect(b.MakeMove(d, wr)).To(BeTrue(), "can't capture at pin!")
		Expect(b.Piece(c)).To(BeNil())
		Expect(b.Piece(d)).To(Equal(wr))

		// coords of captured piece should became nil
		Expect(bn.Coord()).To(BeNil())
	})

})
