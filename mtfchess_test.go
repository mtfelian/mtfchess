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

	BeforeEach(func() { b = NewEmptyStdBoard(w, h) })

	It("checks board width and height", func() {
		Expect(b.Width()).To(Equal(w))
		Expect(b.Height()).To(Equal(h))
	})

	Describe("knight", func() {
		It("generates moves", func() {
			wn1 := NewKnightPiece(White)
			wn2 := NewKnightPiece(White)
			bn := NewKnightPiece(Black)
			b.PlacePiece(2, 1, wn1)
			b.PlacePiece(3, 3, wn2)
			b.PlacePiece(4, 2, bn)

			d := wn1.Destinations(b)
			Expect(d).To(HaveLen(2))
			Expect(d).To(Equal(Pairs{{1, 3}, {4, 2}}))
		})

		It("makes legal moves", func() {
			var wn, bn Piece
			var boardCopy Board
			testReset := func() {
				wn, bn = NewKnightPiece(White), NewKnightPiece(Black)
				b.PlacePiece(2, 1, wn)
				b.PlacePiece(4, 2, bn)
				if boardCopy != nil {
					b.Set(boardCopy)
				}
			}
			testReset()
			boardCopy = b.Copy()
			destinations := wn.Destinations(b)

			for i, d := range destinations {
				x, y := wn.X(), wn.Y()
				Expect(b.MakeMove(d.X, d.Y, wn)).To(BeTrue(), "failed at destination %d", i)
				// check source cell to be empty
				Expect(b.Piece(x, y)).To(Equal(NewEmpty(x, y)))
				// check destination cell to contain new piece
				Expect(b.Piece(d.X, d.Y)).To(Equal(wn))
				if bn.X() != d.X || bn.Y() != d.Y { // if not capture
					// then there should be another piece
					Expect(b.Piece(bn.X(), bn.Y())).To(Equal(bn))
				}

				testReset()
			}
		})

		It("makes illegal moves", func() {
			var wn, bn Piece
			var boardCopy Board
			testReset := func() {
				wn, bn = NewKnightPiece(White), NewKnightPiece(Black)
				b.PlacePiece(2, 1, wn)
				b.PlacePiece(4, 2, bn)
				if boardCopy != nil {
					b.Set(boardCopy)
				}
			}
			testReset()
			boardCopy = b.Copy()
			offsets := Pairs{{3, 1}, {-1, 3}}

			for i, o := range offsets {
				x, y := wn.X(), wn.Y()
				x1, y1 := wn.X()+o.X, wn.Y()+o.Y
				Expect(b.MakeMove(x1, y1, wn)).To(BeFalse(), "failed at offset %d", i)
				// check source cell to contain unmoved piece
				Expect(b.Piece(x, y)).To(Equal(wn))
				// check destination cell to be empty
				Expect(b.Piece(x1, y1)).To(Equal(NewEmpty(x1, y1)))
				// check another cell to contain another piece
				Expect(b.Piece(bn.X(), bn.Y())).To(Equal(bn))

				testReset()
			}
		})
	})

	Describe("king", func() {
		It("generates moves", func() {
			wk := NewKingPiece(White)
			wn := NewKnightPiece(White)
			bn := NewKnightPiece(Black)
			b.PlacePiece(2, 2, wk)
			b.PlacePiece(2, 3, wn)
			b.PlacePiece(1, 1, bn)
			d := wk.Destinations(b)
			Expect(d).To(HaveLen(6))
			Expect(d).To(Equal(Pairs{{1, 1}, {1, 2}, {1, 3}, {2, 1}, {3, 1}, {3, 3}}))
		})
	})

	Describe("find pieces", func() {
		var wn1, wn2, wn3, bn1, bn2, bn3, wk, bk Piece
		BeforeEach(func() {
			wn1, wn2, wn3 = NewKnightPiece(White), NewKnightPiece(White), NewKnightPiece(White)
			bn1, bn2, bn3 = NewKnightPiece(Black), NewKnightPiece(Black), NewKnightPiece(Black)
			wk, bk = NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(1, 1, wn1)
			b.PlacePiece(1, 2, wn2)
			b.PlacePiece(3, 4, wn3)
			b.PlacePiece(5, 5, bn1)
			b.PlacePiece(5, 6, bn2)
			b.PlacePiece(4, 3, bn3)
			b.PlacePiece(2, 1, wk)
			b.PlacePiece(5, 4, bk)
		})
		It("normally", func() {
			filter := PieceFilter{ // find all white knights
				Colours: []Colour{White},
				Names:   []string{NewKnightPiece(Transparent).Name()},
			}
			coords := b.FindPieces(filter)
			Expect(coords).To(HaveLen(3))
			Expect(coords).To(Equal(Pieces{wn3, wn2, wn1}))
		})

		It("is with piece / board condition", func() {
			notOnEdge := func(p Piece) bool {
				return p.X() > 1 && p.Y() > 1 && p.X() < b.Width() && p.Y() < b.Height()
			}
			filter := PieceFilter{ // find all knights
				Names:     []string{NewKnightPiece(Transparent).Name()},
				Condition: notOnEdge,
			}

			coords := b.FindPieces(filter)
			Expect(coords).To(HaveLen(2))
			Expect(coords).To(Equal(Pieces{wn3, bn3}))
		})
	})

	Describe("find attacked cells", func() {
		It("works", func() {
			wn, bn := NewKnightPiece(White), NewKnightPiece(Black)
			wk, bk := NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(1, 1, bk)
			b.PlacePiece(2, 4, wn)
			b.PlacePiece(5, 5, wk)
			b.PlacePiece(4, 4, bn)

			attackedByWhite := b.FindAttackedCellsBy(PieceFilter{Colours: []Colour{White}})
			Expect(attackedByWhite).To(HaveLen(10))
			sort.Sort(attackedByWhite)
			Expect(attackedByWhite).To(Equal(Pairs{{1, 2}, {3, 2}, {4, 3}, {4, 4}, {5, 4},
				{4, 5}, {1, 6}, {3, 6}, {4, 6}, {5, 6}}))

			attackedByBlack := b.FindAttackedCellsBy(PieceFilter{Colours: []Colour{Black}})
			Expect(attackedByBlack).To(HaveLen(9))
			sort.Sort(attackedByBlack)
			Expect(attackedByBlack).To(Equal(Pairs{{2, 1}, {1, 2}, {2, 2}, {3, 2}, {5, 2},
				{2, 3}, {2, 5}, {3, 6}, {5, 6}}))
		})
	})

	Describe("check detection", func() {
		It("is white in check", func() {
			wn, bn := NewKnightPiece(White), NewKnightPiece(Black)
			wk, bk := NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(1, 1, wk)
			b.PlacePiece(3, 2, bn)
			b.PlacePiece(5, 4, bk)
			b.PlacePiece(4, 4, wn)

			Expect(b.InCheck(White)).To(BeTrue())
			Expect(b.InCheck(Black)).To(BeFalse())
		})

		It("is black in check", func() {
			wn, bn := NewKnightPiece(White), NewKnightPiece(Black)
			wk, bk := NewKingPiece(White), NewKingPiece(Black)
			b.PlacePiece(1, 1, bk)
			b.PlacePiece(3, 2, wn)
			b.PlacePiece(5, 4, wk)
			b.PlacePiece(4, 4, bn)

			Expect(b.InCheck(White)).To(BeFalse())
			Expect(b.InCheck(Black)).To(BeTrue())
		})
	})
})
