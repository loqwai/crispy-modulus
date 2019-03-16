package game_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("Player", func() {
	BeforeEach(func() {
		rand.Seed(0)
	})

	Describe("NewPlayer", func() {
		var player *Player
		var state PlayerState

		BeforeEach(func() {
			player = NewPlayer(5)
			state = player.State()
		})

		It("Should instantiate a player with 0 cards by default", func() {
			Expect(state.Cards).To(HaveLen(0))
		})

		It("Should have a deck of 5 cards", func() {
			Expect(state.Deck).To(HaveLen(5))
		})
	})

	Describe("Draw()", func() {
		Describe("When called", func() {
			var state PlayerState
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(PlayerState{
					CardCount: 1,
					Cards:     []int{},
					Deck:      []int{1},
				})
				err := player.Draw()
				Expect(err).NotTo(HaveOccurred())
				state = player.State()
			})

			It("Should populate the hand with numbers between 1-5", func() {
				Expect(state.Cards[0]).To(Equal(1))
			})

			Describe("When called called a second time", func() {
				var err error

				BeforeEach(func() {
					err = player.Draw()
				})

				It("Should return an error cause the deck is empty", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("Steal()", func() {
		Describe("When a player has a 1", func() {
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(PlayerState{
					CardCount: 3,
					Cards:     []int{1},
					Deck:      []int{},
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
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(PlayerState{
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
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(PlayerState{
					CardCount: 3,
					Cards:     []int{},
				})
				player.Update()
			})

			It("Should calculate the player's score to be 0", func() {
				s := player.State().Score
				Expect(s).To(Equal(0))
			})
		})

		Describe("When the player has 1 card, a 2", func() {
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(PlayerState{
					CardCount: 3,
					Cards:     []int{2},
				})
				player.Update()
			})

			It("Should calculate the player's score to be 2", func() {
				s := player.State().Score
				Expect(s).To(Equal(2))
			})
		})

		Describe("When the player has 2 cards, a 1 and a 2", func() {
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(PlayerState{
					CardCount: 3,
					Cards:     []int{1, 2},
				})
				player.Update()
			})

			It("Should calculate the player's score to be 0", func() {
				s := player.State().Score
				Expect(s).To(Equal(0))
			})
		})

		Describe("When the player has 2 stolen cards, -1 and -2", func() {
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(PlayerState{
					CardCount: 3,
					Cards:     []int{-1, -2},
				})
				player.Update()
			})

			It("Should calculate the player's score to be -3", func() {
				s := player.State().Score
				Expect(s).To(Equal(-3))
			})
		})
	})
})
