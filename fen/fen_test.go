package fen_test

import (
	"github.com/mtfelian/mtfchess/fen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FEN test", func() {
	It("is test dummy", func() {
		b, err := fen.NewFromStandardFEN("1/2/3")
		Expect(b).To(BeNil())
		Expect(err).To(HaveOccurred())
	})
})
