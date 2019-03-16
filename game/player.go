package game

import (
	"fmt"
	"math/rand"
)

// Player represents the data of a single player.
type Player struct {
	state PlayerState
}

// PlayerState represents the state of a particular player
type PlayerState struct {
	CardCount int
	Cards     []int
	Deck      []int
	Score     int
}

// NewPlayer returns a new Player instance
func NewPlayer(cardCount int) *Player {
	deck := []int{}

	for i := 1; i <= cardCount; i++ {
		deck = append(deck, i)
	}

	p := &Player{
		state: PlayerState{
			CardCount: cardCount,
			Cards:     []int{},
			Deck:      deck,
		},
	}

	return p
}

// NewPlayerFromState constructs a new player
// instance from a player state object
func NewPlayerFromState(state PlayerState) *Player {
	return &Player{
		state: state,
	}
}

// Draw draws a card
func (p *Player) Draw() error {
	card, err := nextCard(p.state.Deck)
	if err != nil {
		return err
	}
	p.state.Deck, err = removeCard(p.state.Deck, card)
	if err != nil {
		return err
	}
	p.state.Cards = append(p.state.Cards, card)
	return nil
}

// Give adds the negative card value to the player's hand
func (p *Player) Give(card int) {
	p.state.Cards = append(p.state.Cards, -1*card)
}

// State returns the state of the player
func (p *Player) State() PlayerState {
	return p.state
}

// Update updates the score, etc based on the cards the player has
func (p *Player) Update() {
	if len(p.state.Cards) == 0 {
		p.state.Score = 0
		return
	}
	p.state.Score = ScoreHand(p.state.Cards, p.state.CardCount)
	return
}

// ScoreHand finds the...score...of a...hand.
func ScoreHand(cards []int, modulus int) int {
	mySum := 0
	theirSum := 0
	for _, c := range cards {
		if c < 0 {
			theirSum += c
			continue
		}
		mySum += c
	}

	if -1*theirSum > mySum {
		return mySum + theirSum
	}

	return (mySum - theirSum) % modulus
}

// Steal removes the card from the player's hand
func (p *Player) Steal(card int) error {
	var err error
	p.state.Cards, err = removeCard(p.state.Cards, card)
	return err
}

func removeCard(cards []int, card int) ([]int, error) {
	for i, c := range cards {
		if c == card {
			return append(cards[:i], cards[i+1:]...), nil
		}
	}
	return nil, fmt.Errorf("Card %v not found in hand: %v", card, cards)
}

func nextCard(deck []int) (int, error) {
	if len(deck) == 0 {
		return 0, fmt.Errorf("no next card")
	}

	i := rand.Intn(len(deck))
	return deck[i], nil
}
