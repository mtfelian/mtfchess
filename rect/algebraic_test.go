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
		testCases := []struct {
			algebraic    string
			coord        base.ICoord
			errorOccured bool
		}{
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

		for i, testCase := range testCases {
			By(fmt.Sprintf("Checking testCase %v at index %d...", testCase, i))
			coord, err := rect.FromAlgebraic(testCase.algebraic)
			Expect((err != nil) == testCase.errorOccured).To(BeTrue())
			if testCase.coord == nil {
				Expect(coord).To(BeNil())
			} else {
				Expect(coord).To(Equal(testCase.coord))
			}
		}
	})

	It("checks converting rect.Coord to algebraic coord", func() {
		testCases := []struct {
			coord     base.ICoord
			algebraic string
		}{
			{rect.Coord{1, 1}, "a1"},
			{rect.Coord{1, 8}, "a8"},
			{rect.Coord{9, 1}, "i1"},
			{rect.Coord{9, 8}, "i8"},
			{rect.Coord{5, 4}, "e4"},
			{rect.Coord{5, 10}, "e10"},
			{nil, ""},
		}

		for i, testCase := range testCases {
			By(fmt.Sprintf("Checking testCase %v at index %d...", testCase, i))
			algebraic := rect.ToAlgebraic(testCase.coord)
			Expect(algebraic).To(Equal(testCase.algebraic))
		}
	})
})
