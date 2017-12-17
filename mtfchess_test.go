package mtfchess

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

	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("checks board width and height", func() {
		Expect(b.Dim().(rect.Coord).X).To(Equal(w))
		Expect(b.Dim().(rect.Coord).Y).To(Equal(h))
	})

	Describe("knight", func() {
		It("generates moves", func() {
			wn1 := piece.NewKnight(White)
			wn2 := piece.NewKnight(White)
			bn := piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{X: 2, Y: 1}, wn1)
			b.PlacePiece(rect.Coord{X: 3, Y: 3}, wn2)
			b.PlacePiece(rect.Coord{X: 4, Y: 2}, bn)

			d := wn1.Destinations(b)
			Expect(d.Len()).To(Equal(2))
			Expect(d.Equals(rect.NewCoords([]base.Coord{rect.Coord{1, 3}, rect.Coord{4, 2}}))).To(BeTrue())
		})

		It("makes legal moves", func() {
			var wn, bn base.IPiece
			var boardCopy base.IBoard
			testReset := func() {
				wn, bn = piece.NewKnight(White), piece.NewKnight(Black)
				b.PlacePiece(rect.Coord{X: 2, Y: 1}, wn)
				b.PlacePiece(rect.Coord{X: 4, Y: 2}, bn)
				if boardCopy != nil {
					b.Set(boardCopy)
				}
			}
			testReset()
			boardCopy = b.Copy()
			destinations := wn.Destinations(b)

			for destinations.HasNext() {
				d := destinations.Next().(base.Coord)
				c := wn.Coord()
				Expect(b.MakeMove(d, wn)).To(BeTrue(), "failed at destination %d", destinations.I())
				// check source cell to be empty
				Expect(b.Piece(c)).To(BeNil())
				// check destination cell to contain new piece
				Expect(b.Piece(d)).To(Equal(wn))
				if !bn.Coord().Equals(d) { // if not capture
					// then there should be another piece
					Expect(b.Piece(bn.Coord())).To(Equal(bn))
				}

				testReset()
			}
		})

		It("makes illegal moves", func() {
			var wn, bn base.IPiece
			var boardCopy base.IBoard
			testReset := func() {
				wn, bn = piece.NewKnight(White), piece.NewKnight(Black)
				b.PlacePiece(rect.Coord{X: 2, Y: 1}, wn)
				b.PlacePiece(rect.Coord{X: 4, Y: 2}, bn)
				if boardCopy != nil {
					b.Set(boardCopy)
				}
			}
			testReset()
			boardCopy = b.Copy()
			offsets := rect.NewCoords([]base.Coord{rect.Coord{3, 1}, rect.Coord{-1, 3}})

			for offsets.HasNext() {
				o := offsets.Next().(base.Coord)
				c := wn.Coord()
				c1 := c.Add(o)
				Expect(b.MakeMove(c1, wn)).To(BeFalse(), "failed at offset %d", offsets.I())
				// check source cell to contain unmoved piece
				Expect(b.Piece(c)).To(Equal(wn))
				// check destination cell to be empty
				Expect(b.Piece(c1)).To(BeNil())
				// check another cell to contain another piece
				Expect(b.Piece(bn.Coord())).To(Equal(bn))

				testReset()
			}
		})
	})

	Describe("king", func() {
		It("generates moves", func() {
			wk := piece.NewKing(White)
			wn := piece.NewKnight(White)
			bn := piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{X: 2, Y: 2}, wk)
			b.PlacePiece(rect.Coord{X: 2, Y: 3}, wn)
			b.PlacePiece(rect.Coord{X: 1, Y: 1}, bn)
			d := wk.Destinations(b)
			Expect(d.Len()).To(Equal(6))
			Expect(d.Equals(rect.NewCoords([]base.Coord{
				rect.Coord{1, 1}, rect.Coord{1, 2}, rect.Coord{1, 3},
				rect.Coord{2, 1}, rect.Coord{3, 1}, rect.Coord{3, 3},
			}))).To(BeTrue())
		})
	})

	Describe("find pieces", func() {
		var wn1, wn2, wn3, bn1, bn2, bn3, wk, bk base.IPiece
		BeforeEach(func() {
			wn1, wn2, wn3 = piece.NewKnight(White), piece.NewKnight(White), piece.NewKnight(White)
			bn1, bn2, bn3 = piece.NewKnight(Black), piece.NewKnight(Black), piece.NewKnight(Black)
			wk, bk = piece.NewKing(White), piece.NewKing(Black)
			b.PlacePiece(rect.Coord{X: 1, Y: 1}, wn1)
			b.PlacePiece(rect.Coord{X: 1, Y: 2}, wn2)
			b.PlacePiece(rect.Coord{X: 3, Y: 4}, wn3)
			b.PlacePiece(rect.Coord{X: 5, Y: 5}, bn1)
			b.PlacePiece(rect.Coord{X: 5, Y: 6}, bn2)
			b.PlacePiece(rect.Coord{X: 4, Y: 3}, bn3)
			b.PlacePiece(rect.Coord{X: 2, Y: 1}, wk)
			b.PlacePiece(rect.Coord{X: 5, Y: 4}, bk)
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
			b.PlacePiece(rect.Coord{X: 1, Y: 1}, bk)
			b.PlacePiece(rect.Coord{X: 2, Y: 4}, wn)
			b.PlacePiece(rect.Coord{X: 5, Y: 5}, wk)
			b.PlacePiece(rect.Coord{X: 4, Y: 4}, bn)

			attackedByWhite := b.FindAttackedCellsBy(rect.PieceFilter{
				PieceFilter: base.PieceFilter{Colours: []Colour{White}},
			})
			Expect(attackedByWhite.Len()).To(Equal(10))
			sort.Sort(attackedByWhite)
			Expect(attackedByWhite.Equals(rect.NewCoords([]base.Coord{
				rect.Coord{1, 2}, rect.Coord{3, 2}, rect.Coord{4, 3},
				rect.Coord{4, 4}, rect.Coord{5, 4}, rect.Coord{4, 5},
				rect.Coord{1, 6}, rect.Coord{3, 6}, rect.Coord{4, 6}, rect.Coord{5, 6},
			}))).To(BeTrue())

			attackedByBlack := b.FindAttackedCellsBy(rect.PieceFilter{
				PieceFilter: base.PieceFilter{Colours: []Colour{Black}},
			})
			Expect(attackedByBlack.Len()).To(Equal(9))
			sort.Sort(attackedByBlack)
			Expect(attackedByBlack.Equals(rect.NewCoords([]base.Coord{
				rect.Coord{2, 1}, rect.Coord{1, 2}, rect.Coord{2, 2},
				rect.Coord{3, 2}, rect.Coord{5, 2}, rect.Coord{2, 3},
				rect.Coord{2, 5}, rect.Coord{3, 6}, rect.Coord{5, 6},
			}))).To(BeTrue())
		})
	})

	Describe("check detection", func() {
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
})
