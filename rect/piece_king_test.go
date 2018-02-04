package rect

import (
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("King test", func() {
	var b base.IBoard
	resetBoard := func() { b = NewEmptyTestBoard() }
	BeforeEach(func() { resetBoard() })

	It("generates moves", func() {
		wk, wn, bn := NewKing(White), NewKnight(White), NewKnight(Black)
		b.PlacePiece(Coord{2, 2}, wk)
		b.PlacePiece(Coord{2, 3}, wn)
		b.PlacePiece(Coord{1, 1}, bn)
		d := wk.Destinations(b)
		sort.Sort(d)
		Expect(d.Equals(NewCoords([]base.ICoord{
			Coord{1, 1}, Coord{2, 1}, Coord{3, 1},
			Coord{1, 2}, Coord{1, 3}, Coord{3, 3},
		}))).To(BeTrue())
	})

	It("attacks right cells, can release check by capture", func() {
		wn, bq := NewKnight(White), NewQueen(Black)
		bn, wr, wk := NewKnight(Black), NewRook(White), NewKing(White)
		b.PlacePiece(Coord{2, 5}, bq)
		b.PlacePiece(Coord{1, 5}, wn)
		b.PlacePiece(Coord{1, 4}, wk)
		b.PlacePiece(Coord{4, 2}, bn)
		b.PlacePiece(Coord{2, 4}, wr)

		attacking := wk.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{1, 3}, Coord{2, 3}, Coord{2, 4},
			Coord{1, 5}, Coord{2, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{1, 5}, wk)).To(BeFalse(), "captured own piece and still in check")
		Expect(b.MakeMove(Coord{2, 4}, wk)).To(BeFalse(), "captured piece but still in check")
		Expect(b.MakeMove(Coord{2, 3}, wk)).To(BeFalse(), "king still in check (from bn)")

		// successful capture releasing check
		Expect(b.MakeMove(Coord{2, 5}, wk)).To(BeTrue(), "can't capture releasing check")
	})

	It("attacks right cells, can release check by pinning another piece", func() {
		wn, bq := NewKnight(White), NewQueen(Black)
		wr, wk := NewRook(White), NewKing(White)
		b.PlacePiece(Coord{2, 5}, bq)
		b.PlacePiece(Coord{1, 5}, wn)
		b.PlacePiece(Coord{1, 4}, wk)
		b.PlacePiece(Coord{2, 4}, wr)

		attacking := wk.Attacks(b)
		sort.Sort(attacking)
		Expect(attacking.Equals(NewCoords([]base.ICoord{
			Coord{1, 3}, Coord{2, 3}, Coord{2, 4},
			Coord{1, 5}, Coord{2, 5},
		}))).To(BeTrue())

		Expect(b.MakeMove(Coord{1, 5}, wk)).To(BeFalse(), "captured own piece and still in check")
		Expect(b.MakeMove(Coord{2, 4}, wk)).To(BeFalse(), "captured piece but still in check")

		// pin another piece (wr) and release check
		Expect(b.MakeMove(Coord{2, 3}, wk)).To(BeTrue(), "can't release check by pin")
	})

	It("makes legal moves", func() {
		var wk, bk, br base.IPiece
		testReset := func() {
			resetBoard()
			wk, bk, br = NewKing(White), NewKing(Black), NewRook(Black)
			b.PlacePiece(Coord{2, 3}, wk)
			b.PlacePiece(Coord{3, 4}, br)
			b.PlacePiece(Coord{4, 2}, bk)
		}
		testReset()
		destinations := wk.Destinations(b)

		wkCoord, brCoord := wk.Coord().Copy(), br.Coord().Copy()
		for destinations.HasNext() {
			d := destinations.Next().(base.ICoord)
			Expect(b.MakeMove(d, wk)).To(BeTrue(), "failed at destination %d", destinations.I())
			// check source cell to be empty
			Expect(b.Piece(wkCoord)).To(BeNil())
			// check destination cell to contain new piece
			Expect(b.Piece(d)).To(Equal(wk))
			if !brCoord.Equals(d) { // if not capture
				// then there should be another piece
				Expect(b.Piece(brCoord)).To(Equal(br))
			} else { // capture
				// capturing piece's coords is destination
				Expect(wk.Coord()).To(Equal(d))
				// captured piece's coords is nil
				Expect(br.Coord()).To(BeNil())
			}

			testReset()
		}
	})

	It("can't go under check of rook", func() {
		wk, bk, br := NewKing(White), NewKing(Black), NewRook(Black)
		b.PlacePiece(Coord{2, 4}, wk)
		b.PlacePiece(Coord{4, 3}, bk)
		b.PlacePiece(Coord{4, 5}, br)

		d, c := Coord{2, 5}, wk.Coord()
		Expect(b.MakeMove(d, wk)).To(BeFalse(), "gone under rook's check!")
		Expect(b.Piece(c)).To(Equal(wk))
		Expect(b.Piece(d)).To(BeNil())
		Expect(b.Piece(br.Coord())).To(Equal(br))
		Expect(b.Piece(bk.Coord())).To(Equal(bk))

		wkDestinations := wk.Destinations(b)
		Expect(wkDestinations.Len()).To(Equal(3))
		sort.Sort(wkDestinations)
		Expect(wkDestinations.Equals(NewCoords([]base.ICoord{
			Coord{1, 3}, Coord{2, 3}, Coord{1, 4},
		}))).To(BeTrue())
	})

	It("can't go to a opponent's king neighbour cell", func() {
		wk, bk, br := NewKing(White), NewKing(Black), NewRook(Black)
		b.PlacePiece(Coord{2, 4}, wk)
		b.PlacePiece(Coord{4, 3}, bk)
		b.PlacePiece(Coord{4, 5}, br)

		d, c := Coord{3, 4}, wk.Coord()
		Expect(b.MakeMove(d, wk)).To(BeFalse(), "two kings on a neighbour cells!")
		Expect(b.Piece(c)).To(Equal(wk))
		Expect(b.Piece(d)).To(BeNil())
		Expect(b.Piece(br.Coord())).To(Equal(br))
		Expect(b.Piece(bk.Coord())).To(Equal(bk))
	})

	It("can't go into same cell (skip move)", func() {
		wk, bk, br := NewKing(White), NewKing(Black), NewRook(Black)
		b.PlacePiece(Coord{2, 4}, wk)
		b.PlacePiece(Coord{4, 3}, bk)
		b.PlacePiece(Coord{4, 5}, br)

		c := wk.Coord()
		Expect(b.MakeMove(c, wk)).To(BeFalse(), "performed a move into same cell!")
		Expect(b.Piece(c)).To(Equal(wk))
		Expect(b.Piece(br.Coord())).To(Equal(br))
		Expect(b.Piece(bk.Coord())).To(Equal(bk))
	})
})
