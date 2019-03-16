package game

import (
	"fmt"
	"math/rand"
)

// Player represents the data of a single player.
type Player struct {
	cardCount int
	hand      []int
	deck      []int
}

// PlayerState represents the state of a particular player
type PlayerState struct {
	CardCount int
	Hand      []int
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
		cardCount: cardCount,
		hand:      []int{},
		deck:      deck,
	}

	return p
}

// NewPlayerFromState constructs a new player
// instance from a player state object
func NewPlayerFromState(state PlayerState) *Player {
	return &Player{
		cardCount: state.CardCount,
		deck:      state.Deck,
		hand:      state.Hand,
	}
}

// Draw draws a card
func (p *Player) Draw() error {
	card, err := nextCard(p.deck)
	if err != nil {
		return err
	}
	p.deck, err = removeCard(p.deck, card)
	if err != nil {
		return err
	}
	p.hand = append(p.hand, card)
	return nil
}

// Give adds the negative card value to the player's hand
func (p *Player) Give(card int) {
	p.hand = append(p.hand, -1*card)
}

// State returns the state of the player
func (p *Player) State() PlayerState {
	return PlayerState{
		Deck:  p.deck,
		Hand:  p.hand,
		Score: p.Score(),
	}
}

// Score finds the...score...of a...hand.
func (p *Player) Score() int {
	return scoreHand(p.hand, p.cardCount)
}

func scoreHand(cards []int, modulus int) int {
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
	cards, err := removeCard(p.hand, card)
	if err != nil {
		return err
	}
	p.hand = cards
	return nil
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
