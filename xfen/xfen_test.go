package xfen_test

import (
	"github.com/mtfelian/mtfchess/xfen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("XFEN test", func() {
	It("is test dummy", func() {
		b, err := xfen.NewFromStandardXFEN("1/2/3")
		Expect(b).To(BeNil())
		Expect(err).To(HaveOccurred())
	})
})
