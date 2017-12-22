package piece_test

import (
	"fmt"
	"sort"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pawn test", func() {
	w, h := 5, 6
	var b base.IBoard
	BeforeEach(func() { b = rect.NewEmptyBoard(w, h) })

	It("generates moves", func() {
		wp, bn := piece.NewPawn(White), piece.NewKnight(Black)
		b.PlacePiece(rect.Coord{2, 2}, wp)
		b.PlacePiece(rect.Coord{1, 3}, bn)

		a := wp.Attacks(b)
		fmt.Println("###", a)

		d := wp.Destinations(b)
		fmt.Println(">>>", d)

		Expect(d.Len()).To(Equal(2))
		sort.Sort(d)
		Expect(d.Equals(rect.NewCoords([]base.ICoord{rect.Coord{1, 3}, rect.Coord{2, 3}}))).To(BeTrue())
	})
})
