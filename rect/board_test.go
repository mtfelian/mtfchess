package rect_test

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board test", func() {
	w, h := 5, 6
	var b base.IBoard

	BeforeEach(func() {
		b = rect.NewEmptyBoard(w, h, base.Settings{PawnLongFunc: rect.NoPawnLongMoveFunc})
	})

	It("checks board width and height", func() {
		Expect(b.Dim().(rect.Coord).X).To(Equal(w))
		Expect(b.Dim().(rect.Coord).Y).To(Equal(h))
	})

	Describe("find pieces", func() {
		var wn1, wn2, wn3, bn1, bn2, bn3, wk, bk base.IPiece
		BeforeEach(func() {
			wn1, wn2, wn3 = piece.NewKnight(White), piece.NewKnight(White), piece.NewKnight(White)
			bn1, bn2, bn3 = piece.NewKnight(Black), piece.NewKnight(Black), piece.NewKnight(Black)
			wk, bk = piece.NewKing(White), piece.NewKing(Black)
			b.PlacePiece(rect.Coord{1, 1}, wn1)
			b.PlacePiece(rect.Coord{1, 2}, wn2)
			b.PlacePiece(rect.Coord{3, 4}, wn3)
			b.PlacePiece(rect.Coord{5, 5}, bn1)
			b.PlacePiece(rect.Coord{5, 6}, bn2)
			b.PlacePiece(rect.Coord{4, 3}, bn3)
			b.PlacePiece(rect.Coord{2, 1}, wk)
			b.PlacePiece(rect.Coord{5, 4}, bk)
		})
		It("normally", func() {
			filter := rect.PieceFilter{ // find all white knights
				PieceFilter: base.PieceFilter{
					Colours: []Colour{White},
					Names:   []string{piece.NewKnight(Transparent).Name()},
				},
			}
			coords := b.FindPieces(filter)
			Expect(coords).To(HaveLen(3))
			Expect(coords).To(Equal(base.Pieces{wn3, wn2, wn1}))
		})

		It("is with piece / board condition", func() {
			notOnEdge := func(p base.IPiece) bool {
				x, y := p.Coord().(rect.Coord).X, p.Coord().(rect.Coord).Y
				w, h := b.Dim().(rect.Coord).X, b.Dim().(rect.Coord).Y
				return x > 1 && y > 1 && x < w && y < h
			}
			filter := rect.PieceFilter{ // find all knights
				PieceFilter: base.PieceFilter{
					Names:     []string{piece.NewKnight(Transparent).Name()},
					Condition: notOnEdge,
				},
			}

			pieces := b.FindPieces(filter)
			Expect(pieces).To(HaveLen(2))
			Expect(pieces).To(Equal(base.Pieces{wn3, bn3}))
		})
	})

	Describe("find attacked cells", func() {
		It("works", func() {
			wn, bn := piece.NewKnight(White), piece.NewKnight(Black)
			wk, bk := piece.NewKing(White), piece.NewKing(Black)
			b.PlacePiece(rect.Coord{1, 1}, bk)
			b.PlacePiece(rect.Coord{2, 4}, wn)
			b.PlacePiece(rect.Coord{5, 5}, wk)
			b.PlacePiece(rect.Coord{4, 4}, bn)

			attackedByWhite := b.FindAttackedCellsBy(rect.PieceFilter{
				PieceFilter: base.PieceFilter{Colours: []Colour{White}},
			})
			Expect(attackedByWhite.Len()).To(Equal(10))
			sort.Sort(attackedByWhite)
			Expect(attackedByWhite.Equals(rect.NewCoords([]base.ICoord{
				rect.Coord{1, 2}, rect.Coord{3, 2}, rect.Coord{4, 3},
				rect.Coord{4, 4}, rect.Coord{5, 4}, rect.Coord{4, 5},
				rect.Coord{1, 6}, rect.Coord{3, 6}, rect.Coord{4, 6}, rect.Coord{5, 6},
			}))).To(BeTrue())

			attackedByBlack := b.FindAttackedCellsBy(rect.PieceFilter{
				PieceFilter: base.PieceFilter{Colours: []Colour{Black}},
			})
			Expect(attackedByBlack.Len()).To(Equal(9))
			sort.Sort(attackedByBlack)
			Expect(attackedByBlack.Equals(rect.NewCoords([]base.ICoord{
				rect.Coord{2, 1}, rect.Coord{1, 2}, rect.Coord{2, 2},
				rect.Coord{3, 2}, rect.Coord{5, 2}, rect.Coord{2, 3},
				rect.Coord{2, 5}, rect.Coord{3, 6}, rect.Coord{5, 6},
			}))).To(BeTrue())
		})
	})
})
