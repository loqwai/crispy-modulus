package game_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("game", func() {
	BeforeEach(func() {
		rand.Seed(0)
	})

	Describe("Game", func() {
		var g game.Game

		BeforeEach(func() {
			g = game.New(3)
		})

		It("Should instantiate 2 players by default", func() {
			players := g.GetState().Players
			Expect(len(players)).To(Equal(2))
		})

		It("Should be the player with the greatest modulus's turn", func() {
			Expect(g.GetState().CurrentPlayer).To(Equal(0))
		})

		Describe("When the second player has the higher modulus", func() {
			BeforeEach(func() {
				g.SetState(game.State{
					CardCount: 3,
					Players: []game.Player{
						game.Player{MyCards: []int{3}},
						game.Player{MyCards: []int{1}},
					},
				})
			})

			It("Should be the second player's turn", func() {
				Expect(g.GetState().CurrentPlayer).To(Equal(1))
			})
		})
	})

	Describe("Player", func() {
		var p game.Player

		BeforeEach(func() {
			p = game.NewPlayer(5)
		})

		It("Should instantiate a player with 2 cards by default", func() {
			Expect((len(p.MyCards))).To(Equal(2))
		})

		It("Should populate the hand with numbers between 1-5", func() {
			for i := 0; i < 2; i++ {
				Expect(p.MyCards[i]).To(BeNumerically(">=", 1))
				Expect(p.MyCards[i]).To(BeNumerically("<=", 5))
			}
		})

		It("Should not populate the hand with the same card twice", func() {
			for i := 0; i < 100; i++ {
				p2 := game.NewPlayer(5)
				for j := 0; j < 2; j++ {
					for k := 0; k < 2; k++ {
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
