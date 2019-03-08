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
type Game interface {
	GetState() State
}

// State describe what state the game is currently in. It is serializable
type State struct {
	Players []Player
}

// New returns a new Game instance
func New() Game {
	return &_Game{
		state: &State{
			Players: []Player{
				NewPlayer(),
				NewPlayer(),
			},
		},
	}
}

type _Game struct {
	state *State
}

// NewPlayer returns a new Player instance
func NewPlayer() Player {
	cardCount := 10
	initialHandCount := cardCount / 2
	deck := rand.Perm(cardCount)
	p := Player{
		MyCards:    make([]int, 5),
		TheirCards: make([]int, 0),
	}

	for i := 0; i < initialHandCount; i++ {
		p.MyCards[i] = deck[i] + 1
	}

	return p
}

func (g *_Game) GetState() State {
	return *g.state
}

func (g *_Game) String() string {
	return `
          | Dealt | Remaining | Total | Modulus |
  Player1 | 2 3 4 |       1 5 |     9 |       4 |
  Player2 | 1 4 5 |       2 3 |    10 |       0 |
	`
}
