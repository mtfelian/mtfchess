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

var _ = Describe("Pawn test with 0-modifier", func() {
	var b base.IBoard
	resetBoard := func() {
		b = rect.NewTestEmptyBoard()
		s := b.(*rect.Board).Settings()
		s.PawnLongModifier = 0
		b.SetSettings(s)
	}

	BeforeEach(func() { resetBoard() })

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

	It("makes legal moves", func() {
		var wp, bn base.IPiece
		testReset := func() {
			resetBoard()
			wp, bn = piece.NewPawn(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 2}, wp)
			b.PlacePiece(rect.Coord{1, 3}, bn)
		}
		testReset()
		destinations := wp.Destinations(b)

		wpCoord, bnCoord := wp.Coord().Copy(), bn.Coord().Copy()
		Expect(destinations.Len()).To(Equal(2))
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wp)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wpCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wp))
			if !bnCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(bnCoord)).To(Equal(bn))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wp.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(bn.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wp, bn base.IPiece
		testReset := func() {
			resetBoard()
			wp, bn = piece.NewPawn(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 2}, wp)
			b.PlacePiece(rect.Coord{1, 3}, bn)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{
			rect.Coord{3, 3}, rect.Coord{2, 4}, wp.Coord(), rect.Coord{2, 1},
		})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wp.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wp)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wp))

			// check that destination cell was not changed
			if dCellCopy.Piece() == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(dCellCopy.Piece()))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bn.Coord())).To(Equal(bn))

			testReset()
		}
	})
})

var _ = Describe("Pawn test with non-0-modifier", func() {
	var b base.IBoard
	resetBoard := func() {
		b = rect.NewTestEmptyBoard()
		s := b.(*rect.Board).Settings()
		s.PawnLongModifier = 1
		b.SetSettings(s)
	}

	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wp, bn := piece.NewPawn(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bn)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}})))

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(3))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{1, 3}, rect.Coord{2, 3}, rect.Coord{2, 4},
		}))).To(BeTrue())
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

	It("makes legal moves", func() {
		var wp, bn base.IPiece
		testReset := func() {
			resetBoard()
			wp, bn = piece.NewPawn(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 2}, wp)
			b.PlacePiece(rect.Coord{1, 3}, bn)
		}
		testReset()
		destinations := wp.Destinations(b)

		wpCoord, bnCoord := wp.Coord().Copy(), bn.Coord().Copy()
		Expect(destinations.Len()).To(Equal(3))
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wp)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wpCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wp))
			if !bnCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(bnCoord)).To(Equal(bn))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wp.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(bn.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wp, bn base.IPiece
		testReset := func() {
			resetBoard()
			wp, bn = piece.NewPawn(White), piece.NewKnight(Black)
			b.PlacePiece(rect.Coord{2, 2}, wp)
			b.PlacePiece(rect.Coord{1, 3}, bn)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{
			rect.Coord{3, 3}, wp.Coord(), rect.Coord{2, 1},
		})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wp.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wp)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wp))

			// check that destination cell was not changed
			if dCellCopy.Piece() == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(dCellCopy.Piece()))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bn.Coord())).To(Equal(bn))

			testReset()
		}
	})
})
