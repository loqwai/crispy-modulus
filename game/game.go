package game

import (
	"math/rand"
)

//Player represents the data of a single player.
type Player struct {
	MyCards    []int
	TheirCards []int
}

// Game represents an instance of the game
type Game struct {
	//the gamestate of player 1
	Players []Player
}

// New returns a new Game instance
func New() Game {
	return Game{
		Players: []Player{
			NewPlayer(),
			NewPlayer(),
		},
	}
}

//NewPlayer returns a new Player instance
func NewPlayer() Player {
	cardCount := 10
	initialHandCount := cardCount / 2

	p := Player{
		MyCards:    make([]int, initialHandCount),
		TheirCards: make([]int, 0),
	}

	for i := 0; i < initialHandCount; i++ {
		p.MyCards[i] = rand.Intn(cardCount) + 1
	}

	return p
}

func (g *Game) String() string {
	return `
          | Dealt | Remaining | Total | Modulus |
  Player1 | 2 3 4 |       1 5 |     9 |       4 |
  Player2 | 1 4 5 |       2 3 |    10 |       0 |
	`
}
