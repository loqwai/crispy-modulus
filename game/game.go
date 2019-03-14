package game

// Game represents an instance of the game
type Game interface {
	// Start()
	ComputeFirstPlayer()
	State() State
	SetState(state State)
	Start() error
	Steal(card int) error
	Draw() error
}

// State describe what state the game is currently in. It is serializable
type State struct {
	CardCount     int
	CurrentPlayer int
	Players       []PlayerState
}

type _Game struct {
	state *State
}

// New returns a new Game instance
func New(CardCount int) Game {
	return &_Game{
		state: &State{
			CardCount:     CardCount,
			CurrentPlayer: 0,
			Players: []PlayerState{
				NewPlayer(CardCount).State(),
				NewPlayer(CardCount).State(),
			},
		},
	}
}

func (g *_Game) State() State {
	return *g.state
}

func (g *_Game) ComputeFirstPlayer() {
	currentPlayer := 0
	maxScore := 0

	for i, p := range g.state.Players {
		score := ScoreHand(p.Cards, g.state.CardCount)
		if score > maxScore {
			maxScore = score
			currentPlayer = i
		}
	}
	g.state.CurrentPlayer = currentPlayer
}

func (g *_Game) Start() error {
	g.Draw()
	g.Draw()
	return nil
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

func (g *_Game) Draw() error {
	player := NewPlayerFromState(g.state.Players[g.state.CurrentPlayer])
	err := player.Draw()
	if err != nil {
		return err
	}

	g.state.Players[g.state.CurrentPlayer] = player.State()

	g.state.CurrentPlayer = (g.state.CurrentPlayer + 1) % len(g.state.Players)
	return nil
}

func (g *_Game) Steal(card int) error {
	otherPlayerIndex := (g.state.CurrentPlayer + 1) % len(g.state.Players)
	otherPlayer := NewPlayerFromState(g.state.Players[otherPlayerIndex])
	err := otherPlayer.Steal(card)
	if err != nil {
		return err
	}

	g.state.Players[otherPlayerIndex] = otherPlayer.State()

	player := NewPlayerFromState(g.state.Players[g.state.CurrentPlayer])
	player.Give(card)
	g.state.Players[g.state.CurrentPlayer] = player.State()

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
