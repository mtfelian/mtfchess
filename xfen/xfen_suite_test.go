package xfen

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestXFEN(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XFEN Suite")
}
