package rect_test

import (
	"fmt"

	"github.com/mtfelian/mtfchess/base"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Algebraic test", func() {
	It("checks converting algebraic coord to rect.Coord", func() {
		type testCase struct {
			algebraic    string
			coord        base.ICoord
			errorOccured bool
		}
		testTable := []testCase{
			{"a1", rect.Coord{1, 1}, false},
			{"a8", rect.Coord{1, 8}, false},
			{"i1", rect.Coord{9, 1}, false},
			{"i8", rect.Coord{9, 8}, false},
			{"e4", rect.Coord{5, 4}, false},
			{"E4", rect.Coord{5, 4}, false},
			{"4e", nil, true},
			{"4", nil, true},
			{"e", nil, true},
			{"", nil, true},
			{"e10", rect.Coord{5, 10}, false},
			{"10e", nil, true},
		}

		for i, entry := range testTable {
			By(fmt.Sprintf("Checking entry %v at index %d", entry, i))
			coord, err := rect.FromAlgebraic(entry.algebraic)
			Expect((err != nil) == entry.errorOccured).To(BeTrue())
			if entry.coord == nil {
				Expect(coord).To(BeNil())
			} else {
				Expect(coord).To(Equal(entry.coord))
			}
		}
	})
})
