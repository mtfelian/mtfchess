package mtfchess_test

import (
	"fmt"

	. "github.com/mtfelian/mtfchess"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board test", func() {
	w, h := 6, 8
	var b *Board

	BeforeEach(func() { b = NewEmptyBoard(w, h) })

	It("checks board creation", func() {
		fmt.Println(b)
		Expect(b.Width()).To(Equal(w))
		Expect(b.Height()).To(Equal(h))
	})

	It("checks knight moves generation", func() {
		whiteKnight1 := NewKnight(b, White)
		whiteKnight2 := NewKnight(b, White)
		blackKnight := NewKnight(b, Black)
		b.PlacePiece(2, 1, whiteKnight1)
		b.PlacePiece(3, 3, whiteKnight2)
		b.PlacePiece(4, 2, blackKnight)
		fmt.Println(b)
		o := whiteKnight1.Offsets()
		fmt.Println(o)
		Expect(o).To(HaveLen(2))
		Expect(o).To(Equal(Offsets{{-1, 2}, {2, 1}}))
	})
})
