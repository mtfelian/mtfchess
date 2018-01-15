package rect_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
	"github.com/mtfelian/mtfchess/xfen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("stalemate test", func() {
	var b base.IBoard
	var err error

	Context("not, position 1", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandardXFEN(`rn2k1r1/ppp1pp1p/3p2p1/5bn1/P7/2N2B2/1PPPPP2/2BNK1RR w Gkq - 4 11`)
			Expect(err).NotTo(HaveOccurred())
		})

		It("is not checkmate", func() {
			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(33))
			Expect(b.InStaleMate(White)).To(BeFalse())
			Expect(b.InStaleMate(Black)).To(BeFalse())
		})
	})

	Context("not, position 2", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandardXFEN(`r3kb1r/3Q3p/p3P1n1/2p1p1P1/2P1bp2/7P/PB3P2/R4RK1 b kq - 2 24`)
			Expect(err).NotTo(HaveOccurred())
		})

		It("is checkmate", func() {
			Expect(b.InStaleMate(White)).To(BeFalse())
			Expect(b.InStaleMate(Black)).To(BeFalse())
			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(0))
		})
	})

	Context("stalemate", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandardXFEN(`5bnr/4p1pq/4Qpkr/7p/2P4P/8/PP1PPPP1/RNB1KBNR b KQ - 0 10`)
			Expect(err).NotTo(HaveOccurred())
		})

		It("is stalemate", func() {
			Expect(b.InStaleMate(White)).To(BeFalse())
			Expect(b.InStaleMate(Black)).To(BeTrue())

			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(0))

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
