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
			Expect(state.Hand).To(HaveLen(0))
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
				player = NewPlayerFromState(1, PlayerState{
					Hand: []int{},
					Deck: []int{1},
				})
				err := player.Draw()
				Expect(err).NotTo(HaveOccurred())
				state = player.State()
			})

			It("Should populate the hand with numbers between 1-5", func() {
				Expect(state.Hand[0]).To(Equal(1))
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
				player = NewPlayerFromState(3, PlayerState{
					Hand: []int{1},
					Deck: []int{},
				})
			})

			Describe("When Steal is called with 1", func() {
				BeforeEach(func() {
					err := player.Steal(1)
					Expect(err).ToNot(HaveOccurred())
				})

				It("Should remove the 1 card from the player's hand", func() {
					s := player.State()
					Expect(s.Hand).To(Equal([]int{}))
				})
			})

			Describe("When Steal is called with 2", func() {
				var err error
				var state PlayerState

				BeforeEach(func() {
					err = player.Steal(2)
					state = player.State()
				})

				It("Should return an error", func() {
					Expect(err).To(HaveOccurred())
				})

				It("Should not remove any cards from the player's hand", func() {
					Expect(state.Hand).To(Equal([]int{1}))
				})
			})
		})
	})

	Describe("Give()", func() {
		Describe("When the player has no cards", func() {
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(3, PlayerState{
					Hand: []int{},
					Deck: []int{},
				})
			})

			Describe("When called with a 1", func() {
				BeforeEach(func() {
					player.Give(1)
				})

				It("Adds the negative card value to the player's hand", func() {
					cards := player.State().Hand
					Expect(cards).To(Equal([]int{-1}))
				})
			})
		})
	})

	Describe("Score()", func() {
		Describe("When the player has no cards", func() {
			var player *Player
			var score int

			BeforeEach(func() {
				player = NewPlayerFromState(3, PlayerState{
					Deck: []int{},
					Hand: []int{},
				})
				score = player.Score()
			})

			It("Should calculate the player's score to be 0", func() {
				Expect(score).To(Equal(0))
			})
		})

		Describe("When the player has 1 card, a 2", func() {
			var player *Player
			var score int

			BeforeEach(func() {
				player = NewPlayerFromState(3, PlayerState{
					Deck: []int{},
					Hand: []int{2},
				})
				score = player.Score()
			})

			It("Should calculate the player's score to be 2", func() {
				Expect(score).To(Equal(2))
			})
		})

		Describe("When the player has 2 cards, a 1 and a 2", func() {
			var score int
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(3, PlayerState{
					Deck: []int{},
					Hand: []int{1, 2},
				})
				score = player.Score()
			})

			It("Should calculate the player's score to be 0", func() {
				Expect(score).To(Equal(0))
			})
		})

		Describe("When the player has 2 stolen cards, -1 and -2", func() {
			var score int
			var player *Player

			BeforeEach(func() {
				player = NewPlayerFromState(2, PlayerState{
					Deck: []int{},
					Hand: []int{-1, -2},
				})
				score = player.Score()
			})

			It("Should calculate the player's score to be -3", func() {
				Expect(score).To(Equal(-3))
			})
		})
	})
})
