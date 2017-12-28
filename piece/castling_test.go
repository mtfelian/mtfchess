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

	It("checks that both castlings are enabled", func() {
		wr1, wr2, wk := piece.NewRook(White), piece.NewRook(White), piece.NewKing(White)
		br1, br2, bk := piece.NewRook(Black), piece.NewRook(Black), piece.NewKing(Black)
		b.PlacePiece(rect.Coord{1, 1}, wr1)
		b.PlacePiece(rect.Coord{8, 1}, wr2)
		b.PlacePiece(rect.Coord{5, 1}, wk)
		b.PlacePiece(rect.Coord{1, 8}, br1)
		b.PlacePiece(rect.Coord{8, 8}, br2)
		b.PlacePiece(rect.Coord{5, 8}, bk)

		wc, bc := b.Castlings(White), b.Castlings(Black)

		Expect(wc).To(HaveLen(2))
		Expect(bc).To(HaveLen(2))

		wcA, wcH := wc[0], wc[1]
		bcA, bcH := bc[0], bc[1]

		Expect(wcA.Enabled).To(BeTrue())
		Expect(wcA.To).To(HaveLen(2))
		Expect(wcA.Piece).To(HaveLen(2))
		Expect(wcA.To).To(Equal([2]base.ICoord{rect.Coord{3, 1}, rect.Coord{4, 1}}))
		Expect(wcA.Piece[0].Name()).To(Equal("king"))
		Expect(wcA.Piece[1].Name()).To(Equal("rook"))

		Expect(wcH.Enabled).To(BeTrue())
		Expect(wcH.To).To(HaveLen(2))
		Expect(wcH.Piece).To(HaveLen(2))
		Expect(wcH.To).To(Equal([2]base.ICoord{rect.Coord{7, 1}, rect.Coord{6, 1}}))
		Expect(wcH.Piece[0].Name()).To(Equal("king"))
		Expect(wcH.Piece[1].Name()).To(Equal("rook"))

		Expect(bcA.Enabled).To(BeTrue())
		Expect(bcA.To).To(HaveLen(2))
		Expect(bcA.Piece).To(HaveLen(2))
		Expect(bcA.To).To(Equal([2]base.ICoord{rect.Coord{3, 8}, rect.Coord{4, 8}}))
		Expect(bcA.Piece[0].Name()).To(Equal("king"))
		Expect(bcA.Piece[1].Name()).To(Equal("rook"))

		Expect(bcH.Enabled).To(BeTrue())
		Expect(bcH.To).To(HaveLen(2))
		Expect(bcH.Piece).To(HaveLen(2))
		Expect(bcH.To).To(Equal([2]base.ICoord{rect.Coord{7, 8}, rect.Coord{6, 8}}))
		Expect(bcH.Piece[0].Name()).To(Equal("king"))
		Expect(bcH.Piece[1].Name()).To(Equal("rook"))
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
