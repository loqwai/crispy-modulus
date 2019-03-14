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
	// initialHandCount := CardCount / 2
	deck := make([]int, cardCount, cardCount)

	for i := 1; i <= cardCount; i++ {
		deck[i] = i
	}

	p := &Player{
		state: PlayerState{
			CardCount: cardCount,
			Cards:     []int{},
			Deck:      deck,
		},
	}

	// for i := 0; i < initialHandCount; i++ {
	// 	p.Cards[i] = deck[i] + 1
	// }

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
	for i, c := range p.state.Cards {
		if c == card {
			p.state.Cards = append(p.state.Cards[:i], p.state.Cards[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Card %v not found in hand: %v", card, p.state.Cards)
}

func nextCard(deck []int) (int, error) {
	if len(deck) == 0 {
		return 0, fmt.Errorf("no next card")
	}

	i := rand.Intn(len(deck))
	return deck[i], nil
}
