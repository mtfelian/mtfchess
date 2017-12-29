package rect_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Castling test", func() {
	var b base.IBoard
	resetBoard := func() { b = rect.NewStandardChessBoard() }
	BeforeEach(func() { resetBoard() })

	checkCommonCastlingProperties := func(c base.Castling) {
		Expect(c.Enabled).To(BeTrue())
		Expect(c.To).To(HaveLen(2))
		Expect(c.Piece).To(HaveLen(2))
		Expect(c.Piece[0].Name()).To(Equal("king"))
		Expect(c.Piece[1].Name()).To(Equal("rook"))
	}

	checkMakeCastling := func(c base.Castling) {
		boardCopy := b.Copy()
		Expect(boardCopy.MakeCastling(c.Copy(boardCopy))).To(BeTrue())
	}

	checkWhiteCastlingASideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{3, 1}, rect.Coord{4, 1}}))
		checkMakeCastling(c)
	}

	checkWhiteCastlingHSideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{7, 1}, rect.Coord{6, 1}}))
		checkMakeCastling(c)
	}

	checkBlackCastlingASideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{3, 8}, rect.Coord{4, 8}}))
		checkMakeCastling(c)
	}

	checkBlackCastlingHSideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{7, 8}, rect.Coord{6, 8}}))
		checkMakeCastling(c)
	}

	Context("4 rooks, 2 kings", func() {
		var wr1, wr2, wk, br1, br2, bk base.IPiece
		setupPosition := func() {
			wr1, wr2, wk = piece.NewRook(White), piece.NewRook(White), piece.NewKing(White)
			br1, br2, bk = piece.NewRook(Black), piece.NewRook(Black), piece.NewKing(Black)
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
			checkWhiteCastlingHSideEnabled(wc[1])
			checkBlackCastlingASideEnabled(bc[0])
			checkBlackCastlingHSideEnabled(bc[1])
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
			checkWhiteCastlingHSideEnabled(wc[0])
			checkBlackCastlingASideEnabled(bc[0])
			checkBlackCastlingHSideEnabled(bc[1])
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
			b.PlacePiece(rect.Coord{6, 1}, piece.NewKnight(Black))
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(1))
			Expect(bc).To(HaveLen(2))
			checkWhiteCastlingASideEnabled(wc[0])
			checkBlackCastlingASideEnabled(bc[0])
			checkBlackCastlingHSideEnabled(bc[1])
		})

		It("checks that only one castling is enabled due to opponent's piece at king's dst", func() {
			setupPosition()
			b.PlacePiece(rect.Coord{7, 1}, piece.NewKnight(Black))
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(1))
			Expect(bc).To(HaveLen(2))
			checkWhiteCastlingASideEnabled(wc[0])
			checkBlackCastlingASideEnabled(bc[0])
			checkBlackCastlingHSideEnabled(bc[1])
		})

		It("checks that only one castling is enabled due to opponent's piece at king's path", func() {
			setupPosition()
			b.PlacePiece(rect.Coord{4, 8}, piece.NewKnight(Black))
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(2))
			Expect(bc).To(HaveLen(1))
			checkWhiteCastlingASideEnabled(wc[0])
			checkWhiteCastlingHSideEnabled(wc[1])
			checkBlackCastlingHSideEnabled(bc[0])
		})

		It("checks that only one castling is enabled due to opponent's piece at king's dst", func() {
			setupPosition()
			b.PlacePiece(rect.Coord{3, 8}, piece.NewKnight(Black))
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(2))
			Expect(bc).To(HaveLen(1))
			checkWhiteCastlingASideEnabled(wc[0])
			checkWhiteCastlingHSideEnabled(wc[1])
			checkBlackCastlingHSideEnabled(bc[0])
		})

		It("checks that no castlings if in check", func() {
			setupPosition()
			b.PlacePiece(rect.Coord{4, 2}, piece.NewBishop(Black))
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(0))
			Expect(bc).To(HaveLen(2))
			checkBlackCastlingASideEnabled(bc[0])
			checkBlackCastlingHSideEnabled(bc[1])
		})

		It("checks that no castlings if king moved", func() {
			setupPosition()
			b.King(Black).MarkMoved()
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(2))
			Expect(bc).To(HaveLen(0))
			checkWhiteCastlingASideEnabled(wc[0])
			checkWhiteCastlingHSideEnabled(wc[1])
		})

		It("checks that no castlings if king not in standard position", func() {
			setupPosition()
			b.MakeMove(rect.Coord{4, 8}, b.King(Black))
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(2))
			Expect(bc).To(HaveLen(0))
			checkWhiteCastlingASideEnabled(wc[0])
			checkWhiteCastlingHSideEnabled(wc[1])
		})

		It("checks that only one castling is enabled if one of rook moved", func() {
			setupPosition()
			br1.MarkMoved()
			wr2.MarkMoved()
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(1))
			Expect(bc).To(HaveLen(1))
			checkWhiteCastlingASideEnabled(wc[0])
			checkBlackCastlingHSideEnabled(bc[0])
		})

		It("checks that no castlings if both rooks moved", func() {
			setupPosition()
			br1.MarkMoved()
			br2.MarkMoved()
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(2))
			Expect(bc).To(HaveLen(0))
			checkWhiteCastlingASideEnabled(wc[0])
			checkWhiteCastlingHSideEnabled(wc[1])
		})

		It("checks that no castlings if both rooks not in standard position", func() {
			setupPosition()
			Expect(b.MakeMove(rect.Coord{1, 7}, br1)).To(BeTrue())
			Expect(b.MakeMove(rect.Coord{8, 7}, br2)).To(BeTrue())
			wc, bc := b.Castlings(White), b.Castlings(Black)
			Expect(wc).To(HaveLen(2))
			Expect(bc).To(HaveLen(0))
			checkWhiteCastlingASideEnabled(wc[0])
			checkWhiteCastlingHSideEnabled(wc[1])
		})
	})

	Context("2 rooks, 2 kings", func() {
		var wr, wk, br, bk base.IPiece
		setupPosition := func() {
			wr, wk = piece.NewRook(White), piece.NewKing(White)
			br, bk = piece.NewRook(Black), piece.NewKing(Black)
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
			checkBlackCastlingHSideEnabled(bc[0])
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
