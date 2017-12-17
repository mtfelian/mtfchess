package piece_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("King test", func() {
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("generates moves", func() {
		wk, wn, bn := piece.NewKing(White), piece.NewKnight(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{X: 2, Y: 2}, wk)
		b.PlacePiece(rect.Coord{X: 2, Y: 3}, wn)
		b.PlacePiece(rect.Coord{X: 1, Y: 1}, bn)
		d := wk.Destinations(b)
		Expect(d.Len()).To(Equal(6))
		Expect(d.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{1, 1}, rect.Coord{1, 2}, rect.Coord{1, 3},
			rect.Coord{2, 1}, rect.Coord{3, 1}, rect.Coord{3, 3},
		}))).To(BeTrue())
	})
})
