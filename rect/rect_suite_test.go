package rect_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rectangular Board Suite")
}
