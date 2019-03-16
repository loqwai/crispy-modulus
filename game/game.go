package game

import "encoding/json"

// Game represents an instance of the game
type Game interface {
	ComputeFirstPlayer()
	Draw() error
	IsDone() bool
	SetState(state State)
	Start() error
	State() State
	String() (string, error)
	Steal(card int) error
	WhoIsWinning() int
}

// State describe what state the game is currently in. It is serializable
type State struct {
	CardCount     int
	CurrentPlayer int
	Players       []PlayerState
}

type _Game struct {
	cardCount     int
	currentPlayer int
	players       []*Player
}

// New returns a new Game instance
func New(cardCount int) Game {
	return &_Game{
		cardCount:     cardCount,
		currentPlayer: 0,
		players: []*Player{
			NewPlayer(cardCount),
			NewPlayer(cardCount),
		},
	}
}

func (g *_Game) State() State {
	playerStates := make([]PlayerState, len(g.players))
	for i, p := range g.players {
		playerStates[i] = p.State()
	}
	return State{
		CardCount:     g.cardCount,
		CurrentPlayer: g.currentPlayer,
		Players:       playerStates,
	}
}

func (g *_Game) ComputeFirstPlayer() {
	currentPlayer := 0
	maxMod := 0

	for i, p := range g.players {
		mod := p.Sum() % g.cardCount
		if mod > maxMod {
			maxMod = mod
			currentPlayer = i
		}
	}
	g.currentPlayer = currentPlayer
}

func (g *_Game) Start() error {
	g.Draw()
	g.Draw()
	g.ComputeFirstPlayer()

	return nil
}

func (g *_Game) SetState(state State) {
	g.cardCount = state.CardCount
	g.currentPlayer = state.CurrentPlayer
	g.players = make([]*Player, len(state.Players))
	for i, p := range state.Players {
		g.players[i] = NewPlayerFromState(state.CardCount, p)
	}
}

func (g *_Game) String() (string, error) {
	output, err := json.MarshalIndent(g.State(), "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (g *_Game) Draw() error {
	player := g.players[g.currentPlayer]
	err := player.Draw()
	if err != nil {
		return err
	}

	g.currentPlayer = (g.currentPlayer + 1) % len(g.players)
	return nil
}

func (g *_Game) IsDone() bool {
	if len(g.State().Players[0].Deck) == 0 {
		return true
	}
	if len(g.State().Players[1].Deck) == 0 {
		return true
	}
	return false
}

func (g *_Game) Steal(card int) error {
	otherPlayerIndex := (g.currentPlayer + 1) % len(g.players)

	err := g.players[otherPlayerIndex].Steal(card)
	if err != nil {
		return err
	}

	g.players[g.currentPlayer].Give(card)

	g.currentPlayer = (g.currentPlayer + 1) % len(g.players)
	return nil
}

func (g *_Game) WhoIsWinning() int {
	bestPlayer := struct {
		index    int
		mod      int
		multiple int
	}{
		index:    0,
		mod:      g.players[0].Sum() % g.cardCount,
		multiple: g.players[0].Sum() / g.cardCount,
	}

	bestWorstPlayer := -1
	mostNegativeSum := 0

	for i, p := range g.players {
		if p.Sum() < mostNegativeSum {
			bestWorstPlayer = i
			mostNegativeSum = p.Sum()
			continue
		}
		mod := p.Sum() % g.cardCount
		multiple := p.Sum() / g.cardCount

		if mod > bestPlayer.mod {
			continue
		}

		if mod < bestPlayer.mod {
			bestPlayer.index = i
			bestPlayer.mod = mod
			bestPlayer.multiple = multiple
			continue
		}

		//ties
		if p.RemainingCardCount() == 0 {
			continue
		}

		if multiple <= bestPlayer.multiple {
			bestPlayer.index = i
			bestPlayer.mod = mod
			bestPlayer.multiple = multiple
			continue
		}
	}

	if g.players[bestPlayer.index].Sum() >= 0 {
		return bestPlayer.index
	}
	return bestWorstPlayer
}
