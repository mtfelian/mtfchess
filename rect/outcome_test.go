package rect_test

import (
	"fmt"

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
			b, err = xfen.XFEN(`rn2k1r1/ppp1pp1p/3p2p1/5bn1/P7/2N2B2/1PPPPP2/2BNK1RR w Gkq - 4 11`).RectBoard()
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
			b, err = xfen.XFEN(`r3kb1r/3Q3p/p3P1n1/2p1p1P1/2P1bp2/7P/PB3P2/R4RK1 b kq - 2 24`).RectBoard()
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
			b, err = xfen.XFEN(`5bnr/4p1pq/4Qpkr/7p/2P4P/8/PP1PPPP1/RNB1KBNR b KQ - 0 10`).RectBoard()
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

	Context("50 moves draw rule", func() {
		BeforeEach(func() {
			b, err = xfen.NewStandardChessStartingPosition().RectBoard()
			Expect(err).NotTo(HaveOccurred())
		})

		// to returns a board coord from coord given in algebraic notation
		to := func(coord string) base.ICoord {
			n := rect.NewLongAlgebraicNotation()
			Expect(n.DecodeCoord(coord)).To(Succeed())
			return n.Coord
		}

		// from returns a piece staying on a cell with coord given in algebraic notation
		from := func(coord string) base.IPiece { return b.Piece(to(coord)) }

		// makeMove helper func
		makeMove := func(srcCoord, dstCoord string) { Expect(b.MakeMove(to(dstCoord), from(srcCoord))).To(BeTrue()) }

		// makeMoves helper func
		makeMoves := func(coords [][2]string) {
			for i := range coords {
				By(fmt.Sprintf("making move %d...", b.MoveNumber()))
				makeMove(coords[i][0], coords[i][1])
			}
		}

		It("plays full game, draw by 50 moves rule", func() {
			/*
						Anatoly Karpov vs Garry Kasparov
				"When We Were Kings" (game of the day May-09-2011)
				Interpolis 15th (1991), Tilburg NED, rd 7, Oct-25
				King's Indian Defense: Orthodox Variation. Modern System (E97)  ·  1/2-1/2
						1.d4 Nf6 2.c4 g6 3.Nc3 Bg7 4.e4 d6 5.Nf3 O-O 6.Be2 e5 7.O-O
						Nc6 8.d5 Ne7 9.Nd2 a5 10.Rb1 Nd7 11.a3 f5 12.b4 Kh8 13.f3 Ng8
						14.Qc2 Ngf6 15.Nb5 axb4 16.axb4 Nh5 17.g3 Ndf6 18.c5 Bd7
						19.Rb3 Nxg3 20.hxg3 Nh5 21.f4 exf4 22.c6 bxc6 23.dxc6 Nxg3
						24.Rxg3 fxg3 25.cxd7 g2 26.Rf3 Qxd7 27.Bb2 fxe4 28.Rxf8+ Rxf8
						29.Bxg7+ Qxg7 30.Qxe4 Qf6 31.Nf3 Qf4 32.Qe7 Rf7 33.Qe6 Rf6
						34.Qe8+ Rf8 35.Qe7 Rf7 36.Qe6 Rf6 37.Qb3 g5 38.Nxc7 g4 39.Nd5
						Qc1+ 40.Qd1 Qxd1+ 41.Bxd1 Rf5 42.Ne3 Rf4 43.Ne1 Rxb4 44.Bxg4
						h5 45.Bf3 d5 46.N3xg2 h4 47.Nd3 Ra4 48.Ngf4 Kg7 49.Kg2 Kf6
						50.Bxd5 Ra5 51.Bc6 Ra6 52.Bb7 Ra3 53.Be4 Ra4 54.Bd5 Ra5 55.Bc6
						Ra6 56.Bf3 Kg5 57.Bb7 Ra1 58.Bc8 Ra4 59.Kf3 Rc4 60.Bd7 Kf6
						61.Kg4 Rd4 62.Bc6 Rd8 63.Kxh4 Rg8 64.Be4 Rg1 65.Nh5+ Ke6
						66.Ng3 Kf6 67.Kg4 Ra1 68.Bd5 Ra5 69.Bf3 Ra1 70.Kf4 Ke6 71.Nc5+
						Kd6 72.Nge4+ Ke7 73.Ke5 Rf1 74.Bg4 Rg1 75.Be6 Re1 76.Bc8 Rc1
						77.Kd4 Rd1+ 78.Nd3 Kf7 79.Ke3 Ra1 80.Kf4 Ke7 81.Nb4 Rc1
						82.Nd5+ Kf7 83.Bd7 Rf1+ 84.Ke5 Ra1 85.Ng5+ Kg6 86.Nf3 Kg7
						87.Bg4 Kg6 88.Nf4+ Kg7 89.Nd4 Re1+ 90.Kf5 Rc1 91.Be2 Re1
						92.Bh5 Ra1 93.Nfe6+ Kh6 94.Be8 Ra8 95.Bc6 Ra1 96.Kf6 Kh7
						97.Ng5+ Kh8 98.Nde6 Ra6 99.Be8 Ra8 100.Bh5 Ra1 101.Bg6 Rf1+
						102.Ke7 Ra1 103.Nf7+ Kg8 104.Nh6+ Kh8 105.Nf5 Ra7+ 106.Kf6 Ra1
						107.Ne3 Re1 108.Nd5 Rg1 109.Bf5 Rf1 110.Ndf4 Ra1 111.Ng6+ Kg8
						112.Ne7+ Kh8 113.Ng5 Ra6+ 114.Kf7 Rf6+ 1/2-1/2*/

			makeMoves([][2]string{
				{`d2`, `d4`}, {`g8`, `f6`},
				{`c2`, `c4`}, {`g7`, `g6`},
				{`b1`, `c3`}, {`f8`, `g7`},
				{`e2`, `e4`}, {`d7`, `d6`},
				{`g1`, `f3`}, /*black*/
			})

			// todo implement converting algebraic O-O, O-O-O to castle, keep in mind that board.Castling()
			// returns exactly amount of POSSIBLE castlings i.e. [0] may be O-O in one position and O-O-O in another.
			// use I property of Castling, in INotation.DecodeCastling()

			fmt.Println("castling I:", b.Castlings(Black)[0].I)
			Expect(b.MakeCastling(b.Castlings(Black)[0])).To(BeTrue())

			makeMoves([][2]string{
				{`f1`, `e2`}, {`e7`, `e5`}, // 6
			})

			fmt.Println("castling I:", b.Castlings(White)[0].I)
			Expect(b.MakeCastling(b.Castlings(White)[0])).To(BeTrue())

			makeMoves([][2]string{
				/*white,*/ {`b8`, `c6`}, // 7
				{`d4`, `d5`}, {`c6`, `e7`},
				{`f3`, `d2`}, {`a7`, `a5`},
				{`a1`, `b1`}, {`f6`, `d7`},
				{`a2`, `a3`}, {`f7`, `f5`},
				{`b2`, `b4`}, {`g8`, `h8`},
				{`f2`, `f3`}, {`e7`, `g8`},
				{`d1`, `c2`}, {`g8`, `f6`},
				{`c3`, `b5`}, {`a5`, `b4`}, // 15
				{`a3`, `b4`}, {`f6`, `h5`},
				{`g2`, `g3`}, {`d7`, `f6`}, // 17
				{`c4`, `c5`}, {`c8`, `d7`},
				{`b1`, `b3`}, {`h5`, `g3`}, // 19
				{`h2`, `g3`}, {`f6`, `h5`},
				{`f3`, `f4`}, {`e5`, `f4`},
				{`c5`, `c6`}, {`b7`, `c6`},
				{`d5`, `c6`}, {`h5`, `g3`},
				{`b3`, `g3`}, {`f4`, `g3`},
				{`c6`, `d7`}, {`g3`, `g2`},
				{`f1`, `f3`}, {`d8`, `d7`},
				{`c1`, `b2`}, {`f5`, `e4`},
				{`f3`, `f8`}, {`a8`, `f8`},
				{`b2`, `g7`}, {`d7`, `g7`},
				{`c2`, `e4`}, {`g7`, `f6`},
				{`d2`, `f3`}, {`f6`, `f4`}, // 31
				{`e4`, `e7`}, {`f8`, `f7`},
				{`e7`, `e6`}, {`f7`, `f6`},
				{`e6`, `e8`}, {`f6`, `f8`},
				{`e8`, `e7`}, {`f8`, `f7`},
				{`e7`, `e6`}, {`f7`, `f6`},
				{`e6`, `b3`}, {`g6`, `g5`},
				{`b5`, `c7`}, {`g5`, `g4`},
				{`c7`, `d5`}, {`f4`, `c1`}, // 39
				{`b3`, `d1`}, {`c1`, `d1`},
				{`e2`, `d1`}, {`f6`, `f5`}, // 41
				{`d5`, `e3`}, {`f5`, `f4`},
				{`f3`, `e1`}, {`f4`, `b4`},
				{`d1`, `g4`}, {`h7`, `h5`},
				{`g4`, `f3`}, {`d6`, `d5`},
				{`e3`, `g2`}, {`h5`, `h4`},
				{`e1`, `d3`}, {`b4`, `a4`},
				{`g2`, `f4`}, {`h8`, `g7`},
				{`g1`, `g2`}, {`g7`, `f6`},
				{`f3`, `d5`}, {`a4`, `a5`},
				{`d5`, `c6`}, {`a5`, `a6`},
				{`c6`, `b7`}, {`a6`, `a3`},
				{`b7`, `e4`}, {`a3`, `a4`}, // 53
				{`e4`, `d5`}, {`a4`, `a5`},
				{`d5`, `c6`}, {`a5`, `a6`},
				{`c6`, `f3`}, {`f6`, `g5`},
				{`f3`, `b7`}, {`a6`, `a1`},
				{`b7`, `c8`}, {`a1`, `a4`},
				{`g2`, `f3`}, {`a4`, `c4`},
				{`c8`, `d7`}, {`g5`, `f6`},
				{`f3`, `g4`}, {`c4`, `d4`},
				{`d7`, `c6`}, {`d4`, `d8`},
				{`g4`, `h4`}, {`d8`, `g8`}, // 63
				{`c6`, `e4`}, {`g8`, `g1`},
				{`f4`, `h5`}, {`f6`, `e6`},
				{`h5`, `g3`}, {`e6`, `f6`},
				{`h4`, `g4`}, {`g1`, `a1`},
				{`e4`, `d5`}, {`a1`, `a5`},
				{`d5`, `f3`}, {`a5`, `a1`},
				{`g4`, `f4`}, {`f6`, `e6`},
				{`d3`, `c5`}, {`e6`, `d6`},
				{`g3`, `e4`}, {`d6`, `e7`},
				{`f4`, `e5`}, {`a1`, `f1`}, // 73
				{`f3`, `g4`}, {`f1`, `g1`},
				{`g4`, `e6`}, {`g1`, `e1`},
				{`e6`, `c8`}, {`e1`, `c1`},
				{`e5`, `d4`}, {`c1`, `d1`},
				{`c5`, `d3`}, {`e7`, `f7`},
				{`d4`, `e3`}, {`d1`, `a1`},
				{`e3`, `f4`}, {`f7`, `e7`},
				{`d3`, `b4`}, {`a1`, `c1`},
				{`b4`, `d5`}, {`e7`, `f7`},
				{`c8`, `d7`}, {`c1`, `f1`},
				{`f4`, `e5`}, {`f1`, `a1`},
				{`e4`, `g5`}, {`f7`, `g6`},
				{`g5`, `f3`}, {`g6`, `g7`},
				{`d7`, `g4`}, {`g7`, `g6`},
				{`d5`, `f4`}, {`g6`, `g7`},
				{`f3`, `d4`}, {`a1`, `e1`},
				{`e5`, `f5`}, {`e1`, `c1`}, // 90
				{`g4`, `e2`}, {`c1`, `e1`},
				{`e2`, `h5`}, {`e1`, `a1`},
				{`f4`, `e6`}, {`g7`, `h6`},
				{`h5`, `e8`}, {`a1`, `a8`},
				{`e8`, `c6`}, {`a8`, `a1`},
				{`f5`, `f6`}, {`h6`, `h7`},
				{`e6`, `g5`}, {`h7`, `h8`},
				{`d4`, `e6`}, {`a1`, `a6`},
				{`c6`, `e8`}, {`a6`, `a8`},
				{`e8`, `h5`}, {`a8`, `a1`},
				{`h5`, `g6`}, {`a1`, `f1`}, // 101
				{`f6`, `e7`}, {`f1`, `a1`},
				{`g5`, `f7`}, {`h8`, `g8`},
				{`f7`, `h6`}, {`g8`, `h8`},
				{`h6`, `f5`}, {`a1`, `a7`},
				{`e7`, `f6`}, {`a7`, `a1`},
				{`f5`, `e3`}, {`a1`, `e1`},
				{`e3`, `d5`}, {`e1`, `g1`},
				{`g6`, `f5`}, {`g1`, `f1`},
				{`d5`, `f4`}, {`f1`, `a1`},
				{`f4`, `g6`}, {`h8`, `g8`},
				{`g6`, `e7`}, {`g8`, `h8`},
				{`e6`, `g5`}, /* draw by 50-moves rule */
			})
			Expect(b.Outcome().Equals(base.NewDrawByMovesRule())).To(BeTrue())
		})
	})
})
