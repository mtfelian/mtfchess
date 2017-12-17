package rect_test

import (
	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coords test", func() {
	It("checks iterations over coords", func() {
		c := []base.Coord{rect.Coord{5, 5}, rect.Coord{4, 5}, rect.Coord{7, 8}}
		func(over rect.Coords) {
			i := 0
			for over.HasNext() {
				nextElement := over.Next().(rect.Coord)
				Expect(over.I()).To(Equal(i))
				Expect(nextElement.Equals(c[over.I()])).To(BeTrue(), "not equals on iter %d", i)
				i++
			}
		}(rect.NewCoords(c))
	})
})
