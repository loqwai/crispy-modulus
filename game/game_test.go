package game_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("game", func() {

	Describe("Game", func() {
		var g game.Game

		BeforeEach(func() {
			rand.Seed(0)
			g = game.New()
		})

		It("Should instantiate 2 players by default", func() {
			players := g.GetState().Players
			Expect(len(players)).To(Equal(2))
		})
	})

	Describe("Player", func() {
		var p game.Player

		BeforeEach(func() {
			p = game.NewPlayer()
		})

		It("Should instantiate a player with 5 cards by default", func() {
			Expect((len(p.MyCards))).To(Equal(5))
		})

		It("Should populate the hand with numbers between 1-10", func() {
			for i := 0; i < 5; i++ {
				Expect(p.MyCards[i]).To(BeNumerically(">", 0))
				Expect(p.MyCards[i]).To(BeNumerically("<", 11))
			}
		})

		It("Should not populate the hand with the same card twice", func() {
			for i := 0; i < 100; i++ {
				p2 := game.NewPlayer()
				for j := 0; j < 5; j++ {
					for k := 0; k < 5; k++ {
						if j == k {
							continue
						}
						Expect(p2.MyCards[j]).NotTo(Equal(p2.MyCards[k]))
					}
				}
			}
		})

		It("Should instantiate a player with no opponent cards by default", func() {
			Expect(len(p.TheirCards)).To(Equal(0))
		})
	})
})
