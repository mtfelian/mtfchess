package piece_test

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

	checkWhiteCastlingASideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{3, 1}, rect.Coord{4, 1}}))
	}

	checkWhiteCastlingHSideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{7, 1}, rect.Coord{6, 1}}))
	}

	checkBlackCastlingASideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{3, 8}, rect.Coord{4, 8}}))
	}

	checkBlackCastlingHSideEnabled := func(c base.Castling) {
		checkCommonCastlingProperties(c)
		Expect(c.To).To(Equal([2]base.ICoord{rect.Coord{7, 8}, rect.Coord{6, 8}}))
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
	})

	/*
		todo test cases:
		only one castling is enabled due to second rook moved
		only one castling is enabled due to second rook not in standard position
		only one castling is enabled due to king's dst attacked
		only one castling is enabled due to king's path attacked
		only one castling is enabled due to opponent's piece at king's path
		only one castling is enabled due to opponent's piece at king's dst
		only one castling is enabled due to own piece at king's path
		only one castling is enabled due to own piece at king's dst
		two castlings enabled, but rook source cell attacked
		no castlings if in check
		no castlings if king moved
		no castlings if king not in standard position
		no castlings if both rooks moved
		no castlings if both rooks not in standard position
	*/
})
