package mtfchess_test

import (
	"sort"

	. "github.com/mtfelian/mtfchess"
	. "github.com/mtfelian/mtfchess/board"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board test", func() {
	w, h := 5, 6
	var b Board

	BeforeEach(func() { b = NewEmptyRectBoard(w, h) })

	It("checks board width and height", func() {
		Expect(b.Dim().(RectCoord).X).To(Equal(w))
		Expect(b.Dim().(RectCoord).Y).To(Equal(h))
	})

	Describe("knight", func() {
		It("generates moves", func() {
			wn1 := NewKnightPiece(White)
			wn2 := NewKnightPiece(White)
			bn := NewKnightPiece(Black)
			b.PlacePiece(RectCoord{X: 2, Y: 1}, wn1)
			b.PlacePiece(RectCoord{X: 3, Y: 3}, wn2)
			b.PlacePiece(RectCoord{X: 4, Y: 2}, bn)

			d := wn1.Destinations(b)
			Expect(d.Len()).To(Equal(2))
			Expect(d.Equals(NewRectCoords([]Coord{RectCoord{1, 3}, RectCoord{4, 2}}))).To(BeTrue())
		})

		It("makes legal moves", func() {
			var wn, bn Piece
			var boardCopy Board
			testReset := func() {
				wn, bn = NewKnightPiece(White), NewKnightPiece(Black)
				b.PlacePiece(RectCoord{X: 2, Y: 1}, wn)
				b.PlacePiece(RectCoord{X: 4, Y: 2}, bn)
				if boardCopy != nil {
					b.Set(boardCopy)
				}
			}
			testReset()
			boardCopy = b.Copy()
			destinations := wn.Destinations(b)

			for destinations.HasNext() {
				d := destinations.Next().(Coord)
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
			var wn, bn Piece
			var boardCopy Board
			testReset := func() {
				wn, bn = NewKnightPiece(White), NewKnightPiece(Black)
				b.PlacePiece(RectCoord{X: 2, Y: 1}, wn)
				b.PlacePiece(RectCoord{X: 4, Y: 2}, bn)
				if boardCopy != nil {
					b.Set(boardCopy)
				}
			}
			testReset()
			boardCopy = b.Copy()
			offsets := NewRectCoords([]Coord{RectCoord{3, 1}, RectCoord{-1, 3}})

			for offsets.HasNext() {
				o := offsets.Next().(Coord)
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
			wk := NewKingPiece(White)
			wn := NewKnightPiece(White)
			bn := NewKnightPiece(Black)
			b.PlacePiece(RectCoord{X: 2, Y: 2}, wk)
			b.PlacePiece(RectCoord{X: 2, Y: 3}, wn)
			b.PlacePiece(RectCoord{X: 1, Y: 1}, bn)
			d := wk.Destinations(b)
			Expect(d.Len()).To(Equal(6))
			Expect(d.Equals(NewRectCoords([]Coord{
				RectCoord{1, 1}, RectCoord{1, 2}, RectCoord{1, 3},
				RectCoord{2, 1}, RectCoord{3, 1}, RectCoord{3, 3},
			}))).To(BeTrue())
		})
	})

	Describe("find pieces", func() {
		var wn1, wn2, wn3, bn1, bn2, bn3, wk, bk Piece
		BeforeEach(func() {
			wn1, wn2, wn3 = NewKnightPiece(White), NewKnightPiece(White), NewKnightPiece(White)
			bn1, bn2, bn3 = NewKnightPiece(Black), NewKnightPiece(Black), NewKnightPiece(Black)
			wk, bk = NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(RectCoord{X: 1, Y: 1}, wn1)
			b.PlacePiece(RectCoord{X: 1, Y: 2}, wn2)
			b.PlacePiece(RectCoord{X: 3, Y: 4}, wn3)
			b.PlacePiece(RectCoord{X: 5, Y: 5}, bn1)
			b.PlacePiece(RectCoord{X: 5, Y: 6}, bn2)
			b.PlacePiece(RectCoord{X: 4, Y: 3}, bn3)
			b.PlacePiece(RectCoord{X: 2, Y: 1}, wk)
			b.PlacePiece(RectCoord{X: 5, Y: 4}, bk)
		})
		It("normally", func() {
			filter := RectPieceFilter{ // find all white knights
				BasePieceFilter: BasePieceFilter{
					Colours: []Colour{White},
					Names:   []string{NewKnightPiece(Transparent).Name()},
				},
			}
			coords := b.FindPieces(filter)
			Expect(coords).To(HaveLen(3))
			Expect(coords).To(Equal(Pieces{wn3, wn2, wn1}))
		})

		It("is with piece / board condition", func() {
			notOnEdge := func(p Piece) bool {
				x, y := p.Coord().(RectCoord).X, p.Coord().(RectCoord).Y
				w, h := b.Dim().(RectCoord).X, b.Dim().(RectCoord).Y
				return x > 1 && y > 1 && x < w && y < h
			}
			filter := RectPieceFilter{ // find all knights
				BasePieceFilter: BasePieceFilter{
					Names:     []string{NewKnightPiece(Transparent).Name()},
					Condition: notOnEdge,
				},
			}

			pieces := b.FindPieces(filter)
			Expect(pieces).To(HaveLen(2))
			Expect(pieces).To(Equal(Pieces{wn3, bn3}))
		})
	})

	Describe("find attacked cells", func() {
		It("works", func() {
			wn, bn := NewKnightPiece(White), NewKnightPiece(Black)
			wk, bk := NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(RectCoord{X: 1, Y: 1}, bk)
			b.PlacePiece(RectCoord{X: 2, Y: 4}, wn)
			b.PlacePiece(RectCoord{X: 5, Y: 5}, wk)
			b.PlacePiece(RectCoord{X: 4, Y: 4}, bn)

			attackedByWhite := b.FindAttackedCellsBy(RectPieceFilter{
				BasePieceFilter: BasePieceFilter{Colours: []Colour{White}},
			})
			Expect(attackedByWhite.Len()).To(Equal(10))
			sort.Sort(attackedByWhite)
			Expect(attackedByWhite.Equals(NewRectCoords([]Coord{
				RectCoord{1, 2}, RectCoord{3, 2}, RectCoord{4, 3},
				RectCoord{4, 4}, RectCoord{5, 4}, RectCoord{4, 5},
				RectCoord{1, 6}, RectCoord{3, 6}, RectCoord{4, 6}, RectCoord{5, 6},
			}))).To(BeTrue())

			attackedByBlack := b.FindAttackedCellsBy(RectPieceFilter{
				BasePieceFilter: BasePieceFilter{Colours: []Colour{Black}},
			})
			Expect(attackedByBlack.Len()).To(Equal(9))
			sort.Sort(attackedByBlack)
			Expect(attackedByBlack.Equals(NewRectCoords([]Coord{
				RectCoord{2, 1}, RectCoord{1, 2}, RectCoord{2, 2},
				RectCoord{3, 2}, RectCoord{5, 2}, RectCoord{2, 3},
				RectCoord{2, 5}, RectCoord{3, 6}, RectCoord{5, 6},
			}))).To(BeTrue())
		})
	})

	Describe("check detection", func() {
		It("is white in check", func() {
			wn, bn := NewKnightPiece(White), NewKnightPiece(Black)
			wk, bk := NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(RectCoord{X: 1, Y: 1}, wk)
			b.PlacePiece(RectCoord{X: 3, Y: 2}, bn)
			b.PlacePiece(RectCoord{X: 5, Y: 4}, bk)
			b.PlacePiece(RectCoord{X: 4, Y: 4}, wn)

			Expect(b.InCheck(White)).To(BeTrue())
			Expect(b.InCheck(Black)).To(BeFalse())
		})

		It("is black in check", func() {
			wn, bn := NewKnightPiece(White), NewKnightPiece(Black)
			wk, bk := NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(RectCoord{X: 1, Y: 1}, bk)
			b.PlacePiece(RectCoord{X: 3, Y: 2}, wn)
			b.PlacePiece(RectCoord{X: 5, Y: 4}, wk)
			b.PlacePiece(RectCoord{X: 4, Y: 4}, bn)

			Expect(b.InCheck(White)).To(BeFalse())
			Expect(b.InCheck(Black)).To(BeTrue())
		})
	})
})
