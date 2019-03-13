package game_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("Player", func() {
	BeforeEach(func() {
		rand.Seed(0)
	})

	Describe("NewPlayer", func() {
		var player *game.Player
		var state game.PlayerState

		BeforeEach(func() {
			player = game.NewPlayer(5)
			state = player.State()
		})

		It("Should instantiate a player with 0 cards by default", func() {
			Expect(len(state.Cards)).To(Equal(0))
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
				Expect(state.Cards[0]).To(BeNumerically(">=", 1))
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
					Expect(s.Cards).To(BeASaneHand())
				}
			})
		})
	})

	Describe("Steal()", func() {
		Describe("When a player has a 1", func() {
			var player *game.Player

			BeforeEach(func() {
				player = game.NewPlayerFromState(game.PlayerState{
					CardCount: 3,
					Cards:     []int{1},
				})
			})

			Describe("When Steal is called with 1", func() {
				BeforeEach(func() {
					err := player.Steal(1)
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should remove the 1 card from the player's hand", func() {
					s := player.State()
					Expect(s.Cards).To(Equal([]int{}))
				})
			})

			Describe("When Steal is called with 2", func() {
				var err error

				BeforeEach(func() {
					err = player.Steal(2)
				})

				It("Should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("Give()", func() {
		Describe("When the player has no cards", func() {
			var player *game.Player

			BeforeEach(func() {
				player = game.NewPlayerFromState(game.PlayerState{
					CardCount: 3,
					Cards:     []int{},
				})
			})

			Describe("When called with a 1", func() {
				BeforeEach(func() {
					player.Give(1)
				})

				It("Adds the negative card value to the player's hand", func() {
					cards := player.State().Cards
					Expect(cards).To(Equal([]int{-1}))
				})
			})
		})
	})

	Describe("Update()", func() {
		Describe("When the player has no cards", func() {
			var player *game.Player

			BeforeEach(func() {
				player = game.NewPlayerFromState(game.PlayerState{
					CardCount: 3,
					Cards:     []int{},
				})
			})

			Describe("When called", func() {
				BeforeEach(func() {
					player.Update()
				})

				It("Should calculate the player's score to be 0", func() {
					s := player.State().Score
					Expect(s).To(Equal(0))
				})
			})
		})
		Describe("When the player has 1 card, and it's a 2", func() {
			var player *game.Player

			BeforeEach(func() {
				player = game.NewPlayerFromState(game.PlayerState{
					CardCount: 3,
					Cards:     []int{2},
				})
			})

			Describe("When called", func() {
				BeforeEach(func() {
					player.Update()
				})

				It("Should calculate the player's score to be 2", func() {
					s := player.State().Score
					Expect(s).To(Equal(2))
				})
			})
		})
	})
})
