package mtfchess_test

import (
	"fmt"
	. "github.com/mtfelian/mtfchess"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board test", func() {
	w, h := 6, 8
	var b Board

	BeforeEach(func() { b = NewEmptyBoard(w, h) })

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

			o := wn1.Offsets(b)
			Expect(o).To(HaveLen(2))
			Expect(o).To(Equal(Offsets{{-1, 2}, {2, 1}}))
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
			offsets, c := wn.Offsets(b), wn.Coords()

			for i, o := range offsets {
				x1, y1 := c.X+o.X, c.Y+o.Y
				Expect(b.MakeMove(x1, y1, wn)).To(BeTrue(), "failed at offset %d", i)
				// check source square to be empty
				Expect(b.Piece(c.X, c.Y)).To(Equal(NewEmpty(c.X, c.Y)))
				// check destination square to contain new piece
				Expect(b.Piece(x1, y1)).To(Equal(wn))
				bnX, bnY := bn.Coords().X, bn.Coords().Y
				if bnX != x1 || bnY != y1 { // if not capture
					// then there should be another piece
					Expect(b.Piece(bnX, bnY)).To(Equal(bn))
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
			offsets, c := Offsets{{3, 1}, {-1, 3}}, wn.Coords()

			for i, o := range offsets {
				x1, y1 := c.X+o.X, c.Y+o.Y
				Expect(b.MakeMove(x1, y1, wn)).To(BeFalse(), "failed at offset %d", i)
				// check source square to contain unmoved piece
				Expect(b.Piece(c.X, c.Y)).To(Equal(wn))
				// check destination square to be empty
				Expect(b.Piece(x1, y1)).To(Equal(NewEmpty(x1, y1)))
				bnX, bnY := bn.Coords().X, bn.Coords().Y
				// chck another square to contain another piece
				Expect(b.Piece(bnX, bnY)).To(Equal(bn))

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
			fmt.Println(b)
			o := wk.Offsets(b)
			Expect(o).To(HaveLen(7))
			Expect(o).To(Equal(Offsets{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}}))
		})
	})

})
