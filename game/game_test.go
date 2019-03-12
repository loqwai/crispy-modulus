package game_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

func IsHandSane(cards []int) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == j {
				continue
			}
			if cards[i] == cards[j] {
				return false
			}
		}
	}
	return true
}

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

		It("Should give both players empty hands", func() {
			Expect(g.GetState().Players[0].MyCards).To(HaveLen(0))
		})

		Describe("ComputeFirstPlayer()()", func() {
			Describe("When the second player has the higher modulus", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.PlayerState{
							game.PlayerState{MyCards: []int{3}},
							game.PlayerState{MyCards: []int{1}},
						},
					})
					g.ComputeFirstPlayer()
				})

				It("Should be the second player's turn", func() {
					Expect(g.GetState().CurrentPlayer).To(Equal(1))
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
					Expect(g.GetState().Players[0].MyCards).To(HaveLen(1))
				})

				It("Should give player 2 1 card", func() {
					Expect(g.GetState().Players[1].MyCards).To(HaveLen(1))
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
							game.PlayerState{MyCards: []int{3}},
							game.PlayerState{MyCards: []int{1}},
						},
					})
				})

				It("Should not set player 2 to be the current player (in case we are restoring a saved game)", func() {
					Expect(g.GetState().CurrentPlayer).To(Equal(0))
				})
			})

		})

		Describe("Draw()", func() {
			var s game.State

			BeforeEach(func() {
				g.Start()
				g.Draw()
				s = g.GetState()
			})

			It("Should add a card to the current player's hand", func() {
				Expect(s.Players[0].MyCards).To(HaveLen(2))
			})

			It("Should add a valid card to the current player's hand", func() {
				Expect(s.Players[0].MyCards).To(HaveLen(2))
			})

			It("Should become the next player's turn", func() {
				Expect(s.CurrentPlayer).To(Equal(1))
			})

			It("Should populate the hand with numbers between 1-5", func() {
				p := s.Players[0]
				for i := 0; i < 2; i++ {
					Expect(p.MyCards[i]).To(BeNumerically(">=", 1))
					Expect(p.MyCards[i]).To(BeNumerically("<=", 5))
				}
			})

			It("Should not populate the hand with the same card twice", func() {
				Expect(s.Players[0].MyCards).To(BeASaneHand())
			})

			Describe("when Draw is called a second time", func() {
				BeforeEach(func() {
					g.Draw()
				})

				It("should be the other player's turn now", func() {
					Expect(g.GetState().CurrentPlayer).To(Equal(0))
				})

				It("should have a sane hand", func() {
					Expect(s.Players[1].MyCards).To(BeASaneHand())
				})
			})

			Describe("when Draw is called and there are no cards to draw", func() {
				var err error
				BeforeEach(func() {
					g.SetState(game.State{
						CardCount: 3,
						Players: []game.PlayerState{
							game.PlayerState{MyCards: []int{1, 2, 3}},
							game.PlayerState{MyCards: []int{1, 2, 3}},
						},
					})
					err = g.Draw()
				})

				It("should be the other player's turn now", func() {
					Expect(err).To(HaveOccurred())
				})

				It("should have a sane hand", func() {
					Expect(s.Players[1].MyCards).To(BeASaneHand())
				})
			})
		})
	})

	Describe("Player", func() {
		Describe("NewPlayer", func() {
			var player *game.Player
			var state game.PlayerState

			BeforeEach(func() {
				player = game.NewPlayer(5)
				state = player.State()
			})

			It("Should instantiate a player with 0 cards by default", func() {
				Expect(len(state.MyCards)).To(Equal(0))
			})
		})

		Describe("Draw()", func() {
			Describe("When called", func() {
				var state game.PlayerState

				BeforeEach(func() {
					player := game.NewPlayer(5)
					err := player.Draw()
					Expect(err).NotTo(HaveOccurred())
					state = player.State()
				})

				It("Should populate the hand with numbers between 1-5", func() {
					Expect(state.MyCards[0]).To(BeNumerically(">=", 1))
				})
			})

			Describe("When called 100 times", func() {
				var states []game.PlayerState

				BeforeEach(func() {
					for i := 0; i < 100; i++ {
						states = append(states, game.NewPlayer(5).State())
					}
				})

				It("Should not populate the hand with the same card twice", func() {
					for _, s := range states {
						Expect(s.MyCards).To(BeASaneHand())
					}
				})
			})
		})

		// It("Should instantiate a player with no opponent cards by default", func() {
		// 	Expect(len(p.TheirCards)).To(Equal(0))
		// })
	})
})
