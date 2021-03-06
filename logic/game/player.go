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
	Hand []int
	Deck []int
	Sum  int
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
func NewPlayerFromState(cardCount int, state PlayerState) *Player {
	return &Player{
		cardCount: cardCount,
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
		Deck: p.deck,
		Hand: p.hand,
		Sum:  p.Sum(),
	}
}

// Sum adds up all the cards in the player's hand
func (p *Player) Sum() int {
	s := 0

	for i := 0; i < len(p.hand); i++ {
		s += p.hand[i]
	}

	return s
}

//RemainingCardCount returns the number of cards left in the player's hand
func (p *Player) RemainingCardCount() int {
	return len(p.deck)
}

// Steal removes the card from the player's hand
func (p *Player) Steal(card int) error {
	var err error
	if card < 1 {
		return fmt.Errorf("No takeseys backseys")
	}

	for _, c := range p.hand {
		if c == -1*card {
			return fmt.Errorf("Card %v is locked", card)
		}
	}

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
