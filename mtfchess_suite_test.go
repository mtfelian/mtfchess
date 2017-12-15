package mtfchess_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMtfchess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MTFchess Suite")
}
