package rect_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
	"github.com/mtfelian/mtfchess/rect/xfen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("outcome test", func() {
	var b base.IBoard
	var err error

	Context("position 1, from chess960", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandard(xfen.XFEN(`rn2k1r1/ppp1pp1p/3p2p1/5bn1/P7/2N2B2/1PPPPP2/2BNK1RR w Gkq - 4 11`))
			Expect(err).NotTo(HaveOccurred())
		})

		It("is not checkmate, is not stalemate", func() {
			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(33))
			Expect(b.Outcome().Equals(base.NewOutcomeNotCompleted())).To(BeTrue())

			Expect(b.InCheckmate(White)).To(BeFalse())
			Expect(b.InCheckmate(Black)).To(BeFalse())
			Expect(b.InStalemate(White)).To(BeFalse())
			Expect(b.InStalemate(Black)).To(BeFalse())
		})
	})

	Context("position 2, checkmate", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandard(xfen.XFEN(`r3kb1r/3Q3p/p3P1n1/2p1p1P1/2P1bp2/7P/PB3P2/R4RK1 b kq - 2 24`))
			Expect(err).NotTo(HaveOccurred())
		})

		It("is checkmate", func() {
			Expect(b.InCheckmate(White)).To(BeFalse())
			Expect(b.InCheckmate(Black)).To(BeTrue())
			Expect(b.InStalemate(White)).To(BeFalse())
			Expect(b.InStalemate(Black)).To(BeFalse())

			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(0))
			Expect(b.Outcome().Equals(base.NewCheckmate(White))).To(BeTrue())

			Expect(b.HasMoves(Black)).To(BeFalse())
			Expect(b.HasMoves(White)).To(BeTrue())

			Expect(b.SideToMove()).To(Equal(Black))

			// check that making move will not be successful
			pieces := b.FindPieces(base.PieceFilter{Colours: []Colour{b.SideToMove()}})
			for i := range pieces {
				dsts := pieces[i].Destinations(b)
				for dsts.HasNext() {
					Expect(b.MakeMove(dsts.Next().(base.ICoord), pieces[i])).To(BeFalse())
				}
			}
		})
	})

	Context("position 3, stalemate", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandard(xfen.XFEN(`5bnr/4p1pq/4Qpkr/7p/2P4P/8/PP1PPPP1/RNB1KBNR b KQ - 0 10`))
			Expect(err).NotTo(HaveOccurred())
		})

		It("is stalemate", func() {
			Expect(b.InCheckmate(White)).To(BeFalse())
			Expect(b.InCheckmate(Black)).To(BeFalse())
			Expect(b.InStalemate(White)).To(BeFalse())
			Expect(b.InStalemate(Black)).To(BeTrue())

			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(0))
			Expect(b.Outcome().Equals(base.NewStalemate())).To(BeTrue())

			Expect(b.HasMoves(Black)).To(BeFalse())
			Expect(b.HasMoves(White)).To(BeTrue())

			Expect(b.SideToMove()).To(Equal(Black))

			// check that no destinations available
			pieces := b.FindPieces(base.PieceFilter{Colours: []Colour{b.SideToMove()}})
			for i := range pieces {
				Expect(pieces[i].Destinations(b).Len()).To(Equal(0))
			}
		})
	})
})
