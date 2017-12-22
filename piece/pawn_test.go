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

var _ = Describe("Pawn test", func() {
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("generates moves", func() {
		wp, bn := piece.NewPawn(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bn)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}})))

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(2))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{2, 3}}))).To(BeTrue())
	})

	It("attacks right cells, can release check by capture", func() {
		wp, bn, wk := piece.NewPawn(White), piece.NewKnight(Black), piece.NewKing(White)
		wp2 := piece.NewPawn(White)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bn)
		b.PlacePiece(rect.Coord{3, 4}, wk)
		b.PlacePiece(rect.Coord{3, 3}, wp2)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}})))

		Expect(b.MakeMove(rect.Coord{3, 3}, wp)).To(BeFalse(), "captured own piece, and king in check")
		Expect(b.MakeMove(rect.Coord{2, 3}, wp)).To(BeFalse(), "king still in check")
		// successful capture releasing check
		Expect(b.MakeMove(rect.Coord{1, 3}, wp)).To(BeTrue(), "can't capture releasing check")
	})
})
