package mtfchess_test

import (
	"fmt"

	. "github.com/mtfelian/mtfchess"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing with Ginkgo", func() {
})

var _ = Describe("Board test", func() {
	It("checks board creation", func() {
		b := NewEmptyBoard(5, 6)
		fmt.Println(b)
		Expect(true).To(BeTrue())
	})
})
