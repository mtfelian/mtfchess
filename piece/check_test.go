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
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("is white in check", func() {
		wn, bn := piece.NewKnight(White), piece.NewKnight(Black)
		wk, bk := piece.NewKing(White), piece.NewKing(Black)
		b.PlacePiece(rect.Coord{X: 1, Y: 1}, wk)
		b.PlacePiece(rect.Coord{X: 3, Y: 2}, bn)
		b.PlacePiece(rect.Coord{X: 5, Y: 4}, bk)
		b.PlacePiece(rect.Coord{X: 4, Y: 4}, wn)

		Expect(piece.InCheck(b, White)).To(BeTrue())
		Expect(piece.InCheck(b, Black)).To(BeFalse())
	})

	It("is black in check", func() {
		wn, bn := piece.NewKnight(White), piece.NewKnight(Black)
		wk, bk := piece.NewKing(White), piece.NewKing(Black)
		b.PlacePiece(rect.Coord{X: 1, Y: 1}, bk)
		b.PlacePiece(rect.Coord{X: 3, Y: 2}, wn)
		b.PlacePiece(rect.Coord{X: 5, Y: 4}, wk)
		b.PlacePiece(rect.Coord{X: 4, Y: 4}, bn)

		Expect(piece.InCheck(b, White)).To(BeFalse())
		Expect(piece.InCheck(b, Black)).To(BeTrue())
	})
})
