package rect_test

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board test", func() {
	var w, h int
	var b base.IBoard

	BeforeEach(func() {
		b = rect.NewEmptyTestBoard()
		bC := b.Dim().(rect.Coord)
		w, h = bC.X, bC.Y
	})

	It("checks board width and height", func() {
		Expect(b.Dim().(rect.Coord).X).To(Equal(w))
		Expect(b.Dim().(rect.Coord).Y).To(Equal(h))
	})

	Describe("find pieces", func() {
		var wn1, wn2, wn3, bn1, bn2, bn3, wk, bk base.IPiece
		BeforeEach(func() {
			wn1, wn2, wn3 = piece.NewKnight(White), piece.NewKnight(White), piece.NewKnight(White)
			bn1, bn2, bn3 = piece.NewKnight(Black), piece.NewKnight(Black), piece.NewKnight(Black)
			wk, bk = piece.NewKing(White), piece.NewKing(Black)
			b.PlacePiece(rect.Coord{1, 1}, wn1)
			b.PlacePiece(rect.Coord{1, 2}, wn2)
			b.PlacePiece(rect.Coord{3, 4}, wn3)
			b.PlacePiece(rect.Coord{5, 5}, bn1)
			b.PlacePiece(rect.Coord{5, 6}, bn2)
			b.PlacePiece(rect.Coord{4, 3}, bn3)
			b.PlacePiece(rect.Coord{2, 1}, wk)
			b.PlacePiece(rect.Coord{5, 4}, bk)
		})
		It("normally", func() {
			filter := rect.PieceFilter{ // find all white knights
				PieceFilter: base.PieceFilter{
					Colours: []Colour{White},
					Names:   []string{piece.NewKnight(Transparent).Name()},
				},
			}
			coords := b.FindPieces(filter)
			Expect(coords).To(Equal(base.Pieces{wn3, wn2, wn1}))
		})

		It("is with piece / board condition", func() {
			notOnEdge := func(p base.IPiece) bool {
				x, y := p.Coord().(rect.Coord).X, p.Coord().(rect.Coord).Y
				w, h := b.Dim().(rect.Coord).X, b.Dim().(rect.Coord).Y
				return x > 1 && y > 1 && x < w && y < h
			}
			filter := rect.PieceFilter{ // find all knights
				PieceFilter: base.PieceFilter{
					Names:     []string{piece.NewKnight(Transparent).Name()},
					Condition: notOnEdge,
				},
			}

			pieces := b.FindPieces(filter)
			Expect(pieces).To(Equal(base.Pieces{wn3, bn3}))
		})
	})

	Describe("find attacked cells", func() {
		It("works", func() {
			wn, bn := piece.NewKnight(White), piece.NewKnight(Black)
			wk, bk := piece.NewKing(White), piece.NewKing(Black)
			b.PlacePiece(rect.Coord{1, 1}, bk)
			b.PlacePiece(rect.Coord{2, 4}, wn)
			b.PlacePiece(rect.Coord{5, 5}, wk)
			b.PlacePiece(rect.Coord{4, 4}, bn)

			attackedByWhite := b.FindAttackedCellsBy(rect.PieceFilter{
				PieceFilter: base.PieceFilter{Colours: []Colour{White}},
			})
			sort.Sort(attackedByWhite)
			Expect(attackedByWhite.Equals(rect.NewCoords([]base.ICoord{
				rect.Coord{1, 2}, rect.Coord{3, 2}, rect.Coord{4, 3},
				rect.Coord{4, 4}, rect.Coord{5, 4}, rect.Coord{4, 5},
				rect.Coord{1, 6}, rect.Coord{3, 6}, rect.Coord{4, 6}, rect.Coord{5, 6},
			}))).To(BeTrue())

			attackedByBlack := b.FindAttackedCellsBy(rect.PieceFilter{
				PieceFilter: base.PieceFilter{Colours: []Colour{Black}},
			})
			sort.Sort(attackedByBlack)
			Expect(attackedByBlack.Equals(rect.NewCoords([]base.ICoord{
				rect.Coord{2, 1}, rect.Coord{1, 2}, rect.Coord{2, 2},
				rect.Coord{3, 2}, rect.Coord{5, 2}, rect.Coord{2, 3},
				rect.Coord{2, 5}, rect.Coord{3, 6}, rect.Coord{5, 6},
			}))).To(BeTrue())
		})
	})

	Describe("check move order control", func() {
		It("is enabled", func() {
			b.Settings().MoveOrder = true
			b.SetSideToMove(White)

			Expect(b.MoveNumber()).To(BeNumerically("==", 1))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", 1))
			var mn, hmc = 3, 6
			b.SetMoveNumber(mn)
			b.SetHalfMoveCount(hmc)
			Expect(b.MoveNumber()).To(BeNumerically("==", mn))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc))
			Expect(b.SideToMove()).To(Equal(White))

			wr, br := piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{1, 1}, wr)
			b.PlacePiece(rect.Coord{3, 3}, br)

			Expect(b.MakeMove(rect.Coord{3, 4}, br)).To(BeFalse(), "white to move, but black moved")
			Expect(b.MakeMove(rect.Coord{1, 2}, wr)).To(BeTrue(), "white to move, but white can't move")
			Expect(b.MoveNumber()).To(BeNumerically("==", mn))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc+1))
			Expect(b.SideToMove()).To(Equal(Black))

			Expect(b.MakeMove(rect.Coord{1, 1}, wr)).To(BeFalse(), "black to move, but white moved")
			Expect(b.MakeMove(rect.Coord{3, 5}, br)).To(BeTrue(), "black to move, but black can't move")
			Expect(b.MoveNumber()).To(BeNumerically("==", mn+1))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc+2))
			Expect(b.SideToMove()).To(Equal(White))
		})

		It("is disabled", func() {
			b.Settings().MoveOrder = false
			b.SetSideToMove(White)

			Expect(b.MoveNumber()).To(BeNumerically("==", 1))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", 1))
			var mn, hmc = 3, 6
			b.SetMoveNumber(mn)
			b.SetHalfMoveCount(hmc)
			Expect(b.MoveNumber()).To(BeNumerically("==", mn))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc))
			Expect(b.SideToMove()).To(Equal(White))

			wr, br := piece.NewRook(White), piece.NewRook(Black)
			b.PlacePiece(rect.Coord{1, 1}, wr)
			b.PlacePiece(rect.Coord{3, 3}, br)

			Expect(b.MakeMove(rect.Coord{3, 4}, br)).To(BeTrue(), "1 white to move, ordering disabled, but move failed")
			Expect(b.MoveNumber()).To(BeNumerically("==", mn+1))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc+1))
			Expect(b.SideToMove()).To(Equal(Black))
			Expect(b.MakeMove(rect.Coord{1, 2}, wr)).To(BeTrue(), "2 black to move, ordering disabled, but move failed")
			Expect(b.MoveNumber()).To(BeNumerically("==", mn+1))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc+2))
			Expect(b.SideToMove()).To(Equal(White))
			Expect(b.MakeMove(rect.Coord{1, 1}, wr)).To(BeTrue(), "3 white to move, ordering disabled, but move failed")
			Expect(b.MoveNumber()).To(BeNumerically("==", mn+1))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc+3))
			Expect(b.SideToMove()).To(Equal(Black))
			Expect(b.MakeMove(rect.Coord{3, 5}, br)).To(BeTrue(), "4 black to move, ordering disabled, but move failed")
			Expect(b.MoveNumber()).To(BeNumerically("==", mn+2))
			Expect(b.HalfMoveCount()).To(BeNumerically("==", hmc+4))
			Expect(b.SideToMove()).To(Equal(White))
		})
	})

	Describe("check LegalMoves", func() {
		var b base.IBoard
		var resetBoard func()
		JustBeforeEach(func() { resetBoard() })
		resetBoard = func() {
			b = rect.NewEmptyStandardChessBoard()
			// set rook initial coords to enable castling
			//b.SetRookInitialCoords(White, 0, rect.Coord{1, 1}) // should not set it, rook moved
			b.SetRookInitialCoords(White, 1, rect.Coord{7, 1})
			b.SetRookInitialCoords(Black, 0, rect.Coord{1, 8})
			b.SetRookInitialCoords(Black, 1, rect.Coord{7, 8})
		}
		var wr1, wr2, wk, br1, br2, bk base.IPiece
		setupPosition := func() {
			wr1, wr2, wk = piece.NewRook(White), piece.NewRook(White), piece.NewKing(White)
			br1, br2, bk = piece.NewRook(Black), piece.NewRook(Black), piece.NewKing(Black)

			b.PlacePiece(rect.Coord{3, 1}, piece.NewBishop(White))
			b.PlacePiece(rect.Coord{4, 1}, piece.NewKnight(White))
			b.PlacePiece(rect.Coord{2, 2}, piece.NewPawn(White))
			b.PlacePiece(rect.Coord{3, 2}, piece.NewPawn(White))
			b.PlacePiece(rect.Coord{4, 2}, piece.NewPawn(White))
			b.PlacePiece(rect.Coord{5, 2}, piece.NewPawn(White))
			b.PlacePiece(rect.Coord{6, 2}, piece.NewPawn(White))
			b.PlacePiece(rect.Coord{3, 3}, piece.NewKnight(White))
			b.PlacePiece(rect.Coord{6, 3}, piece.NewBishop(White))
			b.PlacePiece(rect.Coord{1, 4}, piece.NewPawn(White))
			b.PlacePiece(rect.Coord{6, 5}, piece.NewBishop(Black))
			b.PlacePiece(rect.Coord{7, 5}, piece.NewKnight(Black))
			b.PlacePiece(rect.Coord{4, 6}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{7, 6}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{1, 7}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{2, 7}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{3, 7}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{5, 7}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{6, 7}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{8, 7}, piece.NewPawn(Black))
			b.PlacePiece(rect.Coord{2, 8}, piece.NewKnight(Black))

			b.PlacePiece(rect.Coord{8, 1}, wr1)
			wr1.MarkMoved()
			b.PlacePiece(rect.Coord{7, 1}, wr2)
			b.PlacePiece(rect.Coord{5, 1}, wk)
			b.PlacePiece(rect.Coord{1, 8}, br1)
			b.PlacePiece(rect.Coord{7, 8}, br2)
			b.PlacePiece(rect.Coord{5, 8}, bk)

			b.Settings().MoveOrder = false
			b.SetSideToMove(White)
			b.SetMoveNumber(11)
			b.SetHalfMoveCount(4)
		}

		It("checks legal moves calculation", func() {
			setupPosition()
			moves := b.LegalMoves(rect.NewLongAlgebraicNotation())
			expectedMoves := []string{
				"O-O",
				"a4-a5", "b2-b3", "b2-b4", "d2-d3", "d2-d4", "e2-e3", "e2-e4",
				"Nc3-e4", "Nc3-a2", "Nc3-d5", "Nc3-b5", "Nc3-b1",
				"Bf3-g4", "Bf3-h5", "Bf3-e4", "Bf3-d5", "Bf3-c6", "Bf3xb7", "Bf3-g2",
				"Nd1-e3",
				"Ke1-f1",
				"Rg1-f1", "Rg1-g2", "Rg1-g3", "Rg1-g4", "Rg1xg5",
				"Rh1-h2", "Rh1-h3", "Rh1-h4", "Rh1-h5", "Rh1-h6", "Rh1xh7",
			}
			sort.Strings(moves)
			sort.Strings(expectedMoves)
			Expect(moves).To(Equal(expectedMoves))
		})
	})
})
