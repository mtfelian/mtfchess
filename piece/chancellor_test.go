package piece_test

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chancellor test", func() {
	var b base.IBoard
	resetBoard := func() { b = rect.NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wc, wn, bc := piece.NewChancellor(White), piece.NewKnight(White), piece.NewChancellor(Black)
		b.PlacePiece(rect.Coord{4, 2}, wc)
		b.PlacePiece(rect.Coord{4, 4}, wn)
		b.PlacePiece(rect.Coord{4, 1}, bc)

		d := wc.Destinations(b)
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{4, 1}, rect.Coord{1, 2}, rect.Coord{2, 2},
			rect.Coord{3, 2}, rect.Coord{5, 2}, rect.Coord{2, 3}, rect.Coord{4, 3},
			rect.Coord{3, 4}, rect.Coord{5, 4},
		}))).To(BeTrue())
	})

	It("attacks right cells", func() {
		wc, bq := piece.NewChancellor(White), piece.NewQueen(Black)
		bn, wn, wk := piece.NewKnight(Black), piece.NewKnight(White), piece.NewKing(White)
		b.PlacePiece(rect.Coord{1, 2}, bq)
		b.PlacePiece(rect.Coord{2, 2}, wc)
		b.PlacePiece(rect.Coord{3, 2}, wk)
		b.PlacePiece(rect.Coord{2, 1}, bn)
		b.PlacePiece(rect.Coord{2, 5}, wn)

		attacking := wc.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{2, 1}, rect.Coord{4, 1}, rect.Coord{1, 2}, rect.Coord{3, 2},
			rect.Coord{2, 3}, rect.Coord{4, 3}, rect.Coord{1, 4}, rect.Coord{2, 4},
			rect.Coord{3, 4}, rect.Coord{2, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(rect.Coord{2, 5}, wc)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(rect.Coord{3, 2}, wc)).To(BeFalse(), "captured own piece")
		Expect(b.MakeMove(rect.Coord{2, 6}, wc)).To(BeFalse(), "jumped over own piece")
		Expect(b.MakeMove(rect.Coord{1, 2}, wc)).To(BeTrue(), "can't capture")
	})

	It("makes legal moves", func() {
		var wc, bc base.IPiece
		testReset := func() {
			resetBoard()
			wc, bc = piece.NewChancellor(White), piece.NewChancellor(Black)
			b.PlacePiece(rect.Coord{2, 1}, wc)
			b.PlacePiece(rect.Coord{4, 1}, bc)
		}
		testReset()

		destinations := wc.Destinations(b)
		sort.Sort(destinations)
		Expect(destinations.Equals(rect.NewCoords([]base.ICoord{
			rect.Coord{1, 1}, rect.Coord{3, 1}, rect.Coord{4, 1},
			rect.Coord{2, 2}, rect.Coord{4, 2}, rect.Coord{1, 3},
			rect.Coord{2, 3}, rect.Coord{3, 3}, rect.Coord{2, 4},
			rect.Coord{2, 5}, rect.Coord{2, 6},
		}))).To(BeTrue())

		bcCoord, wcCoord := bc.Coord().Copy(), wc.Coord().Copy()
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wc)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wcCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wc))
			if !bcCoord.Equals(d) { // if not capture
				// not captured piece still stands
				Expect(b.Piece(bcCoord)).To(Equal(bc))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wc.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(bc.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("don't makes illegal moves", func() {
		var wc, bc base.IPiece
		testReset := func() {
			resetBoard()
			wc, bc = piece.NewChancellor(White), piece.NewChancellor(Black)
			b.PlacePiece(rect.Coord{2, 1}, wc)
			b.PlacePiece(rect.Coord{4, 1}, bc)
		}
		testReset()

		destinations := rect.NewCoords([]base.ICoord{rect.Coord{3, 2}, rect.Coord{5, 1}, wc.Coord()})
		for destinations.HasNext() {
			d, c := destinations.Next().(rect.Coord), wc.Coord()
			dCellCopy := b.Cell(d).Copy(b)
			Expect(b.MakeMove(d, wc)).To(BeFalse(), "failed at offset %d", destinations.I())
			// check source cell to contain unmoved piece
			Expect(b.Piece(c)).To(Equal(wc))

			// check that destination cell was not changed
			if dCellCopy.Piece() == nil {
				Expect(b.Piece(d)).To(BeNil())
			} else {
				Expect(b.Piece(d)).To(Equal(dCellCopy.Piece()))
			}

			// check another cell to contain another piece
			Expect(b.Piece(bc.Coord())).To(Equal(bc))

			testReset()
		}
	})
})
