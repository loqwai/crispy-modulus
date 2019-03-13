package game

import "fmt"

//Player represents the data of a single player.
type Player struct {
	cardCount int
	state     PlayerState
}

// NewPlayer returns a new Player instance
func NewPlayer(cardCount int) *Player {
	// initialHandCount := cardCount / 2
	// deck := rand.Perm(cardCount)
	p := &Player{
		cardCount: cardCount,
		state: PlayerState{
			MyCards:    make([]int, 0),
			TheirCards: make([]int, 0),
		},
	}

	// for i := 0; i < initialHandCount; i++ {
	// 	p.MyCards[i] = deck[i] + 1
	// }

	return p
}

// NewPlayerFromState constructs a new player
// instance from a player state object
func NewPlayerFromState(cardCount int, data PlayerState) *Player {
	return &Player{
		cardCount: cardCount,
		state:     data,
	}
}

// Draw draws a card
func (p *Player) Draw() error {
	card, err := nextCard(p.cardCount, p.state.MyCards)
	if err != nil {
		return err
	}

	p.state.MyCards = append(p.state.MyCards, card)
	return nil
}

// Give adds the negative card value to the player's hand
func (p *Player) Give(card int) {
	p.state.MyCards = append(p.state.MyCards, -1*card)
}

// State returns the state of the player
func (p *Player) State() PlayerState {
	return p.state
}

// Steal removes the card from the player's hand
func (p *Player) Steal(card int) error {
	for i, c := range p.state.MyCards {
		if c == card {
			p.state.MyCards = append(p.state.MyCards[:i], p.state.MyCards[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Card %v not found in hand: %v", card, p.state.MyCards)
}
