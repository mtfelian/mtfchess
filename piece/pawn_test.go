package piece_test

import (
	"sort"

	"fmt"
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
		s.PawnLongMoveFunc = rect.NoPawnLongMoveFunc
		b.SetSettings(s)
	}

	BeforeEach(func() { resetBoard() })

	It("generates moves 1", func() {
		wp, bn := piece.NewPawn(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bn)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(2))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{2, 3}}))).To(BeTrue())
	})

	It("generates moves 2", func() {
		wp, bp1 := piece.NewPawn(White), piece.NewPawn(Black)
		bp2, bp3 := piece.NewPawn(Black), piece.NewPawn(Black)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bp1)
		b.PlacePiece(rect.Coord{2, 3}, bp2)
		b.PlacePiece(rect.Coord{3, 3}, bp3)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())

		d := wp.Destinations(b)
		fmt.Println(d)
		Expect(d.Len()).To(Equal(2))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())
	})

	It("generates moves 3", func() {
		wp, bp1 := piece.NewPawn(White), piece.NewPawn(Black)
		wn, bp3 := piece.NewKnight(White), piece.NewPawn(Black)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bp1)
		b.PlacePiece(rect.Coord{2, 3}, wn)
		b.PlacePiece(rect.Coord{3, 3}, bp3)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())

		d := wp.Destinations(b)
		fmt.Println(d)
		Expect(d.Len()).To(Equal(2))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())
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
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())

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
		s.PawnLongMoveFunc = rect.StandardLongMoveFunc
		b.SetSettings(s)
	}

	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wp, bn := piece.NewPawn(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bn)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())

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
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{3, 3}}))).To(BeTrue())

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

var _ = Describe("Pawn promotion test", func() {
	var b base.IBoard
	resetBoard := func() {
		b = rect.NewTestEmptyBoard()
		s := b.(*rect.Board).Settings()
		s.PawnLongMoveFunc = rect.StandardLongMoveFunc
		b.SetSettings(s)
	}

	BeforeEach(func() {
		resetBoard()
	})

	It("pawn promotes by non-capturing move", func() {
		wp, bk := piece.NewPawn(White), piece.NewKing(Black)
		b.PlacePiece(rect.Coord{2, 5}, wp)
		b.PlacePiece(rect.Coord{4, 6}, bk)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 6}, rect.Coord{3, 6}}))).To(BeTrue())

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(1))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{2, 6}}))).To(BeTrue())

		wp.SetPromote(piece.NewRook(wp.Colour()))
		Expect(b.MakeMove(rect.Coord{2, 6}, wp)).To(BeTrue())

		Expect(b.InCheck(Black)).To(BeTrue())
	})

	It("pawn promotes by capturing move", func() {
		wp, bk, bn := piece.NewPawn(White), piece.NewKing(Black), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 5}, wp)
		b.PlacePiece(rect.Coord{3, 6}, bn)
		b.PlacePiece(rect.Coord{4, 6}, bk)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 6}, rect.Coord{3, 6}}))).To(BeTrue())

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(2))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{2, 6}, rect.Coord{3, 6}}))).To(BeTrue())

		wp.SetPromote(piece.NewRook(wp.Colour()))
		Expect(b.MakeMove(rect.Coord{3, 6}, wp)).To(BeTrue())

		Expect(b.InCheck(Black)).To(BeTrue())
	})

	It("pawn promotes by capture releasing check", func() {
		wp, bk := piece.NewPawn(White), piece.NewKing(Black)
		bq, wk, br := piece.NewQueen(Black), piece.NewKing(White), piece.NewRook(Black)
		b.PlacePiece(rect.Coord{1, 6}, bq)
		b.PlacePiece(rect.Coord{3, 5}, wp)
		b.PlacePiece(rect.Coord{3, 3}, bk)
		b.PlacePiece(rect.Coord{4, 6}, br)
		b.PlacePiece(rect.Coord{5, 6}, wk)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{2, 6}, rect.Coord{4, 6}}))).To(BeTrue())

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(1))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{4, 6}}))).To(BeTrue())

		wp.SetPromote(piece.NewRook(wp.Colour()))
		Expect(b.MakeMove(rect.Coord{4, 6}, wp)).To(BeTrue())
		Expect(b.InCheck(White)).To(BeFalse())
	})

	It("pawn promotes releasing check", func() {
		wp, bk := piece.NewPawn(White), piece.NewKing(Black)
		bq, wk, br := piece.NewQueen(Black), piece.NewKing(White), piece.NewRook(Black)
		b.PlacePiece(rect.Coord{1, 6}, bq)
		b.PlacePiece(rect.Coord{3, 5}, wp)
		b.PlacePiece(rect.Coord{3, 3}, bk)
		b.PlacePiece(rect.Coord{4, 6}, br)
		b.PlacePiece(rect.Coord{4, 6}, wk)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{2, 6}, rect.Coord{4, 6}}))).To(BeTrue())

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(1))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{3, 6}}))).To(BeTrue())

		wp.SetPromote(piece.NewRook(wp.Colour()))
		Expect(b.MakeMove(rect.Coord{3, 6}, wp)).To(BeTrue())

		Expect(b.InCheck(White)).To(BeFalse())
	})

	It("pawn tries to make an invalid promotion", func() {
		s := b.(*rect.Board).Settings()
		s.AllowedPromotions = []string{"knight", "bishop", "queen"} // exclude rook from promotion list
		b.SetSettings(s)

		wp, bk := piece.NewPawn(White), piece.NewKing(Black)
		bq, wk, br := piece.NewQueen(Black), piece.NewKing(White), piece.NewRook(Black)
		b.PlacePiece(rect.Coord{1, 6}, bq)
		b.PlacePiece(rect.Coord{3, 5}, wp)
		b.PlacePiece(rect.Coord{3, 3}, bk)
		b.PlacePiece(rect.Coord{4, 6}, br)
		b.PlacePiece(rect.Coord{5, 6}, wk)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{2, 6}, rect.Coord{4, 6}}))).To(BeTrue())

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(1))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{4, 6}}))).To(BeTrue())

		wpCoords, bkCoords, bqCoords := wp.Coord().Copy(), bk.Coord().Copy(), bq.Coord().Copy()
		wkCoords, brCoords := wk.Coord().Copy(), br.Coord().Copy()

		boardCopy := b.Copy()
		wp.SetPromote(piece.NewRook(wp.Colour()))
		Expect(b.MakeMove(rect.Coord{4, 6}, wp)).To(BeFalse())
		Expect(b.InCheck(White)).To(BeTrue())

		// check that board did not changed
		Expect(b.Piece(wpCoords)).To(Equal(wp))
		Expect(b.Piece(bkCoords)).To(Equal(bk))
		Expect(b.Piece(bqCoords)).To(Equal(bq))
		Expect(b.Piece(wkCoords)).To(Equal(wk))
		Expect(b.Piece(brCoords)).To(Equal(br))
		Expect(b.Piece(rect.Coord{3, 6})).To(BeNil())

		Expect(b.Equals(boardCopy)).To(BeTrue())
	})

	It("pawn tries to make a promotion from invalid cell", func() {
		wp, bq := piece.NewPawn(White), piece.NewQueen(Black)
		b.PlacePiece(rect.Coord{1, 6}, bq)
		b.PlacePiece(rect.Coord{3, 4}, wp)

		a := wp.Attacks(b)
		Expect(a.Len()).To(Equal(2))
		sort.Sort(a)
		Expect(a.Equals(rect.NewCoords([]base.ICoord{rect.Coord{2, 5}, rect.Coord{4, 5}}))).To(BeTrue())

		d := wp.Destinations(b)
		Expect(d.Len()).To(Equal(1))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{3, 5}}))).To(BeTrue())

		wpCoords, bqCoords := wp.Coord().Copy(), bq.Coord().Copy()

		boardCopy := b.Copy()
		wp.SetPromote(piece.NewRook(wp.Colour()))
		Expect(b.MakeMove(rect.Coord{3, 5}, wp)).To(BeFalse())

		// check that board did not changed
		Expect(b.Piece(wpCoords)).To(Equal(wp))
		Expect(b.Piece(bqCoords)).To(Equal(bq))

		Expect(b.Equals(boardCopy)).To(BeTrue())
	})
})
