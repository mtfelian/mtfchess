package rect_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Castling test", func() {
	var b base.IBoard
	var resetBoard func()
	JustBeforeEach(func() { resetBoard() })

	checkCommonCastlingProperties := func(c base.Castling) {
		Expect(c.Enabled).To(BeTrue())
		Expect(c.To).To(HaveLen(2))
		Expect(c.Piece).To(HaveLen(2))
		Expect(c.Piece[0].Name()).To(Equal(base.KingName))
		Expect(c.Piece[1].Name()).To(Equal(base.RookName))
	}

	checkMakeCastling := func(c base.Castling) {
		boardCopy := b.Copy()
		Expect(boardCopy.MakeCastling(c.Copy(boardCopy))).To(BeTrue())
		Expect(boardCopy.SideToMove()).To(Equal(b.SideToMove().Invert()), "side to move did not changed on castling")
	}

	checkWhiteCastlingASideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{3, 1}, rect.Coord{4, 1}}))
		checkMakeCastling(c)
	}

	checkWhiteCastlingZSideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{7, 1}, rect.Coord{6, 1}}))
		checkMakeCastling(c)
	}

	checkBlackCastlingASideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{3, 8}, rect.Coord{4, 8}}))
		checkMakeCastling(c)
	}

	checkBlackCastlingZSideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{7, 8}, rect.Coord{6, 8}}))
		checkMakeCastling(c)
	}

	Context("for standard position", func() {
		BeforeEach(func() {
			resetBoard = func() {
				b = rect.NewEmptyStandardChessBoard()
				b.Settings().MoveOrder = false
				// set rook initial coords to enable castling
				b.SetRookInitialCoords(White, 0, rect.Coord{1, 1})
				b.SetRookInitialCoords(White, 1, rect.Coord{8, 1})
				b.SetRookInitialCoords(Black, 0, rect.Coord{1, 8})
				b.SetRookInitialCoords(Black, 1, rect.Coord{8, 8})
			}
		})

		Context("4 rooks, 2 kings", func() {
			var wr1, wr2, wk, br1, br2, bk base.IPiece
			setupPosition := func() {
				wr1, wr2, wk = rect.NewRook(White), rect.NewRook(White), rect.NewKing(White)
				br1, br2, bk = rect.NewRook(Black), rect.NewRook(Black), rect.NewKing(Black)
				b.PlacePiece(rect.Coord{1, 1}, wr1)
				b.PlacePiece(rect.Coord{8, 1}, wr2)
				b.PlacePiece(rect.Coord{5, 1}, wk)
				b.PlacePiece(rect.Coord{1, 8}, br1)
				b.PlacePiece(rect.Coord{8, 8}, br2)
				b.PlacePiece(rect.Coord{5, 8}, bk)
			}

			It("checks that both castlings are enabled", func() {
				setupPosition()
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(2))
				Expect(bc).To(HaveLen(2))
				checkWhiteCastlingASideEnabled(wc[0])
				checkWhiteCastlingZSideEnabled(wc[1])
				checkBlackCastlingASideEnabled(bc[0])
				checkBlackCastlingZSideEnabled(bc[1])
			})

			It("checks that only one castling is enabled due to second rook moved", func() {
				setupPosition()
				Expect(b.MakeMove(rect.Coord{8, 8}, wr2)).To(BeTrue())
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(0))
				checkWhiteCastlingASideEnabled(wc[0])
			})

			It("checks that only one castling is enabled due to second rook not in standard position", func() {
				setupPosition()
				Expect(b.MakeMove(rect.Coord{1, 2}, wr1)).To(BeTrue())
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(2))
				checkWhiteCastlingZSideEnabled(wc[0])
				checkBlackCastlingASideEnabled(bc[0])
				checkBlackCastlingZSideEnabled(bc[1])
			})

			It("checks that only one castling is enabled due to king's dst attacked", func() {
				setupPosition()
				Expect(b.MakeMove(rect.Coord{7, 8}, br2)).To(BeTrue())
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(1))
				checkWhiteCastlingASideEnabled(wc[0])
				checkBlackCastlingASideEnabled(bc[0])
			})

			It("checks that only one castling is enabled due to king's path attacked", func() {
				setupPosition()
				Expect(b.MakeMove(rect.Coord{6, 8}, br2)).To(BeTrue())
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(1))
				checkWhiteCastlingASideEnabled(wc[0])
				checkBlackCastlingASideEnabled(bc[0])
			})

			It("checks that only one castling is enabled due to opponent's piece at king's path", func() {
				setupPosition()
				b.PlacePiece(rect.Coord{6, 1}, rect.NewKnight(Black))
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(2))
				checkWhiteCastlingASideEnabled(wc[0])
				checkBlackCastlingASideEnabled(bc[0])
				checkBlackCastlingZSideEnabled(bc[1])
			})

			It("checks that only one castling is enabled due to opponent's piece at king's dst", func() {
				setupPosition()
				b.PlacePiece(rect.Coord{7, 1}, rect.NewKnight(Black))
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(2))
				checkWhiteCastlingASideEnabled(wc[0])
				checkBlackCastlingASideEnabled(bc[0])
				checkBlackCastlingZSideEnabled(bc[1])
			})

			It("checks that only one castling is enabled due to opponent's piece at king's path", func() {
				setupPosition()
				b.PlacePiece(rect.Coord{4, 8}, rect.NewKnight(Black))
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(2))
				Expect(bc).To(HaveLen(1))
				checkWhiteCastlingASideEnabled(wc[0])
				checkWhiteCastlingZSideEnabled(wc[1])
				checkBlackCastlingZSideEnabled(bc[0])
			})

			It("checks that only one castling is enabled due to opponent's piece at king's dst", func() {
				setupPosition()
				b.PlacePiece(rect.Coord{3, 8}, rect.NewKnight(Black))
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(2))
				Expect(bc).To(HaveLen(1))
				checkWhiteCastlingASideEnabled(wc[0])
				checkWhiteCastlingZSideEnabled(wc[1])
				checkBlackCastlingZSideEnabled(bc[0])
			})

			It("checks that no castlings if in check", func() {
				setupPosition()
				b.PlacePiece(rect.Coord{4, 2}, rect.NewBishop(Black))
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(0))
				Expect(bc).To(HaveLen(2))
				checkBlackCastlingASideEnabled(bc[0])
				checkBlackCastlingZSideEnabled(bc[1])
			})

			It("checks that no castlings if king moved", func() {
				setupPosition()
				b.King(Black).MarkMoved()
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(2))
				Expect(bc).To(HaveLen(0))
				checkWhiteCastlingASideEnabled(wc[0])
				checkWhiteCastlingZSideEnabled(wc[1])
			})

			It("checks that no castlings if king not in standard position", func() {
				setupPosition()
				b.MakeMove(rect.Coord{4, 8}, b.King(Black))
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(2))
				Expect(bc).To(HaveLen(0))
				checkWhiteCastlingASideEnabled(wc[0])
				checkWhiteCastlingZSideEnabled(wc[1])
			})

			It("checks that only one castling is enabled if one of rook moved", func() {
				setupPosition()
				br1.MarkMoved()
				wr2.MarkMoved()
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(1))
				checkWhiteCastlingASideEnabled(wc[0])
				checkBlackCastlingZSideEnabled(bc[0])
			})

			It("checks that no castlings if both rooks moved", func() {
				setupPosition()
				br1.MarkMoved()
				br2.MarkMoved()
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(2))
				Expect(bc).To(HaveLen(0))
				checkWhiteCastlingASideEnabled(wc[0])
				checkWhiteCastlingZSideEnabled(wc[1])
			})

			It("checks that no castlings if both rooks not in standard position", func() {
				setupPosition()
				Expect(b.MakeMove(rect.Coord{1, 7}, br1)).To(BeTrue())
				Expect(b.MakeMove(rect.Coord{8, 7}, br2)).To(BeTrue())
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(2))
				Expect(bc).To(HaveLen(0))
				checkWhiteCastlingASideEnabled(wc[0])
				checkWhiteCastlingZSideEnabled(wc[1])
			})
		})

		Context("2 rooks, 2 kings", func() {
			var wr, wk, br, bk base.IPiece
			setupPosition := func() {
				wr, wk = rect.NewRook(White), rect.NewKing(White)
				br, bk = rect.NewRook(Black), rect.NewKing(Black)
				b.PlacePiece(rect.Coord{1, 1}, wr)
				b.PlacePiece(rect.Coord{5, 1}, wk)
				b.PlacePiece(rect.Coord{8, 8}, br)
				b.PlacePiece(rect.Coord{5, 8}, bk)
			}

			It("checks that only one castling for each side", func() {
				setupPosition()
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(1))
				checkWhiteCastlingASideEnabled(wc[0])
				checkBlackCastlingZSideEnabled(bc[0])
			})
		})

		Context("empty board", func() {
			It("checks that no castlings", func() {
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(0))
				Expect(bc).To(HaveLen(0))
			})
		})
	})

	Context("for chess960 specially crafted position", func() {
		BeforeEach(func() {
			resetBoard = func() {
				b = rect.NewEmptyStandardChessBoard()
				// set rook initial coords to enable castling
				b.SetRookInitialCoords(White, 0, rect.Coord{1, 1})
				b.SetRookInitialCoords(White, 1, rect.Coord{7, 1})
				b.SetRookInitialCoords(Black, 0, rect.Coord{1, 8})
				b.SetRookInitialCoords(Black, 1, rect.Coord{7, 8})
			}
		})

		Context("4 rooks, 2 kings", func() {
			var wr1, wr2, wk, br1, br2, bk base.IPiece
			setupPosition := func() {
				wr1, wr2, wk = rect.NewRook(White), rect.NewRook(White), rect.NewKing(White)
				br1, br2, bk = rect.NewRook(Black), rect.NewRook(Black), rect.NewKing(Black)

				b.PlacePiece(rect.Coord{3, 1}, rect.NewBishop(White))
				b.PlacePiece(rect.Coord{4, 1}, rect.NewKnight(White))
				b.PlacePiece(rect.Coord{2, 2}, rect.NewPawn(White))
				b.PlacePiece(rect.Coord{3, 2}, rect.NewPawn(White))
				b.PlacePiece(rect.Coord{4, 2}, rect.NewPawn(White))
				b.PlacePiece(rect.Coord{5, 2}, rect.NewPawn(White))
				b.PlacePiece(rect.Coord{6, 2}, rect.NewPawn(White))
				b.PlacePiece(rect.Coord{3, 3}, rect.NewKnight(White))
				b.PlacePiece(rect.Coord{6, 3}, rect.NewBishop(White))
				b.PlacePiece(rect.Coord{1, 4}, rect.NewPawn(White))
				b.PlacePiece(rect.Coord{6, 5}, rect.NewBishop(Black))
				b.PlacePiece(rect.Coord{7, 5}, rect.NewKnight(Black))
				b.PlacePiece(rect.Coord{4, 6}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{7, 6}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{1, 7}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{2, 7}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{3, 7}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{5, 7}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{6, 7}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{8, 7}, rect.NewPawn(Black))
				b.PlacePiece(rect.Coord{2, 8}, rect.NewKnight(Black))

				b.PlacePiece(rect.Coord{8, 1}, wr1)
				wr1.MarkMoved()
				b.PlacePiece(rect.Coord{7, 1}, wr2)
				b.PlacePiece(rect.Coord{5, 1}, wk)
				b.PlacePiece(rect.Coord{1, 8}, br1)
				b.PlacePiece(rect.Coord{7, 8}, br2)
				b.PlacePiece(rect.Coord{5, 8}, bk)

				b.Settings().MoveOrder = false
				b.SetSideToMove(White)
			}

			It("checks that both castlings are enabled", func() {
				setupPosition()
				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(1))
				checkWhiteCastlingZSideEnabled(wc[0])
				checkBlackCastlingZSideEnabled(bc[0])
			})

			It("checks that can't do castling not in opponent's move", func() {
				setupPosition()
				b.Settings().MoveOrder = true

				wc, bc := b.Castlings(White), b.Castlings(Black)
				Expect(wc).To(HaveLen(1))
				Expect(bc).To(HaveLen(1))

				boardCopy := b.Copy()
				Expect(b.MakeCastling(bc[0])).To(BeFalse(), "did castling in opponent's move")
				Expect(boardCopy.Equals(b)).To(BeTrue(), "castling failed but board changed")

				Expect(b.MakeCastling(wc[0])).To(BeTrue(), "can't to castling in my move")
				Expect(b.SideToMove()).To(Equal(boardCopy.SideToMove().Invert()))
			})
		})
	})
})
