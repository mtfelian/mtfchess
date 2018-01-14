package rect_test

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/rect"
	"github.com/mtfelian/mtfchess/xfen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checkmate test", func() {
	var b base.IBoard
	var err error

	Context("not", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandardXFEN(`rn2k1r1/ppp1pp1p/3p2p1/5bn1/P7/2N2B2/1PPPPP2/2BNK1RR w Gkq - 4 11`)
			Expect(err).NotTo(HaveOccurred())
		})

		It("is not checkmate", func() {
			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(33))
			Expect(b.InCheckMate(White)).To(BeFalse())
			Expect(b.InCheckMate(Black)).To(BeFalse())
		})
	})

	Context("checkmate", func() {
		BeforeEach(func() {
			b, err = xfen.NewFromStandardXFEN(`r3kb1r/3Q3p/p3P1n1/2p1p1P1/2P1bp2/7P/PB3P2/R4RK1 b kq - 2 24`)
			Expect(err).NotTo(HaveOccurred())
		})

		It("is checkmate", func() {
			Expect(b.InCheckMate(White)).To(BeFalse())
			Expect(b.InCheckMate(Black)).To(BeTrue())

			Expect(b.LegalMoves(rect.NewLongAlgebraicNotation())).To(HaveLen(0))

			Expect(b.HasMoves(Black)).To(BeFalse())
			Expect(b.HasMoves(White)).To(BeTrue())

			// check that making move will not be successful
			pawns := b.FindPieces(base.PieceFilter{Names: []string{base.PawnName}})
			for i := range pawns {
				dsts := pawns[i].Destinations(b)
				for dsts.HasNext() {
					dst := dsts.Next().(base.ICoord)
					Expect(b.MakeMove(dst, pawns[i])).To(BeFalse())
				}
			}
		})
	})
})
