package game

import (
	"fmt"
	"math/rand"
)

//Player represents the data of a single player.
type Player struct {
	MyCards    []int
	TheirCards []int
}

// Game represents an instance of the game
type Game interface {
	Start()
	GetState() State
	SetState(state State)
	Draw() error
}

// State describe what state the game is currently in. It is serializable
type State struct {
	CardCount     int
	CurrentPlayer int
	Players       []Player
}

// New returns a new Game instance
func New(cardCount int) Game {
	return &_Game{
		state: &State{
			CardCount:     cardCount,
			CurrentPlayer: 0,
			Players: []Player{
				NewPlayer(cardCount),
				NewPlayer(cardCount),
			},
		},
	}
}

type _Game struct {
	state *State
}

// NewPlayer returns a new Player instance
func NewPlayer(cardCount int) Player {
	initialHandCount := cardCount / 2
	deck := rand.Perm(cardCount)
	p := Player{
		MyCards:    make([]int, initialHandCount),
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

func (g *_Game) Start() {
	currentPlayer := 0
	maxMod := 0

	for i, p := range g.state.Players {
		sum := 0
		for _, c := range p.MyCards {
			sum += c
		}

		mod := sum % g.state.CardCount
		if mod > maxMod {
			maxMod = mod
			currentPlayer = i
		}
	}
	g.state.CurrentPlayer = currentPlayer
}

func (g *_Game) SetState(state State) {
	g.state = &state
}

func (g *_Game) String() string {
	return `
          | Dealt | Remaining | Total | Modulus |
  Player1 | 2 3 4 |       1 5 |     9 |       4 |
  Player2 | 1 4 5 |       2 3 |    10 |       0 |
	`
}

func nextCard(modulus int, hand []int) (int, error) {
	if len(hand) >= modulus {
		return 0, fmt.Errorf("no next card")
	}

	card := 0
	for {
		card = rand.Intn(modulus) + 1
		if !contains(hand, card) {
			return card, nil
		}
	}
}

func (g *_Game) Draw() error {
	currentPlayer := g.state.Players[g.state.CurrentPlayer]
	card, err := nextCard(g.state.CardCount, currentPlayer.MyCards)
	if err != nil {
		return err
	}

	g.state.Players[g.state.CurrentPlayer].MyCards = append(currentPlayer.MyCards, card)
	g.state.CurrentPlayer = (g.state.CurrentPlayer + 1) % len(g.state.Players)
	return nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
