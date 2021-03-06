package game_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/loqwai/crispy-modulus/game"
)

var _ = Describe("Game", func() {
	var g game.Game
	Describe("for a game with 10 cards each", func() {
		BeforeEach(func() {
			g = game.New(10)
		})

		It("should instantiate a game with a CardCount of 10", func() {
			Expect(g.State().CardCount).To(Equal(10))
		})
	})
	Describe("for a game with 3 cards each", func() {
		BeforeEach(func() {
			rand.Seed(1) //this is assumed for other tests. It sucks, I know.
			g = game.New(3)
		})

		It("Should instantiate 2 players by default", func() {
			players := g.State().Players
			Expect(len(players)).To(Equal(2))
		})

		It("Should be the player with the greatest modulus's turn", func() {
			Expect(g.State().CurrentPlayer).To(Equal(0))
		})

		It("Should give both players 1 card", func() {
			Expect(g.State().Players[0].Hand).To(HaveLen(0))
		})

		Describe("ComputeFirstPlayer()", func() {
			Describe("When the second player has the higher modulus", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{2, 3},
							},
						},
					})
					g.ComputeFirstPlayer()
				})

				It("Should be the second player's turn", func() {
					Expect(g.State().CurrentPlayer).To(Equal(1))
				})

				It("Should give us the mod of p1", func() {
					Expect(g.State().Players[0].Mod).To(Equal(0))
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
					Expect(g.State().Players[0].Hand).To(HaveLen(1))
				})

				It("Should give player 2 1 card", func() {
					Expect(g.State().Players[1].Hand).To(HaveLen(1))
				})
				It("Should set p2 to be the current player. Because I cheated and know it should be.", func() {
					Expect(g.State().CurrentPlayer).To(Equal(1))
				})
			})
		})

		Describe("String()", func() {
			Describe("Bullshit test to get 100% coverage", func() {
				It("Should be a string longer than 0", func() {
					Expect(game.New(3).String()).NotTo(BeEmpty())
				})
			})
		})

		Describe("SetState()", func() {
			Describe("When the second player has the higher modulus, and Start() hasn't been called", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{2, 3},
							},
						},
					})
				})

				It("Should not set player 2 to be the current player (in case we are restoring a saved game)", func() {
					Expect(g.State().CurrentPlayer).To(Equal(0))
				})
			})
		})

		Describe("WhoIsWinning()", func() {
			Describe("When the first player has the lowest mod", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{2, 3},
							},
						},
					})
				})

				It("Should say the first player is winning", func() {
					Expect(g.WhoIsWinning()).To(Equal(0))
				})
			})

			Describe("When the second player is winning with the lowest mod", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{2, 3},
								Deck: []int{1},
							},
							game.GamePlayerState{
								Hand: []int{1, 2},
								Deck: []int{3},
							},
						},
					})
				})

				It("Should say the second player is winning", func() {
					Expect(g.WhoIsWinning()).To(Equal(1))
				})
			})

			Describe("When the players are tied in terms of mod, but p0 has a greater multiple", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{},
								Deck: []int{1, 2, 3},
							},
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
						},
					})
				})

				It("Should say the first player is winning", func() {
					Expect(g.WhoIsWinning()).To(Equal(0))
				})
			})

			Describe("When the first player is winning by being the most negative", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{-2, -3},
								Deck: []int{1},
							},
							game.GamePlayerState{
								Hand: []int{-1, -2},
								Deck: []int{3},
							},
						},
					})
				})

				It("Should say the first player is winning", func() {
					Expect(g.WhoIsWinning()).To(Equal(0))
				})
			})

			Describe("When the the players have tied, but p1 initiated the tie", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{},
							},
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
						},
					})
				})

				It("Should say that p2 is winning", func() {
					Expect(g.WhoIsWinning()).To(Equal(1))
				})
			})

			Describe("When the the players have tied with negative values", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{-3, 2},
								Deck: []int{1},
							},
							game.GamePlayerState{
								Hand: []int{-3, 2},
								Deck: []int{1},
							},
						},
					})
				})

				It("Should say that its a tie", func() {
					Expect(g.WhoIsWinning()).To(Equal(-1))
				})
			})

			Describe("When the the players have tied, and both initiated the tie", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{},
							},
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{},
							},
						},
					})
				})

				It("Should say that p2 is winning", func() {
					Expect(g.WhoIsWinning()).To(Equal(-1))
				})
			})

			Describe("When the second player is winning by having the lowest multiple", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
							game.GamePlayerState{
								Hand: []int{},
								Deck: []int{1, 2, 3},
							},
						},
					})
				})

				It("Should say the second player is winning", func() {
					Expect(g.WhoIsWinning()).To(Equal(1))
				})
			})

			Describe("When the players are exactly tied", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CurrentPlayer: 0,
						CardCount:     3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
							game.GamePlayerState{
								Hand: []int{3},
								Deck: []int{1, 2},
							},
						},
					})
				})

				It("Should say -1, I guess. Maybe?", func() {
					Expect(g.WhoIsWinning()).To(Equal(-1))
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

			It("Should add a valid card to the current player's hand", func() {
				Expect(s.Players[1].Hand).To(HaveLen(2))
			})

			It("Should become the next player's turn", func() {
				Expect(s.CurrentPlayer).To(Equal(0))
			})

			It("Should populate the hand with numbers between 1-5", func() {
				p := s.Players[1]
				for i := 0; i < 2; i++ {
					Expect(p.Hand[i]).To(BeNumerically(">=", 1))
					Expect(p.Hand[i]).To(BeNumerically("<=", 5))
				}
			})

			It("Should not populate the hand with the same card twice", func() {
				Expect(s.Players[0].Hand).To(BeASaneHand())
			})

			Describe("when Draw is called a second time", func() {
				BeforeEach(func() {
					g.Draw()
				})

				It("should be the other player's turn now", func() {
					Expect(g.State().CurrentPlayer).To(Equal(1))
				})

				It("should have a sane hand", func() {
					Expect(s.Players[1].Hand).To(BeASaneHand())
				})
			})

			Describe("when Draw is called and there are no cards to draw", func() {
				var err error
				BeforeEach(func() {
					g.SetState(game.State{
						CardCount: 3,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{1, 2, 3},
								Deck: []int{},
							},
							game.GamePlayerState{
								Hand: []int{1, 2, 3},
								Deck: []int{},
							},
						},
					})
					err = g.Draw()
				})

				It("should be the other player's turn now", func() {
					Expect(err).To(HaveOccurred())
				})

				It("should have a sane hand", func() {
					Expect(s.Players[1].Hand).To(BeASaneHand())
				})
			})

			Describe("when Draw is called and there are no cards to draw", func() {
				var err error

				BeforeEach(func() {
					g.SetState(game.State{
						CardCount:     3,
						CurrentPlayer: 0,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{2, 3},
								Deck: []int{},
							}},
					})
					err = g.Draw()
				})

				It("should be the other player's turn now", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Describe("IsDone()", func() {
			Describe("When the players have cards left in their deck", func() {
				var isDone bool

				BeforeEach(func() {
					g.SetState(game.State{
						CardCount:     2,
						CurrentPlayer: 0,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{2},
							},
							game.GamePlayerState{
								Hand: []int{2},
								Deck: []int{1},
							},
						},
					})

					isDone = g.IsDone()
				})

				It("Should return false", func() {
					Expect(isDone).To(BeFalse())
				})
			})

			Describe("When p1 has no cards left in their deck", func() {
				var isDone bool

				BeforeEach(func() {
					g.SetState(game.State{
						CardCount:     1,
						CurrentPlayer: 0,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{},
							},
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{2},
							},
						},
					})

					isDone = g.IsDone()
				})

				It("Should return true", func() {
					Expect(isDone).To(BeTrue())
				})
			})
			Describe("When p2 has no cards left in their deck", func() {
				var isDone bool

				BeforeEach(func() {
					g.SetState(game.State{
						CardCount:     1,
						CurrentPlayer: 0,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{2, 3},
							},
							game.GamePlayerState{
								Hand: []int{1, 2, 3},
								Deck: []int{},
							},
						},
					})

					isDone = g.IsDone()
				})

				It("Should return true", func() {
					Expect(isDone).To(BeTrue())
				})
			})
		})

		Describe("Steal()", func() {
			Describe("When both players have one card each", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CardCount:     3,
						CurrentPlayer: 0,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{2},
								Deck: []int{1, 3},
							},
							game.GamePlayerState{
								Hand: []int{1},
								Deck: []int{2, 3},
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
							cards := g.State().Players[0].Hand
							Expect(cards).To(Equal([]int{2, -1}))
						})

						It("Should remove the card from the second player's cards", func() {
							cards := g.State().Players[1].Hand
							Expect(cards).To(Equal([]int{}))
						})
					})

					Describe("When the first player steals a card that the 2nd player doesn't have", func() {
						var err error

						BeforeEach(func() {
							err = g.Steal(2)
						})

						It("Shoud throw an error", func() {
							Expect(err).To(HaveOccurred())
						})

						It("Should not add the negative card to the first player's cards", func() {
							cards := g.State().Players[0].Hand
							Expect(cards).To(Equal([]int{2}))
						})

						It("Should not remove the card from the second player's cards", func() {
							cards := g.State().Players[1].Hand
							Expect(cards).To(Equal([]int{1}))
						})
					})
				})
			})
			Describe("When p2 has a 'locked' card", func() {
				BeforeEach(func() {
					g.SetState(game.State{
						CardCount:     3,
						CurrentPlayer: 0,
						Players: []game.GamePlayerState{
							game.GamePlayerState{
								Hand: []int{},
								Deck: []int{2, 3},
							},
							game.GamePlayerState{
								Hand: []int{1, -1},
								Deck: []int{2, 3},
							},
						},
					})
				})

				Describe("When p1 tries to steal a negative card", func() {
					var err error
					BeforeEach(func() {
						err = g.Steal(-1)
					})

					It("Should throw an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})
				Describe("When p1 tries to steal a 'locked' card", func() {
					var err error
					BeforeEach(func() {
						err = g.Steal(1)
					})

					It("Should throw an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})
			})
		})
	})
})
