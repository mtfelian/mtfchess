package rect_test

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coords test", func() {
	It("checks iterations over coords", func() {
		data := []base.ICoord{rect.Coord{5, 5}, rect.Coord{4, 5}, rect.Coord{7, 8}}
		func(coords rect.Coords) {
			i := 0
			for coords.HasNext() {
				nextCoord := coords.Next().(rect.Coord)
				Expect(coords.I()).To(Equal(i))
				Expect(nextCoord.Equals(data[coords.I()])).To(BeTrue(), "not equals on iter %d", i)
				i++
			}
		}(rect.NewCoords(data))
	})
})
