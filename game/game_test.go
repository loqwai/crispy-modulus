package game_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("Game", func() {
	var (
		g game.Game
	)

	BeforeEach(func() {
		rand.Seed(0)
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

		It("Should populate the hand with numbers between 1-10", func() {
			for i := 0; i < 5; i++ {
				Expect(p.MyCards[i] > 0).To(BeTrue())
				Expect(p.MyCards[i] < 11).To(BeTrue())
			}
		})

		It("Should not populate the hand with the same card twice", func() {

			for i := 0; i < 1000; i++ {
				p2 := game.NewPlayer()
				for j := 0; j < 5; j++ {
					for k := 0; k < 5; k++ {
						if i == j {
							continue
						}
						Expect(p2.MyCards[j]).NotTo(Equal(p2.MyCards[k]))
					}
				}
			}
		})

		It("Should instantiate a player with no opponent cards by default", func() {
			Expect((len(p.TheirCards))).To(Equal(0))
		})

	})
})
