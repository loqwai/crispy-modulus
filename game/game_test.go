package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("Game", func() {
	It("Should instantiate", func() {
		g := game.New()
		Expect(g).NotTo(BeNil())
	})
})
