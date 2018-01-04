package piece_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPiece(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Piece Suite")
}
