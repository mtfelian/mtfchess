package rect

import (
	"fmt"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("XFEN to rectangular board tests", func() {
	It("checks error on totally invalid XFEN", func() {
		b, err := XFEN(`1/2/3`).Board()
		Expect(b).To(BeNil())
		Expect(err).To(HaveOccurred())
	})

	It("checks getting tokens from one position line", func() {
		testCases := []struct {
			line   string
			tokens []string
		}{
			{"4QRqr5", []string{"4", "Q", "R", "q", "r", "5"}},
			{"QRqr5", []string{"Q", "R", "q", "r", "5"}},
			{"10QRqr5", []string{"10", "Q", "R", "q", "r", "5"}},
			{"QRqr10", []string{"Q", "R", "q", "r", "10"}},
			{"QR10qr5", []string{"Q", "R", "10", "q", "r", "5"}},
			{"5", []string{"5"}},
			{"10", []string{"10"}},
			{"", []string{}},
		}

		for i, testCase := range testCases {
			By(fmt.Sprintf("Checking testCase %v at index %d...", testCase, i))
			tokens := getPosLineTokens(testCase.line)
			Expect(tokens).To(Equal(testCase.tokens))
		}
	})

	Context("valid XFEN, chess960", func() {
		var b *Board
		var resetBoard func()
		JustBeforeEach(func() { resetBoard() })
		resetBoard = func() {
			b = NewEmptyStandardChessBoard()
			// set rook initial coords to enable castling
			//b.SetRookInitialCoords(White, 0, Coord{1, 1}) // should not set it, rook moved
			b.SetRookInitialCoords(White, 1, Coord{7, 1})
			b.SetRookInitialCoords(Black, 0, Coord{1, 8})
			b.SetRookInitialCoords(Black, 1, Coord{7, 8})
		}
		var wr1, wr2, wk, br1, br2, bk base.IPiece
		setupPosition := func() {
			wr1, wr2, wk = NewRook(White), NewRook(White), NewKing(White)
			br1, br2, bk = NewRook(Black), NewRook(Black), NewKing(Black)

			b.PlacePiece(Coord{3, 1}, NewBishop(White))
			b.PlacePiece(Coord{4, 1}, NewKnight(White))
			b.PlacePiece(Coord{2, 2}, NewPawn(White))
			b.PlacePiece(Coord{3, 2}, NewPawn(White))
			b.PlacePiece(Coord{4, 2}, NewPawn(White))
			b.PlacePiece(Coord{5, 2}, NewPawn(White))
			b.PlacePiece(Coord{6, 2}, NewPawn(White))
			b.PlacePiece(Coord{3, 3}, NewKnight(White))
			b.PlacePiece(Coord{6, 3}, NewBishop(White))
			b.PlacePiece(Coord{1, 4}, NewPawn(White))
			b.PlacePiece(Coord{6, 5}, NewBishop(Black))
			b.PlacePiece(Coord{7, 5}, NewKnight(Black))
			b.PlacePiece(Coord{4, 6}, NewPawn(Black))
			b.PlacePiece(Coord{7, 6}, NewPawn(Black))
			b.PlacePiece(Coord{1, 7}, NewPawn(Black))
			b.PlacePiece(Coord{2, 7}, NewPawn(Black))
			b.PlacePiece(Coord{3, 7}, NewPawn(Black))
			b.PlacePiece(Coord{5, 7}, NewPawn(Black))
			b.PlacePiece(Coord{6, 7}, NewPawn(Black))
			b.PlacePiece(Coord{8, 7}, NewPawn(Black))
			b.PlacePiece(Coord{2, 8}, NewKnight(Black))

			b.PlacePiece(Coord{8, 1}, wr1)
			wr1.MarkMoved()
			b.PlacePiece(Coord{7, 1}, wr2)
			b.PlacePiece(Coord{5, 1}, wk)
			b.PlacePiece(Coord{1, 8}, br1)
			b.PlacePiece(Coord{7, 8}, br2)
			b.PlacePiece(Coord{5, 8}, bk)

			b.Settings().MoveOrder = false
			b.SetSideToMove(White)
			b.SetMoveNumber(11)
			b.SetHalfMoveCount(4)
		}

		It("checks that parsed board is equal to hard-coded board", func() {
			setupPosition()
			var input XFEN = `rn2k1r1/ppp1pp1p/3p2p1/5bn1/P7/2N2B2/1PPPPP2/2BNK1RR w Gkq - 4 11`
			parsedBoard, err := input.Board()
			Expect(err).NotTo(HaveOccurred())
			Expect(parsedBoard).NotTo(BeNil())
			Expect(b.Equals(parsedBoard)).To(BeTrue())
		})
	})

})

var _ = Describe("Rectangular board to XFEN tests", func() {
	var b *Board
	var resetBoard func()
	JustBeforeEach(func() { resetBoard() })
	resetBoard = func() {
		b = NewEmptyStandardChessBoard()
		// set rook initial coords to enable castling
		//b.SetRookInitialCoords(White, 0, Coord{1, 1}) // should not set it, rook moved
		b.SetRookInitialCoords(White, 1, Coord{7, 1})
		b.SetRookInitialCoords(Black, 0, Coord{1, 8})
		b.SetRookInitialCoords(Black, 1, Coord{7, 8})
	}
	var wr1, wr2, wk, br1, br2, bk base.IPiece
	setupPosition := func() {
		wr1, wr2, wk = NewRook(White), NewRook(White), NewKing(White)
		br1, br2, bk = NewRook(Black), NewRook(Black), NewKing(Black)

		b.PlacePiece(Coord{3, 1}, NewBishop(White))
		b.PlacePiece(Coord{4, 1}, NewKnight(White))
		b.PlacePiece(Coord{2, 2}, NewPawn(White))
		b.PlacePiece(Coord{3, 2}, NewPawn(White))
		b.PlacePiece(Coord{4, 2}, NewPawn(White))
		b.PlacePiece(Coord{5, 2}, NewPawn(White))
		b.PlacePiece(Coord{6, 2}, NewPawn(White))
		b.PlacePiece(Coord{3, 3}, NewKnight(White))
		b.PlacePiece(Coord{6, 3}, NewBishop(White))
		b.PlacePiece(Coord{1, 4}, NewPawn(White))
		b.PlacePiece(Coord{6, 5}, NewBishop(Black))
		b.PlacePiece(Coord{7, 5}, NewKnight(Black))
		b.PlacePiece(Coord{4, 6}, NewPawn(Black))
		b.PlacePiece(Coord{7, 6}, NewPawn(Black))
		b.PlacePiece(Coord{1, 7}, NewPawn(Black))
		b.PlacePiece(Coord{2, 7}, NewPawn(Black))
		b.PlacePiece(Coord{3, 7}, NewPawn(Black))
		b.PlacePiece(Coord{5, 7}, NewPawn(Black))
		b.PlacePiece(Coord{6, 7}, NewPawn(Black))
		b.PlacePiece(Coord{8, 7}, NewPawn(Black))
		b.PlacePiece(Coord{2, 8}, NewKnight(Black))

		b.PlacePiece(Coord{8, 1}, wr1)
		wr1.MarkMoved()
		b.PlacePiece(Coord{7, 1}, wr2)
		b.PlacePiece(Coord{5, 1}, wk)
		b.PlacePiece(Coord{1, 8}, br1)
		b.PlacePiece(Coord{7, 8}, br2)
		b.PlacePiece(Coord{5, 8}, bk)

		b.Settings().MoveOrder = false
		b.SetSideToMove(White)
		b.SetMoveNumber(11)
		b.SetHalfMoveCount(4)
	}

	It("checks converting board to XFEN", func() {
		setupPosition()
		xfen := NewXFEN(b)
		Expect(xfen).To(Equal(XFEN(`rn2k1r1/ppp1pp1p/3p2p1/5bn1/P7/2N2B2/1PPPPP2/2BNK1RR w Gkq - 4 11`)))
	})
})
