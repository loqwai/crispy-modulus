package game_test

import (
	"fmt"
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("Game", func() {
	var g game.Game

	BeforeEach(func() {
		rand.Seed(0)
		g = game.New(3)
		fmt.Println(g)
	})

	It("Should instantiate 2 players by default", func() {
		players := g.State().Players
		Expect(len(players)).To(Equal(2))
	})

	It("Should be the player with the greatest modulus's turn", func() {
		Expect(g.State().CurrentPlayer).To(Equal(0))
	})

	It("Should give both players empty hands", func() {
		Expect(g.State().Players[0].Cards).To(HaveLen(0))
	})

	Describe("ComputeFirstPlayer()()", func() {
		Describe("When the second player has the higher modulus", func() {
			BeforeEach(func() {
				g.SetState(game.State{
					CurrentPlayer: 0,
					CardCount:     3,
					Players: []game.PlayerState{
						game.PlayerState{
							Cards: []int{3},
							Deck:  []int{1, 2},
						},
						game.PlayerState{
							Cards: []int{1},
							Deck:  []int{2, 3},
						},
					},
				})
				g.ComputeFirstPlayer()
			})

			It("Should be the second player's turn", func() {
				Expect(g.State().CurrentPlayer).To(Equal(1))
			})
		})
	})

	Describe("Start()", func() {
		Describe("When called", func() {
			BeforeEach(func() {
				err := g.Start()
				Expect(err).NotTo(HaveOccurred())
			})

			It("Should give player 1 1 card", func() {
				Expect(g.State().Players[0].Cards).To(HaveLen(1))
			})

			It("Should give player 2 1 card", func() {
				Expect(g.State().Players[1].Cards).To(HaveLen(1))
			})
		})
	})

	Describe("SetState()", func() {
		Describe("When the second player has the higher modulus, and Start() hasn't been called", func() {
			BeforeEach(func() {
				g.SetState(game.State{
					CurrentPlayer: 0,
					CardCount:     3,
					Players: []game.PlayerState{
						game.PlayerState{
							Cards: []int{3},
							Deck:  []int{1, 2},
						},
						game.PlayerState{
							Cards: []int{1},
							Deck:  []int{2, 3},
						},
					},
				})
			})

			It("Should not set player 2 to be the current player (in case we are restoring a saved game)", func() {
				Expect(g.State().CurrentPlayer).To(Equal(0))
			})
		})
	})

	Describe("Draw()", func() {
		var s game.State

		BeforeEach(func() {
			g.Start()
			g.Draw()
			s = g.State()
		})

		It("Should add a card to the current player's hand", func() {
			Expect(s.Players[0].Cards).To(HaveLen(2))
		})

		It("Should add a valid card to the current player's hand", func() {
			Expect(s.Players[0].Cards).To(HaveLen(2))
		})

		It("Should become the next player's turn", func() {
			Expect(s.CurrentPlayer).To(Equal(1))
		})

		FIt("Should populate the hand with numbers between 1-5", func() {
			p := s.Players[0]
			for i := 0; i < 2; i++ {
				Expect(p.Cards[i]).To(BeNumerically(">=", 1))
				Expect(p.Cards[i]).To(BeNumerically("<=", 5))
			}
		})

		It("Should not populate the hand with the same card twice", func() {
			Expect(s.Players[0].Cards).To(BeASaneHand())
		})

		Describe("when Draw is called a second time", func() {
			BeforeEach(func() {
				g.Draw()
			})

			It("should be the other player's turn now", func() {
				Expect(g.State().CurrentPlayer).To(Equal(0))
			})

			It("should have a sane hand", func() {
				Expect(s.Players[1].Cards).To(BeASaneHand())
			})
		})

		Describe("when Draw is called and there are no cards to draw", func() {
			var err error
			BeforeEach(func() {
				g.SetState(game.State{
					CardCount: 3,
					Players: []game.PlayerState{
						game.PlayerState{
							Cards: []int{1, 2, 3},
							Deck:  []int{},
						},
						game.PlayerState{
							Cards: []int{1, 2, 3},
							Deck:  []int{},
						},
					},
				})
				err = g.Draw()
			})

			It("should be the other player's turn now", func() {
				Expect(err).To(HaveOccurred())
			})

			It("should have a sane hand", func() {
				Expect(s.Players[1].Cards).To(BeASaneHand())
			})
		})
		Describe("when Draw is called and there are no cards to draw", func() {
			var err error

			BeforeEach(func() {
				g.SetState(game.State{
					CardCount:     3,
					CurrentPlayer: 0,
					Players: []game.PlayerState{
						game.PlayerState{
							CardCount: 3,
							Cards:     []int{2, 3},
							Deck:      []int{},
						}},
				})
				err = g.Draw()
			})

			It("should be the other player's turn now", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Steal()", func() {
		Describe("When both players have one card each", func() {
			BeforeEach(func() {
				g.SetState(game.State{
					CardCount:     3,
					CurrentPlayer: 0,
					Players: []game.PlayerState{
						game.PlayerState{
							Cards: []int{2},
							Deck:  []int{1, 3},
						},
						game.PlayerState{
							Cards: []int{1},
							Deck:  []int{2, 3},
						},
					},
				})
			})

			Describe("When it's the first player's turn", func() {
				Describe("When the first player steals a card", func() {
					BeforeEach(func() {
						err := g.Steal(1)
						Expect(err).NotTo(HaveOccurred())
					})

					It("Should add the negative card to the first player's cards", func() {
						cards := g.State().Players[0].Cards
						Expect(cards).To(Equal([]int{2, -1}))
					})

					It("Should remove the card from the second player's cards", func() {
						cards := g.State().Players[1].Cards
						Expect(cards).To(Equal([]int{}))
					})
				})
			})
		})
	})
})
