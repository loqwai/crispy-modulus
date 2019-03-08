package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("Game", func() {
	var (
		g game.Game
	)
	BeforeEach(func() {
		g = game.New()
	})
	It("Should instantiate", func() {
		Expect(g).NotTo(BeNil())
	})
	It("Should instantiate 2 players by default", func() {
		g := game.New()
		players := g.Players
		Expect(len(players)).To(Equal(2))
	})
	Describe("Player", func() {
		var (
			p game.Player
		)
		BeforeEach(func() {
			p = g.Players[0]
		})
		It("Should instantiate a player with 5 cards by default", func() {
			Expect((len(p.MyCards))).To(Equal(5))
		})
	})
})
